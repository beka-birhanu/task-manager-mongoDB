package common

import "github.com/gin-gonic/gin"

type IContrller interface {
	Register(route gin.RouterGroup)
}
