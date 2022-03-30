package entity

type Models struct {
	ModList []DeviceModel `json:"body"`
}

// DeviceModel 设备对应的物模型
type DeviceModel struct {
	Name    string     `json:"model"`
	ValNum  string     `json:"num"`
	YXModel []ValModel `json:"yx"`
	YCModel []ValModel `json:"yc"`
}

type ValModel struct {
	Name       string `json:"name"`
	Address    string `json:"addr"`
	Type       string `json:"type"`
	Unit       string `json:"bit"`
	DeadZone   string `json:"deadzone"`
	Ratio      string `json:"ratio"`
	IsReport   string `json:"isReport"`
	UserDefine string `json:"userdefine"`
}
