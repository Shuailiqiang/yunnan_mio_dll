package dllInvoke

import (
	"YunNanDll/entity"
	"YunNanDll/parameter"
	"YunNanDll/util"
	"errors"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func LoadDirDll() error {
	dllPath := "CHSInterfaceYn.dll"
	if _, err := os.Stat(dllPath); os.IsNotExist(err) {
		log.Println("CHSInterfaceYn.dll is not exist:", dllPath)
		return err
	} else {
		log.Println("CHSInterfaceYn.dll is exist")
	}

	parameter.SiInterface_DLL = syscall.NewLazyDLL("CHSInterfaceYn.dll")
	if parameter.SiInterface_DLL == nil {
		log.Println("load CHSInterfaceYn.dll faild")
		return errors.New("load CHSInterfaceYn.dll faild")
	}
	return nil
}

func BusinessHandle_Local(fixmedins_code, infosyscode, infosyssign, inputData string) (string, error) {

	//调用动态库的入参
	fixmedins_code_uintptr := uintptr(unsafe.Pointer(syscall.StringBytePtr(fixmedins_code)))
	infosyscode_uintptr := uintptr(unsafe.Pointer(syscall.StringBytePtr(infosyscode)))
	infosyssign_uintptr := uintptr(unsafe.Pointer(syscall.StringBytePtr(infosyssign)))
	inputData_uintptr := uintptr(unsafe.Pointer(syscall.StringBytePtr(inputData)))

	// 获取出参的指针
	outputInfo := make([]byte, parameter.POINTER_SIZE)
	pOutputInfo := uintptr(unsafe.Pointer(&outputInfo[0]))

	// 获取异常的指针
	errMsg := make([]byte, parameter.POINTER_SIZE)
	pErrMsg := uintptr(unsafe.Pointer(&errMsg[0]))

	//调用动态库
	BusinessHandle := parameter.SiInterface_DLL.NewProc("BusinessHandle")
	ret, _, _ := BusinessHandle.Call(fixmedins_code_uintptr, infosyscode_uintptr, infosyssign_uintptr, inputData_uintptr, pOutputInfo, pErrMsg)

	log.Println("BusinessHandle resultCode ===>", ret)

	// 获取输出数据
	log.Println("BusinessHandle pOutputInfo ===>" + string(outputInfo))
	response := strings.TrimRight(util.Utf8BytesToGB2312(outputInfo), "\x00")
	log.Println("BusinessHandle pOutputInfo ===>" + response)

	// 获取异常数据
	log.Println("BusinessHandle errMsg ===>" + string(errMsg))
	errResponse := strings.TrimRight(util.Utf8BytesToGB2312(errMsg), "\x00")
	log.Println("BusinessHandle pErrMsg ===>" + errResponse)

	// 根据 ret 值判断调用是否成功
	if ret != 0 || errResponse != "" {
		return errResponse, errors.New(errResponse)
	}

	return response, nil
}

func Init_Local(config *entity.Config) (string, error) {

	//调用动态库的入参
	fixmedins_code_uintptr := uintptr(unsafe.Pointer(syscall.StringBytePtr(config.ApiParam.Fixmedins_code)))
	infosyscode_uintptr := uintptr(unsafe.Pointer(syscall.StringBytePtr(config.ApiParam.Infosyscode)))
	infosyssign_uintptr := uintptr(unsafe.Pointer(syscall.StringBytePtr(config.ApiParam.Infosyssign)))
	url_uintptr := uintptr(unsafe.Pointer(syscall.StringBytePtr(config.ApiParam.Url)))

	// 获取出参的指针
	errMsg := make([]byte, parameter.POINTER_SIZE)
	pErrMsg := uintptr(unsafe.Pointer(&errMsg[0]))

	//调用动态库
	BusinessHandle := parameter.SiInterface_DLL.NewProc("Init")
	ret, _, _ := BusinessHandle.Call(fixmedins_code_uintptr, infosyscode_uintptr, infosyssign_uintptr, url_uintptr, pErrMsg)

	log.Println("Init resultCode ===>", ret)

	// 获取异常数据
	errResponse := strings.TrimRight(util.Utf8BytesToGB2312(errMsg), "\x00")
	log.Println("Init pErrMsg ===>" + errResponse)

	// 根据 ret 值判断调用是否成功
	if ret != 0 || errResponse != "" {
		return errResponse, errors.New(errResponse)
	}

	parameter.InitStatus = true
	return "成功", nil
}
