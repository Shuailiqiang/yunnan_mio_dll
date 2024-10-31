package parameter

import (
	"YunNanDll/entity"
	"net/http"
	"syscall"
)

var (
	SiInterface_hsf *syscall.LazyDLL // 动态库驱动
	Srv             *http.Server     // 声明HTTP服务器
	Conf            *entity.Config   // 配置文件
)

const (
	POINTER_SIZE = 10 * 1024 * 1024 // 指针分配内存大小 单位byte
)
