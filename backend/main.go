package main

import (
	"github.com/MuShaf-NMS/zesthub_test/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(cors.Default())
	router.InitializeRouter(server)
	server.Run(":8000")
}
