package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-back/core/internal"
	"go-back/global"
	"os"
)

func Viper(path ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(parseConfig(path...))
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GB_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GB_CONFIG); err != nil {
		panic(err)
	}
	return v
}

func parseConfig(path ...string) string {
	var config string
	var configType string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
				case gin.TestMode:
					config = internal.ConfigTestFile
					configType = ""
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
				}
				configType = "gin模式-" + gin.EnvGinMode
			}
		} else {
			configType = "命令行的-c参数传递的值"
		}
	} else {
		config = path[0]
		configType = "func Viper()传递的"
	}
	fmt.Printf("使用%s，config的路径为%s\n", configType, config)
	return config
}
