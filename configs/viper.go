package configs

import (
	"os"

	"github.com/spf13/viper"
)

func NewViperConfig() *viper.Viper {
	os.Setenv("ENVIRONMENT", "local")

	var env = os.Getenv("ENVIRONMENT")
	viper.SetConfigName("config." + env)
	viper.AddConfigPath("./configs/env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	return viper.GetViper()
}
