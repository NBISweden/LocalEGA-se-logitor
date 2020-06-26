package main

import (
        "os"
	"github.com/kataras/iris/v12"
        "github.com/kataras/iris/v12/middleware/recover"
        log "github.com/sirupsen/logrus"
)

func main() {

        // Setup logger
        var logger = log.New()
	logger.Formatter = new(log.JSONFormatter)
        logger.Formatter = new(log.TextFormatter)
	logger.Formatter.(*log.TextFormatter).DisableColors = true

        var level log.Level
        if os.Getenv("LOG_LEVEL") == "debug" {
          level = log.DebugLevel
        } else {
          level = log.InfoLevel
        }

	logger.Level = level
	logger.Out = os.Stdout

	app := iris.New()
	app.Use(recover.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
                logger.Debug("welcome page called")
		ctx.JSON(iris.Map{"message": "Noniiin!"})
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
                logger.Debug("hello called ")
		ctx.JSON(iris.Map{"message": "Noniiin!"})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
        var bindAddr string
        if os.Getenv("BIND_ADDR") == "" {
          bindAddr = "0.0.0.0:8080"
        } else {
          bindAddr = os.Getenv("BIND_ADDR")
        }

        logger.Debug("Started web server on port", bindAddr)
	app.Run(iris.Addr(bindAddr), iris.WithoutServerError(iris.ErrServerClosed))
}

