package handler

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"ttu-backend/apiv1"
	"ttu-backend/internal/consts"
	"ttu-backend/internal/model"
	"ttu-backend/internal/service"
)

var (
	Data = hData{}
)

type hData struct{}

// DataInit 加载初始数据，设备列表和对应模型
func (h *hData) DataInit(ctx context.Context, req *apiv1.DataInitReq) (res *apiv1.DataRes, err error) {
	res = &apiv1.DataRes{}
	in := model.EmptyIn{
		Token:     "123",
		Timestamp: "2019-12-25T00:00:00.727+0800",
		Body:      "",
	}
_:
	out, err := service.Mqtt().MqttInit(ctx, in)
	gconv.Struct(out, &res)
	return res, nil
}

// DataHistory 按时间段访问历史数据
func (h *hData) DataHistory(ctx context.Context, req *apiv1.DataHistoryReq) (res *apiv1.DataHistoryRes, err error) {
	res = &apiv1.DataHistoryRes{}
	print(req.FrozenType)
	in := model.MqttDatabaseGetHistoryIn{
		Token:      "633",
		TimeType:   req.TimeType,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		TimeSpan:   req.TimeSpan,
		FrozenType: req.FrozenType,
		Body:       req.Body,
	}
_:
	gconv.Struct(req, &in)
	//var valList model.MqttDatabaseGetHistoryInBody
	//valList.Dev = req.Body.Dev
	//// 如果为空表示查询所有变量
	//if len(req.Body.Body) == 0 {
	//	var vals []string
	//	// 查询对应的所有变量
	//	for _, dev := range entity.devices.DevList {
	//		if dev.DevGuid == valList.Dev {
	//			for _, val := range dev.YCVal {
	//				vals = append(vals, val.Name)
	//			}
	//			for _, val := range dev.YXVal {
	//				vals = append(vals, val.Name)
	//			}
	//			valList.Body = vals
	//			break
	//		}
	//	}
	//} else {
	//	valList.Body = req.Body.Body
	//}
	//in.Body = valList
	out, err := service.Mqtt().MqttDatabaseGetHistory(ctx, in)
	if err != nil {
		return nil, err
	}
_:
	gconv.Struct(out.Body, &res)
	return res, nil
}

// DataHistoryN 按上N条访问历史数据
func (h *hData) DataHistoryN(ctx context.Context, req *apiv1.DataHistoryNReq) (res *apiv1.DataHistoryRes, err error) {
	res = &apiv1.DataHistoryRes{}
	print(req.FrozenType)
	in := model.MqttDatabaseGetHistoryInN{
		Token:      "633",
		TimeStamp:  "2022-03-29T23:32:49.223+0800",
		Dev:        req.Dev,
		UpperN:     req.UpperN,
		FrozenType: req.FrozenType,
		Body:       req.Body,
	}
_:
	gconv.Struct(req, &in)
	in.Body = []string{"AMax_phsA", "BMax_phsA", "CMax_phsA", "A_phsA", "B_phsA", "C_phsA",
		"EnvHum", "PhV_phsA", "PhV_phsB", "PhV_phsC", "Tmp", "PTUV_Open_Op_phsA", "PTUV_Open_Op_phsB", "PTUV_Open_Op_phsC",
		"A_neut", "AMax_neut", "Max_Tmp", "Max_TmpValue"}
	out, err := service.Mqtt().MqttDatabaseGetHistoryN(ctx, in)
	if err != nil {
		return nil, err
	}
_:
	gconv.Struct(out.Body, &res)
	return res, nil
}

func (h *hData) DataRealtime(ctx context.Context, req *apiv1.DataRealtimeReq) (res *apiv1.DataRealtimeRes, err error) {
	res = &apiv1.DataRealtimeRes{}
	topic := consts.Publish_realtime_data_get
	in := "{ \"token\":\"123\",\"timestamp\":\"2022-03-21T09:30:08.230+0800\",\"body\":[{\"dev\":\"" + req.Dev + "\", \"totalcall\":\"" + req.TotalCall + "\", \"body\":[\"" + req.Attribute + "\"]}]}"

_:
	// 传入 请求消息体和 请求主题
	fmt.Println(in)
	out, err := service.Mqtt().MqttDatabaseGetRealtime(ctx, topic, in)
	if err != nil {
		return nil, err
	}
_:
	fmt.Println("handler:")
	fmt.Println(out)
	gconv.Struct(out.Body[0], &res)
	fmt.Println(res)
	return res, nil
}

func (h *hData) DataTopo(ctx context.Context, req *apiv1.DataTopoReq) (res *apiv1.DataTopoRes, err error) {
	res = &apiv1.DataTopoRes{}
	topic := consts.Publish_register_get
	in := "{ \"token\":\"123\",\"timestamp\":\"2022-03-21T09:30:08.230+0800\",\"body\":[]}]}"
	modelName := []string{"LTU", "Switch"}
_:
	// 传入 请求消息体和 请求主题
	fmt.Println(in)
	out, err := service.Mqtt().MqttDatabaseGetTopo(ctx, topic, in, modelName)
	if err != nil {
		return nil, err
	}
_:
	fmt.Println("handler:")
	fmt.Println(out)
	gconv.Struct(out, &res)
	fmt.Println(res)
	return res, nil
}
