package main

import (
	"strconv"

	"github.com/SilleCao/golang/go-micro-demo/internal/config"
	"github.com/SilleCao/golang/go-micro-demo/internal/server"
	"github.com/gin-gonic/gin"
)

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
