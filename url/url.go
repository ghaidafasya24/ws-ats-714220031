package url

import (
	"github.com/ghaidafasya24/ws-ats-714220031/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage)
	page.Get("/responden", controller.GetResponden)
	page.Get("/pertanyaan", controller.GetPertanyaan)
	// page.Get("/responden/:id", controller.GetRespondenByID)
	// page.Get("/", controller.GetResponden)
}