package main

import (
	"go-back/core"
)

func main() {
	core.Init()
	defer core.CloseEnv()
	core.RunServe()
}
