package handler

import (
	"fmt"
	"os"

	"github.com/bhill77/music/config"
	"github.com/bhill77/music/entity"
	"github.com/bhill77/music/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SongHandler struct {
	db   *gorm.DB
	conf config.Config
}

func NewSongHandler(db *gorm.DB, conf config.Config) *SongHandler {
	return &SongHandler{
		db:   db,
		conf: conf,
	}
}

func (h SongHandler) Add(c *fiber.Ctx) error {
	var song entity.Song
	c.BodyParser(&song)

	e := helper.Validate(&song)

	if len(e) > 0 {
		return c.Status(400).JSON(e)
	}

	if _, err := os.Stat(h.conf.StoragePath + "/" + song.File); err != nil {
		return c.Status(400).JSON(map[string]string{"message": "file not found"})
	}

	h.db.Create(&song)
	return c.JSON(song)
}

func (h SongHandler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(err)
	}

	fmt.Println(h.conf.StoragePath)
	err = c.SaveFile(file, fmt.Sprintf("%s/%s", h.conf.StoragePath, file.Filename))
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON(file.Filename)
}

func (h SongHandler) List(c *fiber.Ctx) error {
	var songs []entity.Song
	h.db.Find(&songs)
	return c.JSON(songs)
}

func (h SongHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var song entity.Song
	h.db.First(&song, id)
	if song.ID == 0 {
		return c.Status(404).JSON(map[string]string{"message": "song not found"})
	}

	var payload entity.Song
	c.BodyParser(&payload)
	e := helper.Validate(&payload)
	if len(e) > 0 {
		return c.Status(400).JSON(e)
	}

	if _, err := os.Stat(h.conf.StoragePath + "/" + song.File); err != nil {
		return c.Status(400).JSON(map[string]string{"message": "file not found"})
	}

	song.Artist = payload.Artist
	song.Title = payload.Title
	song.Cover = payload.Cover
	song.File = payload.File

	h.db.Updates(&song)
	return c.JSON(song)
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
