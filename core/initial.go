package core

import (
	"fmt"
	"go-back/global"
	"go-back/initialize"
	"go.uber.org/zap"
)

func Init() {
	// 初始化viper，读取配置
	global.GB_VP = Viper()
	// 初始化zap日志框架
	global.GB_LOG = Zap()
	zap.ReplaceGlobals(global.GB_LOG)
	// 初始化一些其他配置
	initialize.OtherInit()

	// 初始化gorm，链接数据库
	global.GB_DB = Gorm()
	if global.GB_DB != nil {
		RegisterTables()
	}
}

func RunServe() {
	// 初始化路由
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GB_CONFIG.System.Addr)
	// 初始化服务
	s := initServer(address, Router)
	global.GB_LOG.Info("server run success on ", zap.String("address", address))
	// 启动服务
	err := s.ListenAndServe()
	if err != nil {
		global.GB_LOG.Error(err.Error())
	}
}
