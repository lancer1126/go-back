package internal

import (
	"go-back/global"
	"go.uber.org/zap/zapcore"
	"os"
)

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotateLogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := NewCutter(global.GB_CONFIG.Zap.Director, level, WithCutterFormat("2006-01-02"))
	if global.GB_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
