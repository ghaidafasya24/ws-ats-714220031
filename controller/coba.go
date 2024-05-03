package controller

import (
	"github.com/ghaidafasya24/Kuisoner"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := _714220031.GetAllKuisoner()
	return c.JSON(ipaddr)
}

