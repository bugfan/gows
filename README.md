# golang 千万级别WebSocket消息推送服务

## 功能
1. 维护所有在线链接
2. 可以根据指定的key去快速定位链接进行推送消息
3. 使用chan处理了线程安全问题
4. 发送/接受消息量为千万级别 (unix内核瓶颈为千万) 
5. 增加了client.html websocket前端代码

## 使用
1. 先执行 `go get github.com/gorilla/websocket`
2. 启动程序 `go run main.go`
