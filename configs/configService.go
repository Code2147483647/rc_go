package configs

import (
	"bytes"
	"io/ioutil"
	"log"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once
var instance *viper.Viper

//GetInstance for configService
func GetInstance() *viper.Viper {
	once.Do(func() {
		content, err := ioutil.ReadFile("./configs/local.json")
		if err != nil {
			log.Fatal("failed to read config file", err)
		}
		viper.SetConfigType("json")
		viper.ReadConfig(bytes.NewBuffer(content))
		if err != nil {
			log.Fatal("failed to read config By viper", err)
		}
		instance = viper.GetViper()
	})
	return instance
}
