package handler

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"ttu-backend/apiv1"
	"ttu-backend/internal/model"
	"ttu-backend/internal/service"
)

var (
	Data = hData{}
)

type hData struct{}

func (h *hData) DataHistory(ctx context.Context, req *apiv1.DataHistoryReq) (res *apiv1.DataHistoryRes, err error) {
	res = &apiv1.DataHistoryRes{}
	in := model.MqttDatabaseGetHistoryIn{
		Token:      "633",
		TimeType:   "timestartgather",
		StartTime:  "2019-12-25T00:00:00.727+0800",
		EndTime:    "2019-12-25T23:42:30.727+0800",
		TimeSpan:   "5",
		FrozenType: "min",
		Body: model.MqttDatabaseGetHistoryInBody{
			Dev:  "ADC_frozen_fa0ad9d877ba7f41",
			Body: []string{},
		},
	}
_:
	gconv.Struct(req, &in)
	in.Body.Dev = req.Dev
	out, err := service.Mqtt().MqttDatabaseGetHistory(ctx, in)
	if err != nil {
		return nil, err
	}
_:
	gconv.Struct(out.Body, &res)
	return res, nil
}
