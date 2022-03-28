package service

import (
	"context"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
	"ttu-backend/internal/model"
)

// 订阅回调
// 可以传一个全局变量进去修改
// 多设计几个全局变量

var realtimeResTmp string

//var realtimeResHum []byte
//var TopoRes []byte

func subCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	// 这一步是将消息进行解析并return

	realtimeResTmp = string(msg.Payload())

}

// 连接MQTT服务
func connMQTT(broker string) (bool, MQTT.Client) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)

	mc := MQTT.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		return false, mc
	}

	return true, mc
}

// 订阅某一个主题
func subscribe(topic string) {
	// sub的用户名和密码
	b, mc := connMQTT("sinpower.3322.org:11883")
	// sinpower.3322.org:11883
	// mnifdv.cn:1883
	if !b {
		fmt.Println("sub connMQTT failed")
		return
	}
	mc.Subscribe(topic, 0x00, subCallBackFunc)
	fmt.Println(topic)
}

// 发布消息
func publish(topic string, str string) {
	b, mc := connMQTT("sinpower.3322.org:11883")
	if !b {
		fmt.Println("pub connMQTT failed")
		return
	}
	// 发送
	mc.Publish(topic, 0x00, true, str)
}

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
	/* 参数类型
	"dev": "string",
	 "end_time": "2006-01-02 15:04:05",
	 "frozen_type": "string",
	 "start_time": "2006-01-02 14:04:05",
	 "time_span": "string",
	 "time_type": "string"
	*/

	out = Simulator()
	return out, nil
}

func (s *sMqtt) MqttDatabaseGetRealtime(ctx context.Context, topic string, in string) (out model.MqttDatabaseGetRealtimeOut, err error) {
	if err != nil {
		return out, err
	}
	// 订阅某主题
	subscribe("database/get/response/TestApp/realtime")
	// 发布消息
	fmt.Println("发布的消息 " + topic)
	publish(topic, in)
	fmt.Println(realtimeResTmp)
	// 对消息体进行解析
	time.Sleep(time.Second)
	return RealtimeSimulator(realtimeResTmp), nil
}

func (s *sMqtt) MqttDatabaseGetTopo(ctx context.Context, topic string, in string) (out model.MqttDatabaseGetTopoOut, err error) {
	if err != nil {
		return out, err
	}
	// 订阅某主题
	subscribe("database/get/response/TestApp/register")
	// 发布消息
	fmt.Println("发布的消息 " + topic)
	publish(topic, in)
	fmt.Println(realtimeResTmp)
	// 对消息体进行解析
	time.Sleep(time.Second)
	return TopoSimulator(realtimeResTmp), nil

}

func TopoSimulator(cont string) (res model.MqttDatabaseGetTopoOut) {
	var model1 model.MqttDatabaseGetTopoOut
	fmt.Println(json.Unmarshal([]byte(cont), &model1))
	fmt.Println(model1)
	return model1
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
