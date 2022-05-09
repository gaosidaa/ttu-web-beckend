package service

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"ttu-backend/internal/consts"
)

var mqttClient MQTT.Client

func init() {
	var b bool
	b, mqttClient = connMQTT(consts.MQTTLAN)
	if !b {
		fmt.Println("pub connMQTT failed")
		return
	}
}

// 各订阅接口的回调函数和MQTT连接，发送功能函数
// 通过全局变量传递信息
func getParamsCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	err := json.Unmarshal(msg.Payload(), &getConfigRes)
	if err != nil {
		fmt.Println(err)
	}
	getConfigChan <- getConfigRes
}

func setParamsCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	err := json.Unmarshal(msg.Payload(), &setConfigRes)
	if err != nil {
		fmt.Println(err)
	}
	setConfigChan <- setConfigRes

}

func alarmCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic())
	err := json.Unmarshal(msg.Payload(), &alarmRes)
	if err != nil {
		fmt.Println(err)
	}
	alarmChan <- alarmRes
}
func deviceCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
}

func schemaCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
}

func modelCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
}

func guidCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
}

func registerCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	err := json.Unmarshal(msg.Payload(), &topoRes)
	if err != nil {
		fmt.Println(err)
	}
	topoChan <- topoRes
}

func realtimeCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	err := json.Unmarshal(msg.Payload(), &realtimeResTmp)
	if err != nil {
		fmt.Println(err)
	}
	realtimeChan <- realtimeResTmp
}

func hisCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	// 这一步是将消息进行解析并return
	err := json.Unmarshal(msg.Payload(), &historyRestmp)
	if err != nil {
		fmt.Println(err)
	}
	historyChan <- historyRestmp
}

func initCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	fmt.Printf("Connection lost: %v \n", err)
	client.Connect()
}

// 连接MQTT服务
func connMQTT(broker string) (bool, MQTT.Client) {
	opts := MQTT.NewClientOptions()
	opts.OnConnectionLost = connectLostHandler
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
	mc.Subscribe(consts.Received_alarm_data_get, consts.MQTTQos, alarmCallBackFunc)
	mc.Subscribe(consts.Received_setParams, consts.MQTTQos, setParamsCallBackFunc)
	mc.Subscribe(consts.Received_getParams, consts.MQTTQos, getParamsCallBackFunc)
	mc.Subscribe(consts.Received_history_data_get, consts.MQTTQos, hisCallBackFunc)
	return true, mc
}

// 发布消息
func publish(topic string, str string) {
	// 发送
	mqttClient.Publish(topic, consts.MQTTQos, true, str)
}
