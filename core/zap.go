package core

import (
	"fmt"
	"go-back/core/internal"
	"go-back/global"
	"go-back/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 启动日志框架
func Zap() (logger *zap.Logger) {
	// 判断是否已存在日志文件夹
	if exists, _ := utils.PathExists(global.GB_CONFIG.Zap.Director); !exists {
		fmt.Printf("create %v directory\n", global.GB_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GB_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GB_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
