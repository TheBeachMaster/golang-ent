package song

import "github.com/gofiber/fiber/v2"

type SongHTTPHandler interface {
	AddNewSong() fiber.Handler
	UpdateSongLyrics() fiber.Handler
}
