package consts

const (
	OpenAPITitle       = `IoT Platform`
	OpenAPIDescription = `物联网平台 `
	//OpenAPITagNameChat = `Chat`
	OpenAPITagNameUser = `User`

	Publish_register_get      = "TestApp/get/request/database/register"
	Publish_realtime_data_get = "TestApp/get/request/database/realtime"
	Publish_history_data_get  = "TestApp/get/request/database/history"
	Publish_get_MonMax        = "TestApp/set/request/database/DataStoreDepth"
	Publish_initData          = "TestApp/action/request/database/InitData"

	Peceived_initData = "database/action/response/TestApp/InitData"

	Received_register_get      = "database/get/response/TestApp/register"
	Received_realtime_data_get = "database/get/response/TestApp/realtime"
	Received_history_data_get  = "database/get/response/TestApp/history"

	Message                = "{\n  \"token\": \"123\", \n  \"timestamp\": \"2019-03-01T09:30:08.230+0800\",\n  \"body\": []\n}"
	Realtime_message       = "{ \"token\":\"123\",\"timestamp\":\"2022-03-21T09:30:08.230+0800\",\"body\":[{\"dev\":\"LTU_bb38620dc4e710b0\", \"totalcall\":\"0\", \"body\":[\"Tmp\"]}]}"
	History_period_message = "{\n\"token\":\"123\",\n\"time_type\":\"timestartgather\",\n\"start_time\":\"2022-03-23T16:00:00.727+0800\",\n\"end_time\":\"2023-03-23T17:00:30.727+0800\",\n\"time_span\":\"5\",\n\"frozentype\":\"min\",\n\"body\":{\n\"dev\":\"LTU_bb38620dc4e710b0\",\n\"body\":[\"PhV_phsB\"]\n}\n}\n"
	History_lastN_message  = "{\"token\": \"123\", 	\"timestamp\": \"2022-03-01T09:30:08.230+0800\", 	\"dev\": \"LTU_bb38620dc4e710b0\", 	\"frozentype\": \"min\", 	\"upperN\": \"10\", 	\"body\": [] }"
	// 数据区初始化
	Init_message = "{\"token\":\"123\",\"timestamp\":\"2019-03-01T09:30:08.230+0800\",\"model\":\"LTU\",\"Action\":\"frozenData\",\"data\":\"All\"}"
)
