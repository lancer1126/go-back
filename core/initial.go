package core

import (
	"go-back/global"
	"go.uber.org/zap"
)

func Init() {
	global.GB_VP = Viper()
	global.GB_LOG = Zap()
	zap.ReplaceGlobals(global.GB_LOG)
	global.GB_DB = Gorm()
	if global.GB_DB != nil {
		RegisterTables()
	}
}

func RunServe() {
	// todo RunServe
}
