package service

import (
	"context"
	"encoding/json"
	"fmt"
	"ttu-backend/internal/consts"
	"ttu-backend/internal/model"
	"ttu-backend/internal/model/entity"
)

// 全局设备数组与全局模型数组
var DeviceList entity.Devices
var ModelList entity.Models

// 消息全局变量
var realtimeResTmp model.MqttDatabaseGetRealtimeOut
var topoRes model.MqttDatabaseGetTopoOut
var historyRestmp model.MqttDatabaseGetHistoryOut
var alarmRes model.MqttDataBaseGetAlarmOut
var getConfigRes model.MqttDataBaseGetConfigOut
var setConfigRes model.MqttDataBaseSetConfigOut
var faultWaveform model.BaseFaultWaveformOut
var (
	historyChan       = make(chan model.MqttDatabaseGetHistoryOut)
	realtimeChan      = make(chan model.MqttDatabaseGetRealtimeOut)
	topoChan          = make(chan model.MqttDatabaseGetTopoOut)
	alarmChan         = make(chan model.MqttDataBaseGetAlarmOut)
	faultWaveformChan = make(chan model.BaseFaultWaveformOut)
	getConfigChan     = make(chan model.MqttDataBaseGetConfigOut)
	setConfigChan     = make(chan model.MqttDataBaseSetConfigOut)
)

type (
	// sMqtt is service struct of module Mqtt.
	sMqtt struct{}
)

var (
	// insMqtt is the instance of service Mqtt.
	insMqtt = sMqtt{}
)

// Mqtt returns the interface of Mqtt service.
func Mqtt() *sMqtt {
	return &insMqtt
}

// MqttDatabaseGetHistory 按时间段获取历史数据
func (s *sMqtt) MqttDatabaseGetHistory(ctx context.Context, in model.MqttDatabaseGetHistoryIn) (out model.MqttDatabaseGetHistoryOut, err error) {
	// 消息转换处理
	for _, dev := range DeviceList.DevList {
		if dev.DevGuid == in.Body.Dev {
			// 转换为冻结模型guid
			in.Body.Dev = dev.FrozenDevGuid
			// 变量为空则查询所有变量
			if len(in.Body.Body) == 0 {
				var vals []string
				for _, val := range dev.YCVal {
					vals = append(vals, val.Name)
				}
				for _, val := range dev.YXVal {
					vals = append(vals, val.Name)
				}
				in.Body.Body = vals
			}
			break
		}
	}

	reqJson, err := json.Marshal(in)
	if err != nil {
		return out, err
	}
	fmt.Println(string(reqJson))
	publish(consts.Publish_history_data_get, string(reqJson))

	out = <-historyChan
	return out, nil
}

// MqttDatabaseGetHistoryN 按上N条获取历史数据
func (s *sMqtt) MqttDatabaseGetHistoryN(ctx context.Context, in model.MqttDatabaseGetHistoryInN) (out model.MqttDatabaseGetHistoryOut, err error) {
	// 消息转换处理
	for _, dev := range DeviceList.DevList {
		if dev.DevGuid == in.Dev {
			// 转换为冻结模型guid
			in.Dev = dev.FrozenDevGuid
			// 变量为空则查询所有变量
			if len(in.Body) == 0 {
				var vals []string
				for _, val := range dev.YCVal {
					vals = append(vals, val.Name)
				}
				for _, val := range dev.YXVal {
					vals = append(vals, val.Name)
				}
				in.Body = vals
			}
			break
		}
	}
	reqJson, err := json.Marshal(in)
	if err != nil {
		return out, err
	}
	fmt.Println(string(reqJson))
	publish(consts.Publish_history_data_get, string(reqJson))

	out = <-historyChan
	return out, nil
}

func (s *sMqtt) MqttDatabaseGetRealtime(ctx context.Context, in string) (out model.MqttDatabaseGetRealtimeOut, err error) {
	if err != nil {
		return out, err
	}
	// 发布消息
	fmt.Println("发布的消息 " + in)
	publish(consts.Publish_realtime_data_get, in)
	out = <-realtimeChan
	return out, nil
}

func (s *sMqtt) MqttDatabaseGetTopo(ctx context.Context, in string, modelName []string) (out model.MqttDatabaseGetTopoOut, err error) {
	if err != nil {
		return out, err
	}
	// 发布消息
	fmt.Println("发布的消息 " + in)
	publish(consts.Publish_register_get, in)
	// 对消息体进行解析
	return TopoHandler(<-topoChan, modelName), nil
}

func TopoHandler(cont model.MqttDatabaseGetTopoOut, modelName []string) (res model.MqttDatabaseGetTopoOut) {
	var info model.MqttDatabaseGetTopoOut
	info = cont
	info.Body = nil
	// 筛选App管理的模型
	for _, value := range cont.Body {
		for _, eachModel := range modelName {
			if value.Model == eachModel {
				info.Body = append(info.Body, value)
				// 数据记录到全局变量
				for _, d := range value.Body {
					var dev entity.Device
					var port entity.DevicePort
					var mod entity.DeviceModel
					dev.Guid = d.Guid
					dev.DevGuid = d.Dev

					port.Address = d.Addr
					port.Port = value.Port
					port.Description = d.Desc
					dev.Port = port

					mod.Name = value.Model
					dev.Model = mod

					dev.ManuID = d.ManuID
					dev.IsReport = d.Isreport
					DeviceList.DevList = append(DeviceList.DevList, dev)
				}
				// 记录Frozen_guid
			} else if value.Model == eachModel+"_frozen" {
				for _, d := range value.Body {
					for j, dev := range DeviceList.DevList {
						if dev.Port.Address == d.Addr {
							DeviceList.DevList[j].FrozenDevGuid = d.Dev
						}
					}
				}
			}
		}
	}

	return info
}

func GetFrozenDev(Dev string) string {
	for _, dev := range DeviceList.DevList {
		if dev.DevGuid == Dev {
			return dev.FrozenDevGuid
		}
	}
	return ""
}
func GetModelNameByDev(Dev string) string {
	for _, dev := range DeviceList.DevList {
		if dev.DevGuid == Dev || dev.FrozenDevGuid == Dev {
			return dev.Model.Name
		}
	}
	return ""
}
func GetModelParasByDev(Dev string) []string {
	paras := make([]string, 0)
	modelName := GetModelNameByDev(Dev)
	for k, _ := range consts.ModelConfig[modelName].Mapping {
		paras = append(paras, k)
	}
	return paras
}
func init() {

}
