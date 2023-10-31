package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var (
	port  int
	local bool
)

func init() {
	flag.IntVar(&port, "port", 9005, "web service port")
	flag.BoolVar(&local, "local", true, "run web service locally")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Create a data structure to store project-related information
	projectInfo := struct {
		ProjectName            string
		ProjectNameExplanation string
		Description            string
		GithubLink             string
		ImageURL               string
		// You can add other project-related information here
	}{
		ProjectName:            os.Getenv("PROJECT_NAME"),
		ProjectNameExplanation: os.Getenv("PROJECT_NAME_EXPLANATION"),
		Description:            os.Getenv("DESCRIPTION"),
		GithubLink:             os.Getenv("GITHUB_LINK"),
		ImageURL:               os.Getenv("IMAGE_URL"),
	}
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", projectInfo)
	})

	// Serve Swagger documentation at /swagger/index.html
	app.Static("/swagger", "./swagger")

	log.Printf("Web service running on [::]:%d\n", port)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
