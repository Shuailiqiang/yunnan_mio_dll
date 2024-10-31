package main

import (
	"YunNanDll/dllInvoke"
	"YunNanDll/parameter"
	"YunNanDll/sysClient"
	"YunNanDll/util"
	"YunNanDll/web"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	// 配置日志信息
	file, logErr := os.OpenFile("./log/dll-"+time.Now().Format(time.DateOnly)+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if logErr != nil {
		log.Fatal("Failed to open log file:", logErr)
	}
	// 设置日志的输出目标为文件和控制台
	multiWriter := io.MultiWriter(file, os.Stdout) // 添加控制台输出
	log.SetOutput(multiWriter)
	// 设置日志的输出目标和格式

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("init log success")

	//本地加载动态库
	log.Println("load dll CHSInterfaceYn.dll ")
	dllInvoke.LoadDirDll()

	//本地加载配置文件
	util.InitConfig("YunNanDll.yaml")
	log.Println("load success")

	//本地初始化动态库
	log.Println("init dll CHSInterfaceYn.dll ")
	dllInvoke.Init_Local(parameter.Conf.ApiParam.Fixmedins_code, parameter.Conf.ApiParam.Infosyscode, parameter.Conf.ApiParam.Infosyssign, parameter.Conf.ApiParam.Url)

	//启动web服务器
	log.Println("Starting WebServer")
	web.Start()

	//启动系统托盘
	log.Println("Starting application")
	sysClient.Start()
	log.Println("Starting success")

}
