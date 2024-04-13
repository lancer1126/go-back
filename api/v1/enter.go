package v1

import "go-back/api/v1/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
