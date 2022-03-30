package service

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"ttu-backend/internal/consts"
)

// 各订阅接口的回调函数和MQTT连接，发送功能函数
// 通过全局变量传递信息

func deviceCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	// 这一步是将消息进行解析并return
	initResTmp = string(msg.Payload())
}

func schemaCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	// 这一步是将消息进行解析并return

	realtimeResTmp = string(msg.Payload())
}

func modelCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	// 这一步是将消息进行解析并return

	realtimeResTmp = string(msg.Payload())
}

func guidCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	// 这一步是将消息进行解析并return

	realtimeResTmp = string(msg.Payload())
}

func registerCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	// 这一步是将消息进行解析并return

	realtimeResTmp = string(msg.Payload())
}

func realtimeCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	// 这一步是将消息进行解析并return

	realtimeResTmp = string(msg.Payload())
}

func hisCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	// 这一步是将消息进行解析并return

	err := json.Unmarshal(msg.Payload(), &historyRestmp)
	if err != nil {
		fmt.Println(err)
	}
}

func initCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
}

func emptyMessage() (str string) {
	return ""
}

// 连接MQTT服务
func connMQTT(broker string) (bool, MQTT.Client) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)

	mc := MQTT.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		return false, mc
	}

	mc.Subscribe(consts.Received_history_data_get, consts.MQTTQos, hisCallBackFunc)
	mc.Subscribe(consts.Received_guid_get, consts.MQTTQos, guidCallBackFunc)
	mc.Subscribe(consts.Received_device_get, consts.MQTTQos, deviceCallBackFunc)
	mc.Subscribe(consts.Received_realtime_data_get, consts.MQTTQos, realtimeCallBackFunc)
	mc.Subscribe(consts.Received_register_get, consts.MQTTQos, registerCallBackFunc)
	mc.Subscribe(consts.Received_modelschema_get, consts.MQTTQos, schemaCallBackFunc)
	mc.Subscribe(consts.Received_model_get, consts.MQTTQos, modelCallBackFunc)
	mc.Subscribe(consts.Received_initData, consts.MQTTQos, initCallBackFunc)
	return true, mc
}

// 发布消息
func publish(topic string, str string) {
	b, mc := connMQTT(consts.MQTTBroker)
	if !b {
		fmt.Println("pub connMQTT failed")
		return
	}
	// 发送
	mc.Publish(topic, consts.MQTTQos, true, str)
}
