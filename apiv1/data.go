package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

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
