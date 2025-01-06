package util

import (
	"YunNanDll/entity"
	"YunNanDll/parameter"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func InitConfig(file string) {
	log.Println("load config file " + file + " begin")
	var config entity.Config
	configFile, err := os.Open(file)
	if err != nil {
		log.Printf("load config file error ===> "+err.Error(), err)
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	if err := yaml.Unmarshal(byteValue, &config); err != nil {
		log.Printf("load config file error ===> "+err.Error(), err)
	}

	parameter.Conf = &config

	log.Println("fixmedins_code ===> ", parameter.Conf.ApiParam.Fixmedins_code)
	log.Println("infosyscode ===> ", parameter.Conf.ApiParam.Infosyscode)
	log.Println("infosyssign ===> ", parameter.Conf.ApiParam.Infosyssign)
	log.Println("url ===> ", parameter.Conf.ApiParam.Url)
	log.Println("port ===> ", parameter.Conf.Port)

	log.Println("load config file " + file + " success")
}
