package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mauljurassicia/go-crm/database"
	"github.com/mauljurassicia/go-crm/lead"
	"github.com/glebarez/sqlite" 
	"gorm.io/gorm"
)

func main(){
	initDatabase()
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(":3000")
}

func setUpRoutes(app *fiber.App){
	app.Get("api/v1/leads",lead.GetLeads)
	app.Get("api/v1/lead/:id",lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DbCoon, err = gorm.Open(sqlite.Open("crm.db"), &gorm.Config{})
	if err != nil{
		panic("fail to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DbCoon.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}
