package core

import (
	"go-back/global"
	"go.uber.org/zap"
)

func CloseEnv() {
	if global.GB_DB != nil {
		db, _ := global.GB_DB.DB()
		err := db.Close()
		if err != nil {
			global.GB_LOG.Error("close db fail", zap.Error(err))
			return
		}
		global.GB_LOG.Info("close db success")
	}
}
