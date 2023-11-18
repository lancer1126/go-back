package global

import (
	"github.com/spf13/viper"
	"go-back/config"
	"gorm.io/gorm"
)

// GB -> go-back的缩写，表全局配置
var (
	GB_CONFIG config.Server
	GB_DB     *gorm.DB
	GB_VP     *viper.Viper
)
