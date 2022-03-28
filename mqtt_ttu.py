import paho.mqtt.client as mqtt

return_realtime = "{\"token\":\"123\",\"timestamp\":\"2019-03-01T09:30:08.230+0800\",\"body\":[{\"dev\":\"ADC_guid\"," \
             "\"body\":[{\"name\":\"PhV_phsA\",\"val\":\"220.331\",\"quality\":\"1\"," \
             "\"timestamp\":\"2019-11-22T14:00:08.230+0800\"}]}]} "

return_topo = "{\"token\": \"234\",\"timestamp\": \"2019-03-01T09:30:09.230+0800\",\"body\":[{\"model\": \"MCCB\",\"port\": \"RS485-2\",\"body\":[{\"guid\": \"guid\",\"dev\": \"MCCB_ guid\",\"addr\": \"1\",\"appname\": \"MCCB_collector\",\"desc\": \"\",\"manuID\":\"1234\",\"isreport\":	\"1\"},{\"guid\": \"guid\",\"dev\": \"MCCB_ guid\",\"addr\": \"2\",\"appname\": \"MCCB_collector\",\"desc\": \"\",\"manuID\":\"1234\",\"isreport\":	\"1\"}]},{\"model\": \"RCD\",\"port\": \"RS485-1\",\"body\":[{\"guid\": \"guid\",\"dev\": \"RCD_ guid\",\"addr\": \"010000000001\",\"appname\": \"RCD_collector\",\"desc\": \"\",\"manuID\":	\"1234\",\"isreport\":\"1\"}]}]}"

def on_connect(client, userdata, flags, rc):
    print("Connected with result code: " + str(rc))

def on_message(client, userdata, msg):
    print(msg.topic + " " + str(msg.payload))
    if str(msg.topic) == "TestApp/get/request/database/realtime":
        client.publish('database/get/response/TestApp/realtime', payload=return_realtime, qos=0)
        print("sent")
    if str(msg.topic) == "TestApp/get/request/database/register":
        client.publish('database/get/response/TestApp/register', payload=return_topo, qos=0)
        print("sent")

client = mqtt.Client()
client.on_connect = on_connect
client.on_message = on_message
client.connect('mnifdv.cn', 1883, 600) # 600为keepalive的时间间隔
#client.publish('TestApp/get/request/database/realtime', payload='amazing', qos=0)
client.subscribe("TestApp/get/request/database/realtime")
client.subscribe("TestApp/get/request/database/register")
client.loop_forever() # 保持连接