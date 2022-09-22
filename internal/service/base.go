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
	"strings"
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
	publish(g.Cfg().MustGet(ctx, "pub_topics.Publish_register_get").String(), string(msg))
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
	publish(g.Cfg().MustGet(nil, "pub_topics.Publish_realtime_data_get").String(), string(msg))
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
	fmt.Println(string(msg))
	publish(g.Cfg().MustGet(nil, "pub_topics.Publish_history_data_get").String(), string(msg))
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
	//fmt.Println(in)
	msg, _ := json.Marshal(model.MqttDataBaseGetAlarmIn{
		Token:     "1000",
		Time_type: "timestartgather",
		StartTime: gtime.NewFromStr(in.StartTime).Layout("2006-01-02T15:04:05.000-0700"),
		EndTime:   gtime.NewFromStr(in.EndTime).Layout("2006-01-02T15:04:05.000-0700"),
		SourType:  "104",
		Body: []model.MqttDataBaseGetAlarmInBody{
			{
				Model:    strings.Split(in.Dev, "_")[0],
				Totaldev: "0",
				Dev: []string{
					in.Dev,
				},
			},
		},
	})
	fmt.Println(string(msg))
	publish(g.Cfg().MustGet(nil, "pub_topics.Publish_alarm_data_get").String(), string(msg))
	alarmData := model.MqttDataBaseGetAlarmOut{}
	//fmt.Println(alarm_data)
	select {
	case alarmData = <-alarmChan:
	case <-time.After(3 * time.Second):
		err = fmt.Errorf("读取超时")
		return
	}
	out = model.BaseAlarmOut{
		Alarm: []model.AlarmResAlarm{},
	}
	for _, bb := range alarmData.Body {
		alarm := model.AlarmResAlarm{}
		alarm.Timestamp = gtime.NewFromStr(bb.Timestamp)
		alarm.AlarmType = bb.Event //consts.AlarmDict[bb.Event]
		alarm.Remark = "分->合"
		alarm.Status = "已读"
		//fmt.Println(bb.Extdata)
		out.Alarm = append(out.Alarm, alarm)
	}
	//fmt.Println(alarmChan)
	//fmt.Println(len(alarmChan))

	return out, nil
}
func (s *sBase) BaseFaultWaveform(ctx context.Context, in model.BaseFaultWaveformIn) (out model.BaseFaultWaveformOut, err error) {
	// 获取已有的alarmRes，找到对应时间戳
	out = model.BaseFaultWaveformOut{
		Waveform: []model.BaseFaultWaveformOutBody{},
	}
	alarmTime, _ := time.Parse("2006-01-02T15:04:05.000+0800", in.Timestamp.Time.Format("2006-01-02T15:04:05.000+0800")) // 记录告警的时间
	fmt.Println(alarmTime)
	for _, bb := range alarmRes.Body {
		fmt.Println(123)
		fmt.Println(bb.Timestamp)
		time2, _ := time.Parse("2006-01-02T15:04:05.000+0800", bb.Timestamp)
		//time1 := gtime.NewFromStr(bb.Timestamp).Layout("2006-01-02T15:04:05.000-0700")
		fmt.Println(time2)
		//fmt.Println(time2.Weekday()) 可以获得星期几
		//fmt.Println(time2.Hour()) 可以获得小时
		if alarmTime == time2 {
			for i := 0; i < len(bb.Extdata)/4; i++ {
				wave := model.BaseFaultWaveformOutBody{}
				newTime, _ := time.Parse("2006-01-02T15:04:05.000+0800", bb.Extdata[i*4].Timestamp)

				wave.OffsetTime = (newTime.Sub(alarmTime)).Seconds()
				value0, _ := strconv.ParseFloat(bb.Extdata[i*4].Val, 64)
				wave.Ia = float32(value0)
				value1, _ := strconv.ParseFloat(bb.Extdata[i*4+1].Val, 64)
				wave.Ib = float32(value1)
				value2, _ := strconv.ParseFloat(bb.Extdata[i*4+2].Val, 64)
				wave.Ic = float32(value2)
				value3, _ := strconv.ParseFloat(bb.Extdata[i*4+3].Val, 64)
				wave.In = float32(value3)
				out.Waveform = append(out.Waveform, wave)
			}
		}
	}
	return out, nil

}

func (s *sBase) BaseGetConfig(ctx context.Context, in model.BaseGetConfigIn) (out model.BaseGetConfigOut, err error) {

	msg, _ := json.Marshal(model.MqttDataBaseGetConfigIn{
		Dev: in.Dev,
	})
	fmt.Println(string(msg))
	publish(g.Cfg().MustGet(nil, "pub_topics.Publish_getParams").String(), string(msg))
	getConfigData := model.MqttDataBaseGetConfigOut{}
	fmt.Println(getConfigData)
	select {
	case getConfigData = <-getConfigChan:
	case <-time.After(3 * time.Second):
		err = fmt.Errorf("读取超时")
		return
	}
	out = model.BaseGetConfigOut{
		Dev: getConfigData.Dev,
	}
	if len(getConfigData.Body) > 4 {
		out.LeakageProtectionStatus = getConfigData.Body[0].Val
		out.RatedLeakageProtectionDifference, _ = strconv.Atoi(getConfigData.Body[3].Val)
		out.InterpolationProtectionActionTime, _ = strconv.Atoi(getConfigData.Body[4].Val)
		out.ThresholdProtectionActionTime, _ = strconv.Atoi(getConfigData.Body[2].Val)
		out.RatedProtectionCurrentThreshold, _ = strconv.Atoi(getConfigData.Body[1].Val)

	}

	fmt.Println(getConfigData)

	return out, nil
}

func (s *sBase) BaseSetConfig(ctx context.Context, in model.BaseSetConfigIn) (out model.BaseSetConfigOut, err error) {
	msg, _ := json.Marshal(model.MqttDataBaseSetConfigIn{
		Dev: in.Dev,
		Body: []model.MqttDatabaseSetConfigInBody{
			{
				Name: "leakage_protection_status",
				Val:  in.LeakageProtectionStatus,
			},
			{
				Name: "rated_protection_current_threshold",
				Val:  strconv.Itoa(in.RatedProtectionCurrentThreshold),
			},
			{
				Name: "threshold_protection_action_time",
				Val:  strconv.Itoa(in.ThresholdProtectionActionTime),
			},
			{
				Name: "rated_leakage_protection_difference",
				Val:  strconv.Itoa(in.RatedLeakageProtectionDifference),
			},
			{
				Name: "interpolation_protection_action_time",
				Val:  strconv.Itoa(in.InterpolationProtectionActionTime),
			},
		},
	})
	fmt.Println(string(msg))
	publish(g.Cfg().MustGet(nil, "pub_topics.Publish_setParams").String(), string(msg))
	setConfigData := model.MqttDataBaseSetConfigOut{}
	select {
	case setConfigData = <-setConfigChan:
	case <-time.After(3 * time.Second):
		err = fmt.Errorf("读取超时")
		return
	}
	fmt.Println(setConfigData)
	out = model.BaseSetConfigOut{}
	return out, nil
}

func (s *sBase) BaseDayAnalysis(ctx context.Context, in model.BaseDayAnaIn) (out model.BaseDayAnaOut, err error) {

	msg, _ := json.Marshal(model.MqttDataBaseGetAlarmIn{
		Token:     "1000",
		Time_type: "timestartgather",
		StartTime: gtime.NewFromStr(in.StartTime.String()).Layout("2006-01-02T15:04:05.000-0700"),
		EndTime:   gtime.NewFromStr(in.EndTime.EndOfDay().String()).Layout("2006-01-02T15:04:05.000-0700"),
		SourType:  "104",
		Body: []model.MqttDataBaseGetAlarmInBody{
			{
				Model:    strings.Split(in.Dev, "_")[0],
				Totaldev: "0",
				Dev: []string{
					in.Dev,
				},
			},
		},
	})

	fmt.Println(string(msg))
	publish(g.Cfg().MustGet(nil, "pub_topics.Publish_alarm_data_get").String(), string(msg))
	alarmData := model.MqttDataBaseGetAlarmOut{}
	//fmt.Println(alarm_data)
	select {
	case alarmData = <-alarmChan:
	case <-time.After(10 * time.Second):
		err = fmt.Errorf("读取超时")
		return
	}
	out = model.BaseDayAnaOut{
		DayAna: [24][7]int{},
	}
	fmt.Println(len(alarmData.Body))
	for _, bb := range alarmData.Body {
		hour := gtime.NewFromStr(bb.Timestamp).Hour()

		day := gtime.NewFromStr(bb.Timestamp).Weekday().String()[0:2]
		//fmt.Println(bb.Timestamp)
		//fmt.Print(hour)
		//fmt.Print(", ")
		//fmt.Print(day)
		switch {
		case day == "Mo":
			out.DayAna[hour][0] += 1
			//fmt.Println("+1")
		case day == "Tu":
			out.DayAna[hour][1] += 1
			//fmt.Println("+1")
		case day == "We":
			out.DayAna[hour][2] += 1
			//fmt.Println("+1")
		case day == "Th":
			out.DayAna[hour][3] += 1
			//fmt.Println("+1")
		case day == "Fr":
			out.DayAna[hour][4] += 1
			//fmt.Println("+1")
		case day == "Sa":
			out.DayAna[hour][5] += 1
			//fmt.Println("+1")
		case day == "Su":
			out.DayAna[hour][6] += 1
			//fmt.Println("+1")
		}
	}
	//fmt.Println(alarmChan)
	//fmt.Println(len(alarmChan))
	fmt.Println(out.DayAna)

	return out, nil
}
