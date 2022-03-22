package song

import "github.com/gofiber/fiber/v2"

type SongHTTPRoutes interface {
	AddNewSong() fiber.Handler
	UpdateSongLyrics() fiber.Handler
	GetAllSongs() fiber.Handler
	GetSong() fiber.Handler
}
