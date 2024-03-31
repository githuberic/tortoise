# 顺序时钟服务器
时钟服务器,每秒一次向客户端发送当前时间

* 一次处理一个连接(gr_clock_v1_main.go)
* 并发处理多个连接(gr_clock_v2_main.go)