package entity

type Devices struct {
	DevList []Device `json:"body"`
}

// Device 设备信息及对应端口和变量数组
type Device struct {
	Guid          string `json:"guid"`
	DevGuid       string `json:"dev"`
	FrozenDevGuid string
	AppName       string      `json:"appname"`
	IsReport      string      `json:"isReport"`
	ManuID        string      `json:"manuID"`
	Port          DevicePort  `json:"devDesc"`
	Model         DeviceModel `json:"model"`
	YXVal         []DeviceVal `json:"yx"`
	YCVal         []DeviceVal `json:"yc"`
}

type DevicePort struct {
	Port        string `json:"port"`
	Address     string `json:"addr"`
	Description string `json:"desc"`
}

type DeviceVal struct {
	TimeStamp  string `json:"timestamp"`
	Name       string `json:"name"`
	FrozenType string `json:"frozentype"`
	Val        string `json:"val"`
	Quality    string `json:"quality"`
}
