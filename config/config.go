package config

import "github.com/spf13/viper"

var cfg *viper.Viper

func init() {
	cfg = viper.New()
	cfg.SetConfigName("config")
	cfg.AddConfigPath("./")
	cfg.SetConfigType("yaml")
	cfg.BindEnv("token", "TOKEN")
	cfg.AutomaticEnv()
	err := cfg.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetCfonfig() *viper.Viper {
	return cfg
}
