package lead

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mauljurassicia/go-crm/database"
	"gorm.io/gorm"
)

type Lead struct{
	gorm.Model
	Name string
	Company string
	Email string
	Phone int
}

func GetLeads(c fiber.Ctx) error{
	db := database.DbCoon
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c fiber.Ctx) error{
	id := c.Params("id")
	db := database.DbCoon
	var lead Lead
	if err := db.First(&lead, id).Error; err != nil{
		return c.Status(404).Send([]byte(err.Error()))
	}
	return c.JSON(lead)
}

func NewLead(c fiber.Ctx) error{
	db := database.DbCoon
	lead := new(Lead)
	if err := c.Bind().Body(lead); err != nil{
		c.Status(503).Send([]byte(err.Error()))
		return err
	}
	db.Create(lead)
	return c.JSON(lead)
}

func DeleteLead(c fiber.Ctx) error{
	id := c.Params("id")
	db := database.DbCoon
	var lead Lead
	if err := db.First(&lead, id).Error; err != nil {
		return c.Status(404).Send([]byte(err.Error()))
	}
	db.Delete(&lead)
	return c.Send([]byte("Lead succesfully deleted"))
}