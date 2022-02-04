package routers
import (
	"github.com/gofiber/fiber/v2"
)
func SetupApiRouters(app *fiber.App ){
	apiGroup := app.Group("/api")
	SetupProjectsRouter(apiGroup)
}


