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

func (h *hData) DataHistory(ctx context.Context, req *apiv1.DataHistoryReq) (res *apiv1.DataHistoryRes, err error) {
	res = &apiv1.DataHistoryRes{}
	in := model.MqttDatabaseGetHistoryIn{
		Token:      "633",
		TimeType:   "timestartgather",
		StartTime:  "2019-12-25T00:00:00.727+0800",
		EndTime:    "2019-12-25T23:42:30.727+0800",
		TimeSpan:   "5",
		FrozenType: "min",
		Body: model.MqttDatabaseGetHistoryInBody{
			Dev:  "ADC_frozen_fa0ad9d877ba7f41",
			Body: []string{},
		},
	}
_:
	gconv.Struct(req, &in)
	in.Body.Dev = req.Dev
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
	in := "{ \"token\":\"123\",\"timestamp\":\"2022-03-21T09:30:08.230+0800\",\"body\":[{\"dev\":\"" + req.Dev + "\", \"totalcall\":\"0\", \"body\":[\"" + req.Attribute + "\"]}]}"
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
