package parameter

import (
	"YunNanDll/entity"
	"net/http"
	"syscall"
)

var (
	SiInterface_DLL *syscall.LazyDLL // 动态库驱动
	User32_DLL      *syscall.LazyDLL // 动态库驱动
	Srv             *http.Server     // 声明HTTP服务器
	Conf            *entity.Config   // 配置文件
	InitStatus      bool             = false
)

const (
	POINTER_SIZE       = 10 * 1024 * 1024 // 指针分配内存大小 单位byte
	HWND_TOPMOST       = 0xFFFFFFFF
	SWP_NOSIZE         = 0x0001
	SWP_NOMOVE         = 0x0002
	SWP_SHOWWINDOW     = 0x0040
	MB_OK              = 0x00000000
	MB_ICONINFORMATION = 0x00000040
)
