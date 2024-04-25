package handler

import (
	"github.com/bhill77/music/entity"
	"github.com/bhill77/music/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/thedevsaddam/govalidator"
	"gorm.io/gorm"
)

type SongHandler struct {
	db *gorm.DB
}

func NewSongHandler(db *gorm.DB) *SongHandler {
	return &SongHandler{
		db: db,
	}
}

func (h SongHandler) Add(c *fiber.Ctx) error {
	var song entity.Song
	c.BodyParser(&song)

	rules := govalidator.MapData{
		"title":  []string{"required"},
		"artist": []string{"required"},
	}

	e := helper.Validate(rules, &song)

	if len(e) > 0 {
		return c.Status(400).JSON(e)
	}

	h.db.Create(&song)
	return c.JSON(song)
}

func (h SongHandler) List(c *fiber.Ctx) error {
	var songs []entity.Song
	h.db.Find(&songs)
	return c.JSON(songs)
}

func (h SongHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var song entity.Song
	h.db.First(&song, id)
	if song.ID == 0 {
		return c.Status(404).JSON(map[string]string{"message": "song not found"})
	}

	h.db.Delete(&song)
	return c.JSON(map[string]string{"message": "song deleted"})
}
