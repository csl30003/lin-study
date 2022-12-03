package main

import (
	"server/server"
	"server/tool"
)

//
// main
//  @Description: 主函数
//
func main() {
	tool.RabbitMQSimpleLogger.PublishSimpleLogger(tool.INFO, "项目启动")
	server.Start()
}
