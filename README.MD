# ttu-backend

我的思路：
api的data.go文件里只需要必须的参数，并定义返回前端的参数列表
比如实时数据：设备号+数据属性名
台区拓扑：分成两种，一种是全部获取。一种是分类型获取，需要model名

model的mqtt.go:
定义mqtt 向ttu发送请求时的json消息格式 （可跳过，直接到handler的部分）
定义mqtt ttu返回响应的json格式

handler的data.go
将前端的请求参数字符串拼接到mqtt请求消息体中，用于service的mqtt.go的发送

service的mqtt.go
定义全局变量 