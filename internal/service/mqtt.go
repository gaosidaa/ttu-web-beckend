package service

import (
	"context"
	"encoding/json"
	"fmt"
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
var realtimeResTmp model.MqttDatabaseGetRealtimeOut
var topoRes model.MqttDatabaseGetTopoOut
var historyRestmp model.MqttDatabaseGetHistoryOut
var (
	historyChan  = make(chan model.MqttDatabaseGetHistoryOut)
	realtimeChan = make(chan model.MqttDatabaseGetRealtimeOut)
	topoChan     = make(chan model.MqttDatabaseGetTopoOut)
)

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

// MqttDatabaseGetHistory 按时间段获取历史数据
func (s *sMqtt) MqttDatabaseGetHistory(ctx context.Context, in model.MqttDatabaseGetHistoryIn) (out model.MqttDatabaseGetHistoryOut, err error) {
	// 如果为空表示查询所有变量
	if len(in.Body.Body) == 0 {
		var vals []string
		// 查询对应的所有变量
		for _, dev := range DeviceList.DevList {
			if dev.DevGuid == in.Body.Dev {
				for _, val := range dev.YCVal {
					vals = append(vals, val.Name)
				}
				for _, val := range dev.YXVal {
					vals = append(vals, val.Name)
				}
				in.Body.Body = vals
				break
			}
		}
	}

	reqJson, err := json.Marshal(in)
	if err != nil {
		return out, err
	}
	fmt.Println(string(reqJson))
	publish(consts.Publish_history_data_get, string(reqJson))

	out = <-historyChan
	return out, nil
}

// MqttDatabaseGetHistoryN 按上N条获取历史数据
func (s *sMqtt) MqttDatabaseGetHistoryN(ctx context.Context, in model.MqttDatabaseGetHistoryInN) (out model.MqttDatabaseGetHistoryOut, err error) {
	// 如果为空表示查询所有变量
	if len(in.Body) == 0 {
		var vals []string
		// 查询对应的所有变量
		for _, dev := range DeviceList.DevList {
			if dev.DevGuid == in.Dev {
				for _, val := range dev.YCVal {
					vals = append(vals, val.Name)
				}
				for _, val := range dev.YXVal {
					vals = append(vals, val.Name)
				}
				in.Body = vals
				break
			}
		}
	}
	reqJson, err := json.Marshal(in)
	if err != nil {
		return out, err
	}
	fmt.Println(string(reqJson))
	publish(consts.Publish_history_data_get, string(reqJson))

	out = <-historyChan
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
	realtime, ok := <-realtimeChan
	for !ok {
		realtime, ok = <-realtimeChan
	}
	//message := <-realtimeChan
	// 对消息体进行解析
	return realtime, nil
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
	message := <-topoChan
	return TopoSimulator(message, modelName), nil
}

func TopoSimulator(cont model.MqttDatabaseGetTopoOut, modelName []string) (res model.MqttDatabaseGetTopoOut) {
	var model2 model.MqttDatabaseGetTopoOut
	model2 = cont
	model2.Body = nil
	for _, value := range cont.Body {
		for _, eachModel := range modelName {
			if value.Model == eachModel {
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
