package model

type MqttConfig struct {
	Mqtt      Mqtt      `json:"mqtt"`
	PubTopics PubTopics `json:"pub_topics"`
	SubTopics SubTopics `json:"sub_topics"`
}
type Mqtt struct {
	MQTTLAN string `json:"MQTTLAN"`
	MQTTQos int    `json:"MQTTQos"`
}
type PubTopics struct {
	PublishRegisterGet     string `json:"Publish_register_get"`
	PublishRealtimeDataGet string `json:"Publish_realtime_data_get"`
	PublishHistoryDataGet  string `json:"Publish_history_data_get"`
	PublishAlarmDataGet    string `json:"Publish_alarm_data_get"`
	PublishGetParams       string `json:"Publish_getParams"`
	PublishSetParams       string `json:"Publish_setParams"`
}
type SubTopics struct {
	ReceivedRegisterGet     string `json:"Received_register_get"`
	ReceivedRealtimeDataGet string `json:"Received_realtime_data_get"`
	ReceivedHistoryDataGet  string `json:"Received_history_data_get"`
	ReceivedAlarmDataGet    string `json:"Received_alarm_data_get"`
	ReceivedGetParams       string `json:"Received_getParams"`
	ReceivedSetParams       string `json:"Received_setParams"`
}
