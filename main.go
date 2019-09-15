package main

import (
	"fmt"
	"github.com/apex/log"
	"github.com/gabeduke/level-api/pkg/handlers"
	"github.com/gabeduke/level-api/pkg/util"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	util.InitLogger()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := gin.Default()
	r.GET("/healthz", handlers.Healthz)
	r.GET("/level", handlers.Level)


	log.WithField("Port", port).Info("Starting service..")
	r.Run(fmt.Sprintf(":%s", port))
}
