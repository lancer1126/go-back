package router

import "go-back/router/system"

type Group struct {
	System system.RouterGroup
}

var GroupApp = new(Group)
