package util

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/apex/log/handlers/text"
	"os"
)

func InitLogger() {


	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.InfoLevel)
		logLevel = "info"
	}

	logHandler := os.Getenv("LOG_HANDLER")
	switch logHandler {
	case "text":
		log.SetHandler(text.New(os.Stderr))
	case "logfmt":
		log.SetHandler(logfmt.New(os.Stderr))
	case "json":
		log.SetHandler(json.New(os.Stderr))
	default:
		log.SetHandler(cli.New(os.Stderr))
		logHandler = "cli"
	}

	log.WithFields(log.Fields{
		"Handler": logHandler,
		"Level": logLevel,
	}).Debug("Initializing logger..")
}