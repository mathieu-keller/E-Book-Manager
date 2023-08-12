package resources

import (
	"e-book-manager/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func SetupRoutes() {
	bodyLimit, err := strconv.Atoi(util.GetEnvOrDefault("BODY_LIMIT", "100"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	app := fiber.New(fiber.Config{
		BodyLimit:   bodyLimit * 1024 * 1024,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	api := app.Group("/api")
	InitUploadApi(api)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	log.Fatalln(app.Listen(":8080"))
}
