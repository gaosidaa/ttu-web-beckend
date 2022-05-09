package service

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
)

var mqttClient MQTT.Client

func init() {
	var b bool
	b, mqttClient = connMQTT(g.Cfg("mqtt_config").MustGet(nil, "mqtt.MQTTLAN").String())
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

// 连接MQTT服务
func connMQTT(broker string) (bool, MQTT.Client) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)

	mc := MQTT.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		return false, mc
	}

	mc.Subscribe(g.Cfg("mqtt_config").MustGet(nil, "sub_topics.Received_history_data_get").String(), g.Cfg("mqtt_config").MustGet(nil, "mqtt.MQTTQos").Bytes()[0], hisCallBackFunc)
	mc.Subscribe(g.Cfg("mqtt_config").MustGet(nil, "sub_topics.Received_realtime_data_get").String(), g.Cfg("mqtt_config").MustGet(nil, "mqtt.MQTTQos").Bytes()[0], realtimeCallBackFunc)
	mc.Subscribe(g.Cfg("mqtt_config").MustGet(nil, "sub_topics.Received_register_get").String(), g.Cfg("mqtt_config").MustGet(nil, "mqtt.MQTTQos").Bytes()[0], registerCallBackFunc)
	mc.Subscribe(g.Cfg("mqtt_config").MustGet(nil, "sub_topics.Received_alarm_data_get").String(), g.Cfg("mqtt_config").MustGet(nil, "mqtt.MQTTQos").Bytes()[0], alarmCallBackFunc)
	mc.Subscribe(g.Cfg("mqtt_config").MustGet(nil, "sub_topics.Received_setParams").String(), g.Cfg("mqtt_config").MustGet(nil, "mqtt.MQTTQos").Bytes()[0], setParamsCallBackFunc)
	mc.Subscribe(g.Cfg("mqtt_config").MustGet(nil, "sub_topics.Received_getParams").String(), g.Cfg("mqtt_config").MustGet(nil, "mqtt.MQTTQos").Bytes()[0], getParamsCallBackFunc)
	return true, mc
}

// 发布消息
func publish(topic string, str string) {
	// 发送
	mqttClient.Publish(topic, g.Cfg("mqtt_config").MustGet(nil, "mqtt.MQTTQos").Bytes()[0], true, str)
}
