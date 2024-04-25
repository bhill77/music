package main

import (
	"github.com/bhill77/music/entity"
	"github.com/bhill77/music/handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:123456@tcp(127.0.0.1:3306)/music_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Song{})

	app := fiber.New()

	songHandler := handler.NewSongHandler(db)

	app.Get("/", handler.HelloHandler)

	// song
	app.Post("/song", songHandler.Add)
	app.Get("/song", songHandler.List)
	app.Delete("/song/:id", songHandler.Delete)

	app.Listen(":4000")
}
