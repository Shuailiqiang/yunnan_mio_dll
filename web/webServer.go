package web

import (
	"YunNanDll/parameter"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Start() {

	// 创建默认的多路复用器
	mux := http.NewServeMux()

	// 设置路由，限制请求方式中间件
	mux.Handle("/Init", onlyGet(http.HandlerFunc(Init)))
	mux.Handle("/BusinessHandle", onlyPost(http.HandlerFunc(BusinessHandle)))

	//启动服务器
	srv := &http.Server{
		Addr:    ":" + parameter.Conf.Port, // 默认端口
		Handler: mux,                       // 使用默认的多路复用器
	}
	// 在一个新的 Goroutine 中启动服务器
	go func() {
		log.Printf("Starting server on port %s...\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
		}
	}()
}

func responseErr(errMesage string) string {
	// 创建一个 map 作为 JSON 数据
	errResponse := map[string]interface{}{
		"code":    "-1",
		"message": errMesage,
	}
	jsonData, err := json.Marshal(errResponse)
	if err != nil {
		log.Println("JSON转换异常:", err)
		return "JSON转换异常" + err.Error()
	}
	// 将 JSON 字符串转换为字符串
	jsonStr := string(jsonData)

	return jsonStr
}

/**
 * 停止web服务器
 */
func ShutdownServer() {
	log.Println("Shutting down server...")

	// 创建一个上下文，设定超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭服务器，等待所有活动连接完成
	if err := parameter.Srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

// 限制仅get方法
func onlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method only allowed GET", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func onlyPost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method only allowed POST", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
