package config

import (
	"github.com/spf13/viper"

	"GO_PROJECT/logger"
	"GO_PROJECT/model"
)

// initialize nil pointer with data type model.Configurations
var Config *model.Configurations

func GetConfigurations() *model.Configurations {

	// if Config has a value i.e not nil\ pointer
	if Config != nil {
		return Config
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("./config/data/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Fatal("Error reading config file ", err)
	}

	// initialize variable configuration with data type model.Configurations
	var configuration model.Configurations

	// viper has config data so it unmarshal it in location of configuration i.e(&configuration)
	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Log.Fatal("Unable to decode config info using viper ", err)
	}

	// now config is pointing towards location of configuration i.e(&configuration)
	Config = &configuration

	return Config
}
