package web

import (
	"YunNanDll/dllInvoke"
	"YunNanDll/entity"
	"YunNanDll/util"
	"log"
	"net/http"
)

func Init(w http.ResponseWriter, r *http.Request) {

	//获取请求参数
	queryParamter := r.URL.Query()

	fixmedins_code := queryParamter.Get("fixmedins_code")
	infosyscode := queryParamter.Get("infosyscode")
	infosyssign := queryParamter.Get("infosyssign")
	url := queryParamter.Get("url")
	log.Println("fixmedins_code ===> ", fixmedins_code)
	log.Println("infosyscode ===> ", infosyscode)
	log.Println("infosyssign ===> ", infosyssign)
	log.Println("url ===> ", url)

	config := entity.Config{}
	config.ApiParam.Fixmedins_code = fixmedins_code
	config.ApiParam.Infosyscode = infosyscode
	config.ApiParam.Infosyssign = infosyssign
	config.ApiParam.Url = url

	response, dllErr := dllInvoke.Init_Local(&config)

	//校验dll
	if dllErr != nil {
		log.Println("Init error ===>", dllErr)
		http.Error(w, responseErr(dllErr.Error()+",response:"+response), http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "text/plain;charset=GBK")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func BusinessHandle(w http.ResponseWriter, r *http.Request) {

	//获取请求参数
	queryParamter := r.URL.Query()

	fixmedins_code := queryParamter.Get("fixmedins_code")
	infosyscode := queryParamter.Get("infosyscode")
	infosyssign := queryParamter.Get("infosyssign")

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	inputData := string(util.Utf8BytesToGB2312(body))

	log.Println("fixmedins_code ===> ", fixmedins_code)
	log.Println("infosyscode ===> ", infosyscode)
	log.Println("infosyssign ===> ", infosyssign)
	log.Println("inputData ===> ", inputData)

	response, dllErr := dllInvoke.BusinessHandle_Local(fixmedins_code, infosyscode, infosyssign, inputData)

	//校验dll
	if dllErr != nil {
		log.Println("ReadydSSECard_Social error ===>", dllErr)
		http.Error(w, responseErr(response), http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "text/plain;charset=GBK")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
