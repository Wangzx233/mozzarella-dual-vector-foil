package router

import (
	"github.com/gin-gonic/gin"
	"mozzarella-dual-vector-foil/api"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/picture/upload-token", api.GetUpToken)

	r.Run(":8081")
}
