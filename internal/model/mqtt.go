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

// MqttDatabaseGetHistoryIn 历史数据
type MqttDatabaseGetHistoryIn struct {
	Token      string   `json:"token"`
	TimeType   string   `json:"time_type"`
	StartTime  string   `json:"start_time"`
	EndTime    string   `json:"end_time"`
	Timestamp  string   `json:"timestamp"`
	Dev        string   `json:"dev"`
	UpperN     string   `json:"upperN"`
	TimeSpan   string   `json:"time_span"`
	FrozenType string   `json:"frozentype"`
	Body       []string `json:"body"`
	//Body       MqttDatabaseGetHistoryInBody `json:"body"`
}
type MqttDatabaseGetHistoryInBody struct {
	Dev  string   `json:"dev"`
	Body []string `json:"body"`
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

// 实时数据,另一个角度，将

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
