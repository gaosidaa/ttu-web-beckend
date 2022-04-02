package handler

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"ttu-backend/apiv1"
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

// DataHistory 按上N条访问历史数据
func (h *hData) DataHistory(ctx context.Context, req *apiv1.DataHistoryReq) (res *apiv1.DataHistoryRes, err error) {
	res = &apiv1.DataHistoryRes{}
	in := model.MqttDatabaseGetHistoryIn{
		Token: "633",
		//TimeType:   "timestartgather",
		//StartTime:  "2022-03-29T23:05:50.296+0800",
		//EndTime:    "2022-03-29T23:16:50.296+0800",
		//TimeSpan:   "5",
		Timestamp:  "2022-03-29T23:32:49.223+0800",
		Dev:        "LTU_frozen_3b2ebaac2c6bc90a",
		UpperN:     "10",
		FrozenType: "min",
		Body:       []string{"PhV_phsA"},
		//Body: model.MqttDatabaseGetHistoryInBody{
		//	Dev:  "LTU_frozen_3b2ebaac2c6bc90a",
		//	Body: []string{"PhV_phsA"},
		//},
	}
_:
	gconv.Struct(req, &in)
	//in.Body.Dev = req.Dev
	out, err := service.Mqtt().MqttDatabaseGetHistory(ctx, in)
	if err != nil {
		return nil, err
	}
_:
	gconv.Struct(out.Body, &res)
	return res, nil
}

func (h *hData) DataRealtime(ctx context.Context, req *apiv1.DataRealtimeReq) (res *apiv1.DataRealtimeRes, err error) {
	res = &apiv1.DataRealtimeRes{}
	topic := "TestApp/get/request/database/realtime"
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
	topic := "TestApp/get/request/database/register"
	in := "{ \"token\":\"123\",\"timestamp\":\"2022-03-21T09:30:08.230+0800\",\"body\":[]}]}"
_:
	// 传入 请求消息体和 请求主题
	fmt.Println(in)
	out, err := service.Mqtt().MqttDatabaseGetTopo(ctx, topic, in)
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
