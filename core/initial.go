package core

import (
	"fmt"
	"go-back/global"
)

func Init() {
	global.GB_VP = Viper()
	global.GB_DB = Gorm()
}

func RunServe() {
	Router := Routers()
	Router.Static("/form-generator", "./resource/page")
	addr := fmt.Sprintf(":%d", global.GB_CONFIG.System.Addr)
	initServer(addr, Router)
}
