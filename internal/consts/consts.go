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
	Publish_alarm_data_get    = uid + "/get/request/database/SOE"
	Publish_get_MonMax        = uid + "/set/request/database/DataStoreDepth"
	Publish_initData          = uid + "/action/request/database/InitData"
	Publish_getParams         = uid + "/get/request/TestApp/params"
	Publish_setParams         = uid + "/set/request/TestApp/params"

	Received_device_get        = "database/get/response/" + uid + "/vT"
	Received_modelschema_get   = "database/get/response/" + uid + "/modelschema"
	Received_model_get         = "database/get/response/" + uid + "/model"
	Received_guid_get          = "database/get/response/" + uid + "/guid"
	Received_initData          = "database/action/response/" + uid + "/InitData"
	Received_register_get      = "database/get/response/" + uid + "/register"
	Received_realtime_data_get = "database/get/response/" + uid + "/realtime"
	Received_history_data_get  = "database/get/response/" + uid + "/history"
	Received_alarm_data_get    = "database/get/response/" + uid + "/SOE"
	Received_getParams         = "TestApp/get/response/" + uid + "/params"
	Received_setParams         = "TestApp/set/response/" + uid + "/params"

	Message                = "{\n  \"token\": \"123\", \n  \"timestamp\": \"2019-03-01T09:30:08.230+0800\",\n  \"body\": []\n}"
	Realtime_message       = "{ \"token\":\"123\",\"timestamp\":\"2022-03-21T09:30:08.230+0800\",\"body\":[{\"dev\":\"LTU_bb38620dc4e710b0\", \"totalcall\":\"0\", \"body\":[\"Tmp\"]}]}"
	History_period_message = "{\n\"token\":\"123\",\n\"time_type\":\"timestartgather\",\n\"start_time\":\"2022-03-23T16:00:00.727+0800\",\n\"end_time\":\"2023-03-23T17:00:30.727+0800\",\n\"time_span\":\"5\",\n\"frozentype\":\"min\",\n\"body\":{\n\"dev\":\"LTU_bb38620dc4e710b0\",\n\"body\":[\"PhV_phsB\"]\n}\n}\n"
	History_lastN_message  = "{\"token\": \"123\", 	\"timestamp\": \"2022-03-01T09:30:08.230+0800\", 	\"dev\": \"LTU_bb38620dc4e710b0\", 	\"frozentype\": \"min\", 	\"upperN\": \"10\", 	\"body\": [] }"
	// 数据区初始化
	Init_message = "{\"token\":\"123\",\"timestamp\":\"2019-03-01T09:30:08.230+0800\",\"model\":\"LTU\",\"Action\":\"frozenData\",\"data\":\"All\"}"
	// 告警消息
	Alarm_message     = "{\"token\": \"123\",\"time_type\":\"timestartgather\",\"start_time\": \"2020-02-24T14:08:34.055+0800\",\"end_time\": \"2020-02-24T16:38:34.055+0800\",\"SourType\": \"104\", \"body\": [{\"model\": \"MultiMeter_frozen\",\"totaldev\": \"0\",\"dev\":[\"MultiMeter_frozen_7fb2132a153c212b\"]}]}"
	GetConfig_message = "{\"dev\":\"123456\"}"
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

var AlarmDict = map[string]string{
	"IndN":              "开入量N",
	"PTOV_Open_Op_phsA": "A相断相",
	"PTUV_Loss_Op_phsA": "A相失压",
	"PTUV_Op_phsA":      "A相欠压",
	"PTOV_Op_phsA":      "A相过压",
	"PTOC_Op_phsA":      "A相过流",
	"PTUC_Op_phsA":      "A相失流",
	"PTOC_Ovld_Op_phsA": "A相过载",
	"PTOC_Hvld_Op_phsA": "A相重载",
	"PTUC_Open_Op_phsA": "A相断流",
	"RDIR_dirPhsA":      "A相潮流反向",
	"PTOC_Dly_Op_phsA":  "A相短路短延时",
	"PIOC_Op_phsA":      "A相短路瞬时保护",
	"PTLK_Op_phsA":      "A相剩余电流保护",
	"PTUV_lim_Op_phsB":  "B相电压越下限",
	"PTOV_lim_Op_phsB":  "B相电压越上限",
	"PTOV_Open_Op_phsB": "B相断相",
	"PTUV_Loss_Op_phsB": "B相失压",
	"PTUV_Op_phsB":      "B相欠压",
	"PTOV_Op_phsB":      "B相过压",
	"PTOC_Op_phsB":      "B相过流",
	"PTUC_Op_phsB":      "B相失流",
	"PTOC_Ovld_Op_phsB": "B相过载",
	"PTOC_Hvld_Op_phsB": "B相重载",
	"PTUC_Open_Op_phsB": "B相断流",
	"RDIR_dirPhsB":      "B相潮流反向",
	"PTOC_Dly_Op_phsB":  "B相短路短延时保护",
	"PIOC_Op_phsB":      "B相短路瞬时保护",
	"PTLK_Op_phsB":      "B相剩余电流保护",
	"PTUV_lim_Op_phsC":  "C相电压越下限",
	"PTOV_lim_Op_phsC":  "C相电压越上限",
	"PTOV_Open_Op_phsC": "C相断相",
	"PTUV_Loss_Op_phsC": "C相失压",
	"PTUV_Op_phsC":      "C相欠压",
	"PTOV_Op_phsC":      "C相过压",
	"PTOC_Op_phsC":      "C相过流",
	"PTUC_Op_phsC":      "C相失流",
	"PTOC_Ovld_Op_phsC": "C相过载",
	"PTOC_Hvld_Op_phsC": "C相重载",
	"PTUC_Open_Op_phsC": "C相断流",
	"RDIR_dirPhsC":      "C相潮流反向",
	"PTOC_Dly_Op_phsC":  "C相短路短延时保护",
	"PIOC_Op_phsC":      "C相短路瞬时保护",
	"PTLK_Op_phsC":      "C相剩余电流保护",
	"SeqAAlm":           "电压逆相序",
	"SeqVAlm":           "电流逆相序",
	"ImbAAlm":           "电压不平衡",
	"ImbVAlm":           "电流不平衡",
	"PwrSupAlm":         "辅助电源失电",
	"PwrOffAlm":         "停电",
	"PwrOnAlm":          "上电",
	"DmdUpAlm":          "需量超限",
	"DoorAlm":           "门磁报警",
	"FloodAlm":          "水浸报警",
	"SmokAlm":           "烟感报警",
	"ShockAlm":          "振动报警",
	"Oilchrom":          "油色谱",
	"CSWI_ pos":         "开关机械状态",
	"TEST_pos":          "断路器自检状态",
	"ImbLAlm":           "负荷不平衡",
}
