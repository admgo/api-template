package server

import (
	"api-template/internal/api"
	"api-template/internal/handler"

	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	handler.RegisterHandlersWithOptions(engine, api.Api{}, handler.GinServerOptions{
		BaseURL:      "",
		Middlewares:  nil,
		ErrorHandler: nil,
	})
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
