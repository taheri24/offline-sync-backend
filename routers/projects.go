package routers

import (
	"archive/zip"
	"bufio"
	"bytes"

	"github.com/gofiber/fiber/v2"

	"taheri24.ir/publish-server/database"
	"taheri24.ir/publish-server/internal"
)

func SetupProjectsRouter(apiRouter fiber.Router) {
	router := apiRouter.Group("/projects")
	router.Get("/", func(c *fiber.Ctx) error {
		db := database.Connection

		var items []*database.Project
		if err := db.Find(&items).Error; err != nil {
			panic(err)
		}
		return c.JSON(items)
	})
	router.Get("/count", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	router.Get("/:id", func(c *fiber.Ctx) error {
		db := database.Connection

		targetItem := new(database.Project)
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString("Invalid URLParams(id)," + err.Error())
		}
		db.Find(targetItem, id)
		if targetItem.ID != uint(id) {
			return c.Status(400).SendString("Invalid URLParams(id), not found")

		}
		var b bytes.Buffer
		memWriter := bufio.NewWriter(&b)
		patchZip := zip.NewWriter(memWriter)
		patchZip.Create("")
		return c.JSON(targetItem)
	})
	router.Post("/", func(c *fiber.Ctx) error {
		db := database.Connection

		freshItem := new(database.Project)
		if err := c.BodyParser(freshItem); err != nil {

			return c.Status(400).SendString(err.Error())
		}
		db.Create(freshItem)
		return c.JSON(freshItem)
	})
	router.Put("/:id", func(c *fiber.Ctx) error {
		db := database.Connection

		targetItem := new(database.Project)
		if err := c.BodyParser(targetItem); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString("Invalid URLParams(id)," + err.Error())
		}
		if targetItem.ID != uint(id) {
			return c.Status(400).SendString("Conflict URLParams(id) with Body")

		}
		db.Save(targetItem)
		return c.JSON(targetItem)

	})
	router.Get("/:id/patch", func(c *fiber.Ctx) error {
		db := database.Connection

		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString("Invalid URLParams(id)," + err.Error())
		}
		targetItem := new(database.Project)
		db.Find(targetItem, id)
		if targetItem.ID != uint(id) {
			return c.Status(400).SendString("Conflict URLParams(id) with Body")

		}
		syncProject := internal.NewSyncProject(targetItem)
		fileName, buff := syncProject.GeneratePatch()
		reader := bytes.NewReader(buff.Bytes())
		c.Attachment(fileName)
		return c.SendStream(reader)
	})
	router.Delete("/:id", func(c *fiber.Ctx) error {
		return c.SendString("sdf")
	})
}
