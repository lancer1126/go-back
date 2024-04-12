package core

import (
	"github.com/gin-gonic/gin"
	"go-back/global"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func Routers() *gin.Engine {
	if global.GB_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router := gin.New()
	if global.GB_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger(), gin.Recovery())
	}
	return Router
}

func initServer(addr string, router *gin.Engine) server {
	return &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
