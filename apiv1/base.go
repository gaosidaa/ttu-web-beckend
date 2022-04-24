package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceListReq 获取设备列表
type DeviceListReq struct {
	g.Meta `path:"/base/deviceList" method:"get" tags:"Base" summary:"获取设备列表"`
}

type DeviceListRes struct {
	Station string               `json:"station" title:"台区名称"`
	Group   []DeviceListResGroup `json:"group"                title:"设备分组"`
}

type DeviceListResGroup struct {
	Name   string                     `json:"name" title:"设备分组名称"`
	Total  int                        `json:"total" title:"组内设备总数"`
	Online int                        `json:"online" title:"组内设备在线数"`
	Icon   string                     `json:"icon" title:"设备分组图标"`
	Device []DeviceListResGroupDevice `json:"device" title:"设备分组内的设备"`
}

type DeviceListResGroupDevice struct {
	Dev    string `json:"dev"    title:"设备的唯一地址"`
	Name   string `json:"name" title:"设备名称"`
	Online bool   `json:"online" title:"在线状态"`
	Badge  int    `json:"badge" title:"角标数值"`
}

// RealtimeReq 获取实时数据
type RealtimeReq struct {
	g.Meta `path:"/base/realtime" method:"post" tags:"Base" summary:"获取实时数据"`
	Dev    string `json:"dev"         v:"required"  title:"设备的唯一地址"`
}
type RealtimeRes RecordRes

// RecordReq 获取记录数据
type RecordReq struct {
	g.Meta         `path:"/base/record" method:"post" tags:"Base" summary:"获取记录数据"`
	Dev            string `json:"dev"         v:"required"  title:"设备的唯一地址"`
	StartTime      string `json:"start_time"  title:"开始时间" dc:""`
	EndTime        string `json:"end_time"    title:"结束时间" dc:""`
	TimeSpanUnit   string `json:"time_span_unit"   title:"记录间隔周期" dc:"可选值：min、hour、day、mon"`
	TimeSpanNumber string `json:"time_span_number"   title:"记录间隔数" dc:""`
	Limit          int    `json:"limit"       title:"返回条数限制" dc:""`
}

type RecordRes struct {
	Dev    string            `json:"dev"         title:"设备的唯一地址"`
	Count  int               `json:"count"         title:"记录数据条数"`
	Record []RecordResRecord `json:"record"                title:"记录数据"`
}
type RecordResRecord struct {
	Timestamp *gtime.Time `json:"time"              title:"时间"`
	InAvg     float32     `json:"In_Avg"            title:"平均漏电电流"`
	InMax     float32     `json:"In_Max"            title:"最大漏电电流"`
	Ia        float32     `json:"Ia"                title:"A相电流"`
	Ib        float32     `json:"Ib"                title:"B相电流"`
	Ic        float32     `json:"Ic"                title:"C相电流"`
	Ua        float32     `json:"Ua"                title:"A相电压"`
	Ub        float32     `json:"Ub"                title:"B相电压"`
	Uc        float32     `json:"Uc"                title:"C相电压"`
	Tn        float32     `json:"Tn"                title:"接点最高温度"`
	T         float32     `json:"T"                 title:"环境温度"`
	H         float32     `json:"H"                 title:"环境适度"`
}

// AlarmReq 获取告警事件
type AlarmReq struct {
	g.Meta    `path:"/base/alarm" method:"post" tags:"Base" summary:"获取告警事件"`
	StartTime string `json:"start_time"  title:"开始时间" dc:""`
	EndTime   string `json:"end_time"    title:"结束时间" dc:""`
}

type AlarmRes struct {
	Alarm []AlarmResAlarm `json:"alarm"                title:"告警事件"`
}
type AlarmResAlarm struct {
	Timestamp *gtime.Time `json:"time"           title:"时间"`
	AlarmType string      `json:"alarm_type"     title:"事件类型"`
	Remark    string      `json:"remark"         title:"事件说明"`
	Status    string      `json:"status"         title:"事件状态"`
}
