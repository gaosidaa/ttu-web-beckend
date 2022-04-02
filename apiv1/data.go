package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type DataInitReq struct {
	g.Meta `path:"/data/init" method:"post" tags:"Data" summary:"初始数据加载"`
}

type DataRes struct {
}

type DataHistoryReq struct {
	g.Meta     `path:"/data/history" method:"post" tags:"Data" summary:"获取历史数据"`
	Dev        string `json:"dev"         v:"required" title:"设备的唯一地址" dc:"由模型和guid组合而成，是数据中心中设备的唯一地址"`
	TimeType   string `json:"time_type"   v:"required" title:"筛选的事件类型" dc:"在timestartgather、timeendgather、timestamp中选择,分别为采集开始时间、采集结束时间、冻结时间"`
	StartTime  string `json:"start_time"  v:"required|date-format:2006-01-02 15:04:05" title:"开始时间" dc:""`
	EndTime    string `json:"end_time"    v:"required|date-format:2006-01-02 15:04:05" title:"结束时间" dc:""`
	TimeSpan   string `json:"time_span"   v:"required" title:"筛选周期" dc:"与frozentype对应，分别为day、mon、hour、min，若frozentype选择SchFroz则单位为min"`
	FrozenType string `json:"frozen_type" v:"required" title:"冻结类型" dc:"在day、billday、mon、hour、min、realtime、SchFroz中选择，分别表示日冻结、结算日、月冻结、小时冻结、分钟冻结、实时冻结、实时数据转冻结"`
}

type DataHistoryRes struct {
	Dev  string               `json:"dev"         title:"设备的唯一地址"`
	Body []DataHistoryResBody `json:"body"                title:"数据体"`
}
type DataHistoryResBody struct {
	Timestamp       *gtime.Time              `json:"time"                title:"冻结时间"`
	TimeStartGather *gtime.Time              `json:"time_start_gather"   title:"采集开始时间"`
	TimeEndGather   *gtime.Time              `json:"time_end_gather"     title:"采集结束时间"`
	Body            []DataHistoryResBodyBody `json:"body"                title:"数据体"`
}
type DataHistoryResBodyBody struct {
	Name string `json:"name"                title:"变量名"`
	Val  string `json:"val"                 title:"变量值"`
}

// 实时数据

type DataRealtimeReq struct {
	g.Meta    `path:"/data/realtime" method:"post" tags:"Data" summary:"获取某一设备的某一属性实时数据"`
	Dev       string `json:"dev"         v:"required" title:"设备的唯一地址" dc:"由模型和guid组合而成，是数据中心中设备的唯一地址"`
	Attribute string `json:"attribute" v:"required" title:"设备的属性" dc:"想要获取的某一项设备属性"`
	TotalCall string `json:"totalCall" v:"required" title:"要全部属性还是单个属性。1 ：所有属性，不用再加属性。 0：加上需要的属性"`
}

type DataRealtimeRes struct {
	Dev  string                `json:"dev"         title:"设备的唯一地址"`
	Body []DataRealtimeResBody `json:"body"                title:"数据体"`
}

type DataRealtimeResBody struct {
	Name      string      `json:"name"                title:"变量名"`
	Val       string      `json:"val"                 title:"变量值"`
	Quality   string      `json:"quality"             title:"质量"`
	Timestamp *gtime.Time `json:"time"                title:"数据时间"`
}

// 台区拓扑数据

type DataTopoReq struct {
	g.Meta `path:"/data/register" method:"get" tags:"Data" summary:"获取某一台区的所有设备列表"`
	Model      []string      `json:"model"                title:"输入需要的模型名"`
}

type DataTopoRes struct {
	Body []DataTopoResBody `json:"body"                title:"数据体"`
}

type DataTopoResBody struct {
	Model string                `json:"model" title:"模型名称"`
	Port  string                `json:"port" title:"端口号"`
	Body  []DataTopoResBodyBody `json:"body" title:"该模型的所有设备"`
}

type DataTopoResBodyBody struct {
	Guid     string `json:"guid"`
	Dev      string `json:"dev"`
	Addr     string `json:"addr"`
	Appname  string `json:"appname"`
	Desc     string `json:"desc"`
	ManuID   string `json:"manuID"`
	Isreport string `json:"isreport"`
}
