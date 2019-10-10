# golang WebSocket消息推送服务

## 功能
1. 维护所有在线链接
2. 可以根据指定的key去快速定位链接进行推送消息
3. 使用chan处理了线程安全问题
4. 发送/接受消息量为千万级别 (unix内核有瓶颈) 
5. 增加了client.html websocket前端代码

## 使用
1. 先执行 `git clone https://github.com/bugfan/gows.git`
2. 进入gows目录 `cd xxxx/gows`
3. 启动程序 `go run main.go -debug=true -p=9000` // -debug参数指定是否将服务端日志打印出来,-p指定监听端口

## 容器化

1. 切换到此目录 ，执行 `docker build -t YourTag .`
