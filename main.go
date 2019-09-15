package main

import (
	"fmt"
	"os"

	log "github.com/apex/log"
	"github.com/gabeduke/level-api/pkg/router"
	"github.com/gabeduke/level-api/pkg/util"
)

func main() {

	util.InitLogger()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := router.GetRouter()

	log.WithField("Port", port).Info("Starting service..")
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err.Error())
	}
}
