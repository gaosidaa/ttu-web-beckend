package model

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
