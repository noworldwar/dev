package starter

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig(name string, path ...string) {

	viper.SetConfigName(name)

	for _, v := range path {
		viper.AddConfigPath(v)
	}

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Fatal error when reader %s config file: %s", name, err.Error())
	}
}
