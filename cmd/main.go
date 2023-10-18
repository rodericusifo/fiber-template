package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/rodericusifo/fiber-template/internal/app/core"
	"github.com/rodericusifo/fiber-template/internal/pkg/config"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/handler"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/runner"

	log "github.com/sirupsen/logrus"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
)

func init() {
	config.ConfigureLog()
	config.ConfigureEnv()
	config.ConfigureDatabaseCache(pkg_constant.REDIS)
	config.ConfigureDatabaseSQL(pkg_constant.MYSQL)
	config.ConfigureAuth()

	runner.RunDatabaseSeederSQL(pkg_constant.MYSQL)
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.HandleHTTPError,
	})

	app.Use(
		requestid.New(),
		logger.New(logger.Config{
			Format: "[${time}] ${pid} | ${locals:requestid} | ${status} | ${latency} | ${method} | ${path}\n",
		}),
		recover.New(),
		cors.New(cors.Config{
			AllowMethods: "GET,POST,DELETE,PATCH",
		}),
	)

	core.InitRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", getter.GetEnvServerPort()))
	if err != nil {
		log.WithFields(log.Fields{
			"message": "application failed to run",
			"detail":  err,
		}).Fatal("[MAIN]")
	}
}
