package main

import (
	"fmt"
	log "github.com/apex/log"
	"github.com/gabeduke/level/docs"
	"github.com/gabeduke/level/pkg/router"
	"github.com/gabeduke/level/pkg/util"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Level API
// @version 1.0
// @description API to get the water level from NWS
// @termsOfService http://swagger.io/terms/

// @contact.name Dukemon
// @contact.url leetserve.com
// @contact.email gabeduke@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

var version = "dev"

func main() {

	port := util.GetPort()

	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// get router
	r := router.GetRouter()

	// swagger
	u := fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)
	url := ginSwagger.URL(u) // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// run
	log.WithField("Port", port).Info("Starting service..")
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	util.InitLogger()
}
