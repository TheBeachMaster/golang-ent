package routes

import (
	"com.thebeachmaster/entexample/config"
	"com.thebeachmaster/entexample/internal/song"
	"com.thebeachmaster/entexample/internal/song/models"
	"github.com/gofiber/fiber/v2"
)

type songHTTPHandler struct {
	cfg         *config.Config
	songUsecase song.SongUseCase
}

func NewSongHTTPHandler(c *config.Config, u song.SongUseCase) song.SongHTTPRoutes {
	return &songHTTPHandler{cfg: c, songUsecase: u}
}

// TODO: Document
func (s *songHTTPHandler) AddNewSong() fiber.Handler {
	// TODO: Securely Parse body
	return func(c *fiber.Ctx) error {
		song := &models.CreateSong{}
		if err := c.BodyParser(song); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		addSong, err := s.songUsecase.AddNewSong(c.Context(), song)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"error":   false,
			"message": nil,
			"song":    addSong,
		})
	}
}

func (s *songHTTPHandler) UpdateSongLyrics() fiber.Handler {
	return func(c *fiber.Ctx) error {
		songLyrics := &models.UpdateSongLyricsFileURL{}
		if err := c.BodyParser(songLyrics); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		updateLyrics, err := s.songUsecase.AddSongLyrics(c.Context(), songLyrics)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"error":      false,
			"message":    nil,
			"songLyrics": updateLyrics,
		})
	}
}

func (s *songHTTPHandler) GetAllSongs() fiber.Handler {
	return func(c *fiber.Ctx) error {
		songList, err := s.songUsecase.ListAllSongs(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"error":   false,
			"message": nil,
			"songs":   songList,
		})
	}
}

func (s *songHTTPHandler) GetSong() fiber.Handler {
	return func(c *fiber.Ctx) error {
		songHash := c.Params("hash")

		isEmpty := len(songHash) == 0

		if isEmpty {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "The param 'hash' can not be empty",
			})
		}

		songInfo, err := s.songUsecase.FetchSong(c.Context(), songHash)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"error":   false,
			"message": nil,
			"song":    songInfo,
		})
	}
}
