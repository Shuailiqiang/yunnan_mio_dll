package sysClient

import (
	"YunNanDll/dllInvoke"
	"YunNanDll/parameter"
	"YunNanDll/web"
	"github.com/getlantern/systray"
	"io/ioutil"
	"log"
	"os"
)

func Start() {
	systray.Run(onReady, onExit)
}

/**
 * 初始化时时调用函数指针
 */
func onReady() {

	systray.SetTitle("西安动态库调用服务")
	systray.SetIcon(getIconData())
	// 添加一个退出菜单项
	quitMenuItem := systray.AddMenuItem("退出", "退出程序")
	initMenuItem := systray.AddMenuItem("重新初始化", "重新初始化")
	// 监听退出菜单项的点击事件
	go func() {
		<-quitMenuItem.ClickedCh
		log.Println("Quitting application...")
		// 关闭服务器
		web.ShutdownServer()
		log.Println("Quitting application success")
		systray.Quit()

	}()

	go func() {
		<-initMenuItem.ClickedCh
		//本地加载动态库
		log.Println("load dll CHSInterfaceYn.dll ")
		dllInvoke.LoadDirDll()

		//本地初始化动态库
		log.Println("init dll CHSInterfaceYn.dll ")
		dllInvoke.Init_Local(parameter.Conf.ApiParam.Fixmedins_code, parameter.Conf.ApiParam.Infosyscode, parameter.Conf.ApiParam.Infosyssign, parameter.Conf.ApiParam.Url)
	}()
}

func onExit() {
	// 停止web服务放在右键菜单监听事件
}

// 获取
func getIconData() []byte {
	// 打开图标文件
	file, err := os.Open("icon.ico")
	if err != nil {
		log.Println("无法打开图标文件:", err)
		return nil
	}

	// 读取图标文件的内容
	iconData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("无法读取图标文件:", err)
		return nil
	}
	defer file.Close()

	return iconData
}
