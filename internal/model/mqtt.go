package model

// EmptyIn 空消息体
type EmptyIn struct {
	Token     string `json:"token"`
	Timestamp string `json:"timestamp"`
	Body      string `json:"body"`
}

// MqttInitOut 挂接设备列表消息体
type MqttInitOut struct {
	Token     string          `json:"token"`
	Timestamp string          `json:"timestamp"`
	Body      MqttInitOutBody `json:"body"`
}

// MqttInitOutBody App所有注册设备
type MqttInitOutBody struct {
	App  string              `json:"appname"`
	Body MqttInitOutBodyBody `json:"body"`
}

// MqttInitOutBodyBody 每个设备的具体信息
type MqttInitOutBodyBody struct {
	Dev     string `json:"dev"`
	Model   string `json:"model"`
	Port    string `json:"port"`
	Address string `json:"addr"`
	Desc    string `json:"desc"`
}

// MqttDatabaseGetHistoryIn 按时间段的历史数据
type MqttDatabaseGetHistoryIn struct {
	Token      string                       `json:"token"`
	TimeType   string                       `json:"time_type"`
	StartTime  string                       `json:"start_time"`
	EndTime    string                       `json:"end_time"`
	TimeSpan   string                       `json:"time_span"`
	FrozenType string                       `json:"frozentype"`
	Body       MqttDatabaseGetHistoryInBody `json:"body"`
}

type MqttDatabaseGetHistoryInBody struct {
	Dev  string   `json:"dev"`
	Body []string `json:"body"`
}

// MqttDatabaseGetHistoryInN 按上N条的历史数据
type MqttDatabaseGetHistoryInN struct {
	Token      string   `json:"token"`
	TimeStamp  string   `json:"timestamp"`
	Dev        string   `json:"dev"`
	FrozenType string   `json:"frozentype"`
	UpperN     string   `json:"upperN"`
	Body       []string `json:"body"`
}

type MqttDatabaseGetHistoryOut struct {
	Token     string                        `json:"token"`
	Timestamp string                        `json:"timestamp"`
	Body      MqttDatabaseGetHistoryOutBody `json:"body"`
}
type MqttDatabaseGetHistoryOutBody struct {
	Dev  string                              `json:"dev"`
	Body []MqttDatabaseGetHistoryOutBodyBody `json:"body"`
}
type MqttDatabaseGetHistoryOutBodyBody struct {
	Timestamp       string                                  `json:"timestamp"`
	TimeStartGather string                                  `json:"timestartgather"`
	TimeEndGather   string                                  `json:"timeendgather"`
	AdditionalCheck string                                  `json:"additionalcheck"`
	Body            []MqttDatabaseGetHistoryOutBodyBodyBody `json:"body"`
}
type MqttDatabaseGetHistoryOutBodyBodyBody struct {
	Name string `json:"name"`
	Val  string `json:"val"`
}

type MqttDatabaseGetRealtimeIn struct {
	Token     string                          `json:"token"`
	Timestamp string                          `json:"timestamp"`
	Body      []MqttDatabaseGetRealtimeInBody `json:"body"`
}

type MqttDatabaseGetRealtimeInBody struct {
	Dev       string   `json:"dev"`
	Totalcall string   `json:"totalcall"`
	Body      []string `json:"body"`
}

type MqttDatabaseGetRealtimeOut struct {
	Token     string                           `json:"token"`
	Timestamp string                           `json:"timestamp"`
	Body      []MqttDatabaseGetRealtimeOutBody `json:"body"`
}

type MqttDatabaseGetRealtimeOutBody struct {
	Dev  string                               `json:"dev"`
	Body []MqttDatabaseGetRealtimeOutBodyBody `json:"body"`
}
type MqttDatabaseGetRealtimeOutBodyBody struct {
	Name      string `json:"name"`
	Val       string `json:"val"`
	Quality   string `json:"quality"`
	Timestamp string `json:"timestamp"`
}

// 台区设备列表
type MqttDatabaseGetTopoIn struct {
	Token     string   `json:"token"`
	Timestamp string   `json:"timestamp"`
	Body      []string `json:"body"`
}

type MqttDatabaseGetTopoOut struct {
	Token     string                       `json:"token"`
	Timestamp string                       `json:"timestamp"`
	Body      []MqttDatabaseGetTopoOutBody `json:"body"`
}

type MqttDatabaseGetTopoOutBody struct {
	Model string                           `json:"model"`
	Port  string                           `json:"port"`
	Body  []MqttDatabaseGetTopoOutBodyBody `json:"body"`
}
type MqttDatabaseGetTopoOutBodyBody struct {
	Guid      string `json:"guid"`
	Dev       string `json:"dev"`
	Addr      string `json:"addr"`
	Desc      string `json:"desc"`
	ManuID    string `json:"manuID"`
	Isreport  string `json:"isreport"`
	NodeID    string `json:"nodeID"`
	ProductID string `json:"productID"`
}

// 告警
type MqttDataBaseGetAlarmIn struct {
	Token     string                       `json:"token"`
	Time_type string                       `json:"time_type"`
	StartTime string                       `json:"start_time"`
	EndTime   string                       `json:"end_time"`
	SourType  string                       `json:"SourType"`
	Body      []MqttDataBaseGetAlarmInBody `json:"body"`
}

type MqttDataBaseGetAlarmInBody struct {
	Model    string   `json:"model"`
	Totaldev string   `json:"totaldev"`
	Dev      []string `json:"dev"`
}

type MqttDataBaseGetAlarmOut struct {
	Token     string                        `json:"token"`
	Timestamp string                        `json:"timestamp"`
	Body      []MqttDataBaseGetAlarmOutBody `json:"body"`
}

type MqttDataBaseGetAlarmOutBody struct {
	Appname         string          `json:"appname"`
	SourType        string          `json:"SourType"`
	Model           string          `json:"model"`
	Dev             string          `json:"dev"`
	Event           string          `json:"event"`
	Timestamp       string          `json:"timestamp"`
	Timestartgather string          `json:"timestartgather"`
	Timeendgather   string          `json:"timeendgather"`
	Startimestamp   string          `json:"starttimestamp"`
	Endtimestamp    string          `json:"endtimestamp"`
	HappenSrc       string          `json:"HappenSrc"`
	IsNeedRpt       string          `json:"IsNeedRpt"`
	occurnum        string          `json:"occurnum"`
	EventLevel      string          `json:"EventLevel"`
	RptStatus       []RptStatusBody `json:"RptStatus"`
	Data            string          `json:"data"`
	Extdata         []ExtdataBody   `json:"extdata"`
}

type RptStatusBody struct {
	Net1 string `json:"Net-1"`
	Eth1 string `json:"Eth-1"`
}

type ExtdataBody struct {
	Name      string `json:"name"`
	Val       string `json:"val"`
	Timestamp string `json:"timestamp"`
}

// get参数
type MqttDataBaseGetConfigIn struct {
	Dev string `json:"dev"`
}

type MqttDataBaseGetConfigOut struct {
	Dev  string                        `json:"dev"`
	Body []MqttDatabaseSetConfigInBody `json:"body"`
}

// set
type MqttDataBaseSetConfigIn struct {
	Dev  string                        `json:"dev"`
	Body []MqttDatabaseSetConfigInBody `json:"body"`
}

type MqttDatabaseSetConfigInBody struct {
	Name string `json:"name" title:"变量名称"`
	Val  string `json:"val" title:"变量数值"`
}
type MqttDataBaseSetConfigOut struct {
	Dev    string `json:"dev"`
	Status string `json:"status"`
}
