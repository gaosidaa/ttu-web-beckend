package service

import (
	"context"
	"encoding/json"
	"fmt"
	"ttu-backend/internal/model"
)

type (
	// sMqtt is service struct of module Mqtt.
	sMqtt struct{}
)

var (
	// insMqtt is the instance of service Mqtt.
	insMqtt = sMqtt{}
)

// Mqtt returns the interface of Mqtt service.
func Mqtt() *sMqtt {
	return &insMqtt
}

func (s *sMqtt) MqttDatabaseGetHistory(ctx context.Context, in model.MqttDatabaseGetHistoryIn) (out model.MqttDatabaseGetHistoryOut, err error) {
	reqJson, err := json.Marshal(in)
	if err != nil {
		return out, err
	}
	fmt.Println(string(reqJson))
	//todo:此处请求数据中心“6.13.1按时间查询”接口,替换下面的模拟函数
	out = Simulator()
	return out, nil
}

func Simulator() (res model.MqttDatabaseGetHistoryOut) {
	return model.MqttDatabaseGetHistoryOut{
		Token:     "633",
		Timestamp: "2016-06-02T10:11:18.067+0800",
		Body: model.MqttDatabaseGetHistoryOutBody{
			Dev: "ADC_frozen_fa0ad9d877ba7f41",
			Body: []model.MqttDatabaseGetHistoryOutBodyBody{
				{
					Timestamp:       "2019-12-25T01:51:42.000+0800",
					TimeStartGather: "2019-12-25T01:51:42.000+0800",
					TimeEndGather:   "2019-12-25T01:51:42.000+0800",
					AdditionalCheck: "88888888",
					Body: []model.MqttDatabaseGetHistoryOutBodyBodyBody{
						{
							Name: "AvSupWh",
							Val:  "2.2565",
						},
						{
							Name: "AvRtlWh",
							Val:  "3.2565",
						},
					},
				},
				{
					Timestamp:       "2019-12-26T01:51:42.000+0800",
					TimeStartGather: "2019-12-26T01:51:42.000+0800",
					TimeEndGather:   "2019-12-26T01:51:42.000+0800",
					AdditionalCheck: "88888888",
					Body: []model.MqttDatabaseGetHistoryOutBodyBodyBody{
						{
							Name: "AvSupWh",
							Val:  "5.2565",
						},
						{
							Name: "AvRtlWh",
							Val:  "6.2565",
						},
					},
				},
			},
		},
	}
}
