package service

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
	"time"
	"ttu-backend/internal/consts"
	"ttu-backend/internal/model"
)

type (
	// sBase is service struct of module Base.
	sBase struct{}
)

var (
	// insBase is the instance of service Base.
	insBase = sBase{}
)

// Base returns the interface of Base service.
func Base() *sBase {
	return &insBase
}

func (s *sBase) BaseDeviceList(ctx context.Context, in model.BaseDeviceListIn) (out model.BaseDeviceListOut, err error) {
	if err != nil {
		return out, err
	}
	out.Group = []model.DeviceListResGroup{}
	msg, _ := json.Marshal(model.MqttDatabaseGetTopoIn{
		Token:     "1000",
		Timestamp: "",
		Body:      []string{},
	})
	publish(consts.Publish_register_get, string(msg))
	topo := model.MqttDatabaseGetTopoOut{}
	select {
	case topo = <-topoChan:
	case <-time.After(3 * time.Second):
		err = fmt.Errorf("读取超时")
		return
	}
	modelName := []string{"LTU", "Switch"}
	for _, t := range topo.Body {
		for _, m := range modelName {
			if t.Model != m {
				continue
			}
			group := model.DeviceListResGroup{
				Name:   consts.ModelConfig[t.Model].Name,
				Total:  len(t.Body),
				Online: len(t.Body),
				Icon:   "",
				Device: []model.DeviceListResGroupDevice{},
			}
			for _, d := range t.Body {
				device := model.DeviceListResGroupDevice{
					Dev:    d.Dev,
					Name:   "设备" + d.Dev[len(d.Dev)-4:len(d.Dev)],
					Online: true,
					Badge:  0,
				}
				group.Device = append(group.Device, device)
			}
			out.Group = append(out.Group, group)
			out.Station = "10kV古城大道线路—邱家咀 3#台区"
		}
	}
	_ = TopoHandler(topo, modelName)
	return
}
func (s *sBase) BaseRealtime(ctx context.Context, in model.BaseRealtimeIn) (out model.BaseRealtimeOut, err error) {
	msg, _ := json.Marshal(model.MqttDatabaseGetRealtimeIn{
		Token:     "1000",
		Timestamp: "",
		Body: []model.MqttDatabaseGetRealtimeInBody{
			{
				Dev:       in.Dev,
				Totalcall: "1",
				Body:      []string{},
			},
		},
	})
	publish(consts.Publish_realtime_data_get, string(msg))
	realtime := model.MqttDatabaseGetRealtimeOut{}
	select {
	case realtime = <-realtimeChan:
	case <-time.After(3 * time.Second):
		err = fmt.Errorf("读取超时")
		return
	}
	if len(realtime.Body) == 0 || realtime.Body[0].Dev != in.Dev {
		err = fmt.Errorf("获取失败")
		return
	}
	paras := g.Map{}
	for _, para := range realtime.Body[0].Body {
		paras[para.Name] = para.Val
	}
	record := model.RecordResRecord{}
	_ = gconv.Struct(paras, &record, map[string]string{
		"AMax_phsA": "InMax",
		"BMax_phsA": "InAvg",
		"A_phsA":    "Ia",
		"B_phsA":    "Ib",
		"C_phsA":    "Ic",
		"PhV_phsA":  "Ua",
		"PhV_phsB":  "Ub",
		"PhV_phsC":  "Uc",
		"Tmp":       "T",
		"EnvHum":    "H",
	})
	record.Tn = record.T
	record.Timestamp = gtime.NewFromStr(realtime.Timestamp)
	out = model.BaseRealtimeOut{
		Dev:   in.Dev,
		Count: 1,
		Record: []model.RecordResRecord{
			record,
		},
	}
	return out, nil
}
func (s *sBase) BaseRecord(ctx context.Context, in model.BaseRecordIn) (out model.BaseRecordOut, err error) {
	frozenDev := GetFrozenDev(in.Dev)
	if DeviceList.DevList == nil {
		_, _ = Base().BaseDeviceList(nil, model.BaseDeviceListIn{})
	}
	msg, _ := json.Marshal(model.MqttDatabaseGetHistoryIn{
		Token:      "1000",
		TimeType:   "timestamp",
		StartTime:  gtime.NewFromStr(in.StartTime).Layout("2006-01-02T15:04:05.000-0700"),
		EndTime:    gtime.NewFromStr(in.EndTime).Layout("2006-01-02T15:04:05.000-0700"),
		TimeSpan:   strconv.Itoa(in.TimeSpanNumber),
		FrozenType: in.TimeSpanUnit,
		Body: model.MqttDatabaseGetHistoryInBody{
			Dev:  frozenDev,
			Body: GetModelParasByDev(in.Dev),
		},
	})
	publish(consts.Publish_history_data_get, string(msg))
	history := model.MqttDatabaseGetHistoryOut{}
	select {
	case history = <-historyChan:
	case <-time.After(3 * time.Second):
		err = fmt.Errorf("读取超时")
		return
	}
	if history.Body.Dev != frozenDev {
		err = fmt.Errorf("获取失败")
		return
	}

	out = model.BaseRecordOut{
		Dev:    in.Dev,
		Count:  len(history.Body.Body),
		Record: []model.RecordResRecord{},
	}

	for _, bb := range history.Body.Body {
		paras := g.Map{}
		for _, para := range bb.Body {
			paras[para.Name] = para.Val
		}
		record := model.RecordResRecord{}
		_ = gconv.Struct(paras, &record, consts.ModelConfig[GetModelNameByDev(in.Dev)].Mapping)
		record.Tn = record.T
		record.Timestamp = gtime.NewFromStr(bb.Timestamp)
		out.Record = append(out.Record, record)
	}
	return out, nil
}
func (s *sBase) BaseAlarm(ctx context.Context, in model.BaseAlarmIn) (out model.BaseAlarmOut, err error) {
	return
}
