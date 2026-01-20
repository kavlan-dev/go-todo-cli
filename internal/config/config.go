package config

import "github.com/spf13/viper"

var MaxTaskLength uint
var TasksPath string

func SetupConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.SetDefault("TasksPath", "tasks.json")
	v.SetDefault("MaxTaskLength", 200)

	v.ReadInConfig()

	MaxTaskLength = v.GetUint("MaxTaskLength")
	TasksPath = v.GetString("TasksPath")
}
