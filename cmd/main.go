package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/rodericusifo/fiber-template/internal/pkg/config"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/getter"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/handler"
	"github.com/rodericusifo/fiber-template/internal/pkg/util/runner"

	log "github.com/sirupsen/logrus"

	internal_app_core_auth_controller_api "github.com/rodericusifo/fiber-template/internal/app/core/auth/controller/api"
	internal_app_core_employee_controller_api "github.com/rodericusifo/fiber-template/internal/app/core/employee/controller/api"
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
		ServerHeader: "Fiber",
		AppName:      fmt.Sprintf("%s v%s", getter.GetEnvConfig().AppsName, getter.GetEnvConfig().AppsVersion),
	})

	app.Use(
		requestid.New(),
		logger.New(logger.Config{
			Format: "[${time}] ${pid} | ${locals:requestid} | ${status} | ${latency} | ${method} | ${path}\n",
		}),
		recover.New(),
		cors.New(cors.Config{
			AllowMethods: "GET,POST,DELETE,PUT",
		}),
	)

	apiVersion := "/v" + strings.Split(getter.GetEnvConfig().AppsVersion, ".")[0]
	internal_app_core_auth_controller_api.InitAPI(app.Group(apiVersion))
	internal_app_core_employee_controller_api.InitAPI(app.Group(apiVersion))

	go func() {
		err := app.Listen(fmt.Sprintf(":%d", func() int {
			serverPort := getter.GetEnvConfig().ServerPort
			if serverPort != 0 {
				return serverPort
			} else {
				return constant.DEFAULT_ENV_SERVER_PORT.(int)
			}
		}()))
		if err != nil {
			log.WithFields(log.Fields{
				"message": "error starting server",
				"detail":  err,
			}).Fatal("[MAIN]")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.WithFields(log.Fields{
		"message": "shutting down server...",
	}).Infoln("[MAIN]")

	if err := app.Shutdown(); err != nil {
		log.WithFields(log.Fields{
			"message": "server forced to shutdown",
			"detail":  err,
		}).Fatal("[MAIN]")
	}

	log.WithFields(log.Fields{
		"message": "server shutdown gracefully",
	}).Infoln("[MAIN]")
}
