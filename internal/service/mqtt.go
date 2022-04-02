package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"ttu-backend/internal/consts"
	"ttu-backend/internal/model"
	"ttu-backend/internal/model/entity"
)

// 全局设备数组与全局模型数组
var DeviceList entity.Devices
var ModelList entity.Models
var FrozenModelList entity.Models

// 消息全局变量
var initResTmp string
var realtimeResTmp string
var topoRes string
var historyRestmp model.MqttDatabaseGetHistoryOut

//var realtimeResHum []byte
//var TopoRes []byte

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

func (s *sMqtt) MqttInit(ctx context.Context, in model.EmptyIn) (out model.MqttInitOut, err error) {
	reqJson, err := json.Marshal(in)
	if err != nil {
		return out, err
	}
	publish(consts.Publish_device_get, string(reqJson))
	err = json.Unmarshal([]byte(initResTmp), &out)
	if err != nil {
		return model.MqttInitOut{}, err
	}
	return out, err
}

func (s *sMqtt) MqttDatabaseGetHistory(ctx context.Context, in model.MqttDatabaseGetHistoryIn) (out model.MqttDatabaseGetHistoryOut, err error) {
	reqJson, err := json.Marshal(in)
	if err != nil {
		return out, err
	}
	fmt.Println(string(reqJson))
	publish(consts.Publish_history_data_get, string(reqJson))
	out = historyRestmp
	return out, nil
}

func (s *sMqtt) MqttDatabaseGetRealtime(ctx context.Context, topic string, in string) (out model.MqttDatabaseGetRealtimeOut, err error) {
	if err != nil {
		return out, err
	}
	// 发布消息
	fmt.Println("发布的消息 " + topic)
	publish(topic, in)
	fmt.Println(realtimeResTmp)
	// 对消息体进行解析
	time.Sleep(time.Second)
	return RealtimeSimulator(realtimeResTmp), nil
}

func (s *sMqtt) MqttDatabaseGetTopo(ctx context.Context, topic string, in string, modelName []string) (out model.MqttDatabaseGetTopoOut, err error) {
	if err != nil {
		return out, err
	}
	// 发布消息
	fmt.Println("发布的消息 " + topic)
	publish(topic, in)
	fmt.Println(topoRes)
	// 对消息体进行解析
	time.Sleep(time.Second)
	return TopoSimulator(topoRes, modelName), nil

}

func TopoSimulator(cont string, modelName []string) (res model.MqttDatabaseGetTopoOut) {
	var model1 model.MqttDatabaseGetTopoOut
	var model2 model.MqttDatabaseGetTopoOut
	fmt.Println(json.Unmarshal([]byte(cont), &model1))
	model2 = model1
	model2.Body = nil
	for _, value := range model1.Body{
		for _, eachModel := range modelName{
			if (value.Model == eachModel){
				model2.Body = append(model2.Body, value)
			}
		}
	}
	fmt.Println(model2)
	return model2
}

func RealtimeSimulator(cont string) (res model.MqttDatabaseGetRealtimeOut) {
	var model1 model.MqttDatabaseGetRealtimeOut
	fmt.Println(json.Unmarshal([]byte(cont), &model1))
	fmt.Println(model1)
	return model1
	/*return model.MqttDatabaseGetRealtimeOut{
		Token:     model1.Token,
		Timestamp: model1.Timestamp,
		Body: []model.MqttDatabaseGetRealtimeOutBody{
			{
				Dev: model1.Body[0].Dev,
				Body: []model.MqttDatabaseGetRealtimeOutBodyBody{
					{
						Name:      model1.Body[0].Body[0].Name,
						Val:       model1.Body[0].Body[0].Val,
						Quality:   model1.Body[0].Body[0].Quality,
						Timestamp: model1.Body[0].Body[0].Timestamp,
					},
				},
			},
		},
	}*/
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
