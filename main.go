package main

import(
	"github.com/gofiber/fiber/v3"
)

func main(){
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(":3000")
}

func setUpRoutes(app *fiber.App){
	app.Get()
	app.Get()
	app.Post()
	app.Delete()
}

func initDatabase(){
	
}