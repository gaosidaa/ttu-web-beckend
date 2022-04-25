package consts

const (
	uid        = "app"
	MQTTBroker = "mqtt://sinpower.3322.org:11883"
	//MQTTLAN    = "mqtt://192.168.2.33:1883"
	MQTTLAN = "mqtt://47.110.134.175:8183"
	MQTTQos = 0x0

	Publish_device_get        = uid + "/get/request/database/vT"
	Publish_modelschema_get   = uid + "/get/request/database/modelschema"
	Publish_model_get         = uid + "/get/request/database/model"
	Publish_guid_get          = uid + "/get/request/database/guid"
	Publish_register_get      = uid + "/get/request/database/register"
	Publish_realtime_data_get = uid + "/get/request/database/realtime"
	Publish_history_data_get  = uid + "/get/request/database/history"
	Publish_get_MonMax        = uid + "/set/request/database/DataStoreDepth"
	Publish_initData          = uid + "/action/request/database/InitData"

	Received_device_get        = "database/get/response/" + uid + "/vT"
	Received_modelschema_get   = "database/get/response/" + uid + "/modelschema"
	Received_model_get         = "database/get/response/" + uid + "/model"
	Received_guid_get          = "database/get/response/" + uid + "/guid"
	Received_initData          = "database/action/response/" + uid + "/InitData"
	Received_register_get      = "database/get/response/" + uid + "/register"
	Received_realtime_data_get = "database/get/response/" + uid + "/realtime"
	Received_history_data_get  = "database/get/response/" + uid + "/history"

	Message                = "{\n  \"token\": \"123\", \n  \"timestamp\": \"2019-03-01T09:30:08.230+0800\",\n  \"body\": []\n}"
	Realtime_message       = "{ \"token\":\"123\",\"timestamp\":\"2022-03-21T09:30:08.230+0800\",\"body\":[{\"dev\":\"LTU_bb38620dc4e710b0\", \"totalcall\":\"0\", \"body\":[\"Tmp\"]}]}"
	History_period_message = "{\n\"token\":\"123\",\n\"time_type\":\"timestartgather\",\n\"start_time\":\"2022-03-23T16:00:00.727+0800\",\n\"end_time\":\"2023-03-23T17:00:30.727+0800\",\n\"time_span\":\"5\",\n\"frozentype\":\"min\",\n\"body\":{\n\"dev\":\"LTU_bb38620dc4e710b0\",\n\"body\":[\"PhV_phsB\"]\n}\n}\n"
	History_lastN_message  = "{\"token\": \"123\", 	\"timestamp\": \"2022-03-01T09:30:08.230+0800\", 	\"dev\": \"LTU_bb38620dc4e710b0\", 	\"frozentype\": \"min\", 	\"upperN\": \"10\", 	\"body\": [] }"
	// 数据区初始化
	Init_message = "{\"token\":\"123\",\"timestamp\":\"2019-03-01T09:30:08.230+0800\",\"model\":\"LTU\",\"Action\":\"frozenData\",\"data\":\"All\"}"
)

var ModelConfig = map[string]struct {
	Name    string
	Mapping map[string]string
}{
	"LTU": {
		Name: "线路终端",
		Mapping: map[string]string{
			"AMax_phsA": "InMax",
			"BMax_phsA": "InAvg",
			"A_phsA":    "Ia",
			"B_phsA":    "Ib",
			"C_phsA":    "Ic",
			"PhV_phsA":  "Ua",
			"PhV_phsB":  "Ub",
			"PhV_phsC":  "Uc",
			"Tmp":       "T",
			"EnvHum":    "H",
		},
	},
	"Switch": {
		Name: "智能开关",
		Mapping: map[string]string{
			"AMax_phsA": "InMax",
			"BMax_phsA": "InAvg",
			"A_phsA":    "Ia",
			"B_phsA":    "Ib",
			"C_phsA":    "Ic",
			"PhV_phsA":  "Ua",
			"PhV_phsB":  "Ub",
			"PhV_phsC":  "Uc",
			"Tmp":       "T",
			"EnvHum":    "H",
		},
	},
}
