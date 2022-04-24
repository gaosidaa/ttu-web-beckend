package handler

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"ttu-backend/apiv1"
	"ttu-backend/internal/model"
	"ttu-backend/internal/service"
)

var (
	Base = hBase{}
)

type hBase struct{}

func (h *hBase) DeviceList(ctx context.Context, req *apiv1.DeviceListReq) (res *apiv1.DeviceListRes, err error) {
	in := model.BaseDeviceListIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseDeviceList(ctx, in)
	res = &apiv1.DeviceListRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}
func (h *hBase) Realtime(ctx context.Context, req *apiv1.RealtimeReq) (res *apiv1.RealtimeRes, err error) {
	in := model.BaseRealtimeIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseRealtime(ctx, in)
	res = &apiv1.RealtimeRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}
func (h *hBase) Record(ctx context.Context, req *apiv1.RecordReq) (res *apiv1.RecordRes, err error) {
	in := model.BaseRecordIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseRecord(ctx, in)
	res = &apiv1.RecordRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}
func (h *hBase) Alarm(ctx context.Context, req *apiv1.AlarmReq) (res *apiv1.AlarmRes, err error) {

	return res, nil
}
