package handlers

import (
	"fmt"
	"template/function"

	"github.com/gofiber/fiber/v2"
)

// Handler to get balance by address
func TestAPIS(c *fiber.Ctx) error {

	address := "0x5009317fd4f6f8feea9dae41e5f0a4737bb7a7d5"
	currency := "bnb"

	message, valid := function.CryptoAddressValidator(address, currency)

	fmt.Println("message =>", message)
	fmt.Println("valid =>", valid)

	return c.SendString(message)

}
