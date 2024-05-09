package main

import (
	"log"

	"github.com/ghaidafasya24/ws-ats-714220031/config"
	// "github.com/ghaidafasya24/Kuisoner/module"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ghaidafasya24/ws-ats-714220031/url"
	"github.com/gofiber/fiber/v2"
	"github.com/aiteung/musik"
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
