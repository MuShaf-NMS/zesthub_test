package router

import (
	"github.com/MuShaf-NMS/zesthub_test/handler"
	"github.com/MuShaf-NMS/zesthub_test/service"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(server *gin.Engine) {
	service := service.NewService()
	handler := handler.NewHandler(service)
	server.GET("/", handler.Index)
	server.POST("/snaki", handler.Snaki)
}
