package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	CHSInterfaceYn = syscall.NewLazyDLL("CHSInterfaceYn.dll")

	fixmedins_code = uintptr(unsafe.Pointer(syscall.StringBytePtr("H53010200351")))
	infosyscode    = uintptr(unsafe.Pointer(syscall.StringBytePtr("915101006818136552")))               //社会统一信用服务编码
	infosyssign    = uintptr(unsafe.Pointer(syscall.StringBytePtr("neu1f61023a346fcafb0b3d1dbd57e64"))) //医院私钥

	gbkDecoder = simplifiedchinese.GBK.NewDecoder()
)

// 声明dll

func main() {
	file, logErr := os.OpenFile("dll.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if logErr != nil {
		log.Fatal("Failed to open log file:", logErr)
	}
	// 设置日志输出目标为文件
	log.SetOutput(file)

	var url uintptr = uintptr(unsafe.Pointer(syscall.StringBytePtr("http://10.114.177.55:8080/eapdomain/callService"))) //URL

	const errMsgSize = 1024 // 根据需要调整大小
	errMsg := make([]byte, errMsgSize)
	errMsgPtr := uintptr(unsafe.Pointer(&errMsg[0])) // 出参

	//初始化
	Init := CHSInterfaceYn.NewProc("Init")
	log.Println("init begin ====>")
	ret, _, _ := Init.Call(fixmedins_code, infosyscode, infosyssign, url, errMsgPtr)

	if ret != 0 {
		log.Println("Error:", string(errMsg))
		return
	}

	http.HandleFunc("/businessHandle", aboutHandler)
	//启动服务器
	port := ":8765"
	log.Printf("Starting server on port %s...\n", port)
	webErr := http.ListenAndServe(port, nil)

	if webErr != nil {
		log.Println("start web server err:", webErr)
	}
	// go build -ldflags "-s -w -H=windowsgui" 执行时不显示控制台

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// 获取请求报文的内容长度
	len := r.ContentLength
	// 新建一个字节切片，长度与请求报文的内容长度相同
	body := make([]byte, len)
	// 读取 r 的请求主体，并将具体内容读入 body 中
	r.Body.Read(body)
	// 将字节切片内容写入相应报文

	inputDataGBK, _ := gbkDecoder.Bytes(body)

	inputString := string(inputDataGBK)
	log.Println("requestBody ===> ", inputString)

	inputData := uintptr(unsafe.Pointer(syscall.StringBytePtr(inputString)))

	// 准备输出数据
	const outputDataSize = 1024 // 根据需要调整大小
	outputData := make([]byte, outputDataSize)
	pOutputData := uintptr(unsafe.Pointer(&outputData[0]))

	const errMsgSize = 1024 // 根据需要调整大小
	errMsg := make([]byte, errMsgSize)
	pErrMsg := uintptr(unsafe.Pointer(&errMsg[0]))

	BusinessHandleW := CHSInterfaceYn.NewProc("BusinessHandle")
	ret, _, _ := BusinessHandleW.Call(fixmedins_code, infosyscode, infosyssign, inputData, pOutputData, pErrMsg)

	if ret != 0 {
		log.Println("Error ===>", string(errMsg))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	outputDataGBK, _ := gbkDecoder.Bytes(outputData)

	// 获取输出数据 utf-8
	response := string(outputDataGBK)
	response = strings.TrimRight(response, "\x00")
	log.Println("Output  ===>" + response)

	// 设置响应头
	w.Header().Set("Content-Type", "text/plain")

	// 发送字符串响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
