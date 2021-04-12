package config

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func Init() {
	Config = viper.New()
	Config.SetEnvPrefix("run")
	Config.AutomaticEnv()
	mode := Config.GetString("mode")
	if len(mode) == 0 {
		mode = "alpha"
		Config.SetDefault("mode", mode)
	}

	Config.SetConfigType("yaml")
	Config.AddConfigPath("config/" + mode)
	Config.SetConfigName("app")
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
