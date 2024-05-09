package controller

import (
	"github.com/gofiber/fiber/v2"
	cek "github.com/ghaidafasya24/Kuisoner/module"
	"github.com/aiteung/musik"
	// "github.com/ghaidafasya24/Kuisoner"
	// "github.com/ghaidafasya24/ws-ats-714220031/config"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetResponden(c *fiber.Ctx) error {
	ps := cek.GetAllResponden()
	return c.JSON(ps)
}
