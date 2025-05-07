package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/marques-mateus/lab-go-cicd-heroku/docs" // importa os docs gerados
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Exemplo Simples API
// @version         1.0
// @description     Esta é uma API simples com Gin e Swagger
// @host            localhost:8080
// @BasePath        /api

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		api.GET("/ping", PingHandler)
	}

	r.Run(":8080")
}

// PingHandler godoc
// @Summary      Ping
// @Description  Retorna pong
// @Tags         ping
// @Success      200  {string}  string  "pong"
// @Router       /ping [get]
func PingHandler(c *gin.Context) {
	c.String(200, "pong")
}
