package main

import (
	"strconv"

	_ "github.com/SilleCao/golang/go-micro-demo/docs"
	"github.com/SilleCao/golang/go-micro-demo/internal/config"
	"github.com/SilleCao/golang/go-micro-demo/internal/server"
	"github.com/gin-gonic/gin"
)

// @title           Go Micro Demo
// @version         1.0
// @description     A demo service in Go using Gin framework
// @termsOfService  https://sille.cn

// @contact.name   Sille Cao
// @contact.email  caoliang1025@163.com

// @host      localhost:7081
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	router := gin.New()

	// Register common middleware.
	// router.Use(Recovery(), Security(conf), Logger())

	// Find and load templates.
	// router.LoadHTMLFiles(conf.TemplateFiles()...)

	// Register HTTP route handlers.
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	conf.InitDB()
	server.RegisterRoutes(router, conf)
	router.Run(":" + strconv.Itoa(conf.Server.Port))
}
