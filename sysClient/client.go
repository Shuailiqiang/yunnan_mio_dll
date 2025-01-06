package sysClient

import (
	"YunNanDll/dllInvoke"
	"YunNanDll/parameter"
	"YunNanDll/web"
	"bytes"
	"github.com/getlantern/systray"
	"io"
	"log"
	"os"
	"os/exec"
)

func Start() {
	systray.Run(onReady, onExit)
}

/**
 * 初始化时时调用函数指针
 */
func onReady() {

	systray.SetTitle("动态库调用服务")
	systray.SetIcon(getIconData())
	// 添加一个退出菜单项
	quitMenuItem := systray.AddMenuItem("退出", "退出程序")
	initMenuItem := systray.AddMenuItem("重新初始化", "重新初始化")
	ecTokenMenuItem := systray.AddMenuItem("读电子凭证", "读电子凭证")
	callExeMenuItem := systray.AddMenuItem("调用EXE", "调用EXE")
	// 监听退出菜单项的点击事件
	go func() {
		for {
			<-quitMenuItem.ClickedCh
			log.Println("Quitting application...")
			// 关闭服务器
			web.ShutdownServer()
			log.Println("Quitting application success")
			systray.Quit()
		}
	}()

	go func() {
		<-initMenuItem.ClickedCh
		for {
			//本地加载动态库
			log.Println("initMenuItem Clicked")
			dllInvoke.LoadDirDll()

			//本地初始化动态库
			log.Println("init dll CHSInterfaceYn.dll ")
			local, err := dllInvoke.Init_Local(parameter.Conf)
			if err != nil {
				log.Println("动态库初始化失败!" + err.Error())
			}
			log.Println(local)
		}
	}()

	go func() {
		for {
			<-ecTokenMenuItem.ClickedCh
			//本地加载动态库
			log.Println("ecTokenMenuItem Clicked")
			log.Println("parameter.InitStatus ===> ", parameter.InitStatus)
			if !parameter.InitStatus {
				log.Println("动态库还未初始化成功,请重新初始化!")
				return
			}

			var param = "{\n\t\t\t\"infno\": \"1191\",\n\t\t\t\"msgid\": \"H53011200683202412121831078729\",\n\t\t\t\"mdtrtarea_admvs\": \"530000\",\n\t\t\t\"insuplc_admdvs\": \"\",\n\t\t\t\"recer_sys_code\": \"mbs\",\n\t\t\t\"dev_no\": \"\",\n\t\t\t\"dev_safe_info\": \"\",\n\t\t\t\"cainfo\": \"\",\n\t\t\t\"signtype\": \"\",\n\t\t\t\"infver\": \"1.0\",\n\t\t\t\"opter_type\": \"1\",\n\t\t\t\"opter\": \"neu\",\n\t\t\t\"opter_name\": \"neu\",\n\t\t\t\"inf_time\": \"2024-12-12 18:31:07\",\n\t\t\t\"fixmedins_code\": \"H53011200683\",\n\t\t\t\"fixmedins_name\": \"医院\",\n\t\t\t\"sign_no\": \"\",\n\t\t\t\"input\": {\n\t\t\t\"data\": {\n\t\t\t\t\"mdtrt_cert_type\": \"01\",\n\t\t\t\t\t\"cardtype\": \"0\",\n\t\t\t\t\t\"businesstype\": \"01101\",\n\t\t\t\t\t\"operatorid\": \"test001\",\n\t\t\t\t\t\"operatorname\": \"超级管理员\",\n\t\t\t\t\t\"officeid\": \"32760\",\n\t\t\t\t\t\"deviceType\": \"\",\n\t\t\t\t\t\"officename\": \"消化内科\"\n\t\t\t}\n\t\t}\n\t\t}"

			dllInvoke.BusinessHandle_Local(parameter.Conf.ApiParam.Fixmedins_code, parameter.Conf.ApiParam.Infosyscode, parameter.Conf.ApiParam.Infosyssign, param)
		}
	}()
	go func() {
		for {
			<-callExeMenuItem.ClickedCh
			log.Println("callExeMenuItem Clicked")

			// 创建一个命令，替换为你的 exe 文件路径
			cmd := exec.Command(".\\YNReadCard.exe")

			// 设置命令的属性以隐藏控制台窗口s
			//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

			// 创建一个字节缓冲区来捕获输出
			var out bytes.Buffer
			cmd.Stdout = &out

			// 执行命令
			err := cmd.Run()
			if err != nil {
				log.Println("Error:", err)
			}
			// 打印输出
			log.Println("Output:", out.String())
		}
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
	iconData, err := io.ReadAll(file)
	if err != nil {
		log.Println("无法读取图标文件:", err)
		return nil
	}
	defer file.Close()

	return iconData
}
