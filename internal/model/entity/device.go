package entity

type Devices struct {
	DevList []Device `json:"body"`
}

// Device 设备信息及对应端口和变量数组
type Device struct {
	Guid          string `json:"guid"`
	DevGuid       string `json:"dev"`
	FrozenGuid    string
	FrozenDevGuid string
	AppName       string      `json:"appname"`
	IsReport      string      `json:"isReport"`
	ManuID        string      `json:"manuID"`
	Init          DevicePort  `json:"devDesc"`
	Model         DeviceModel `json:"model"`
	FrozenModel   DeviceModel
	Connection    DeviceConnection `json:"port"`
	YXVal         []DeviceVal      `json:"yx"`
	YCVal         []DeviceVal      `json:"yc"`
}

type DevicePort struct {
	Port        string `json:"port"`
	Address     string `json:"addr"`
	Description string `json:"desc"`
	ModelName   string `json:"model"`
}

type DeviceConnection struct {
	Baud    string `json:"Baud"`
	Bit     string `json:"bit"`
	StopBit string `json:"stop_bit"`
	Parity  string `json:"Parity"`
}

type DeviceVal struct {
	TimeStamp  string `json:"timestamp"`
	Name       string `json:"name"`
	FrozenType string `json:"frozentype"`
	Val        string `json:"val"`
	Quality    string `json:"quality"`
}
