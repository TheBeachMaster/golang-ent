package song

import (
	"github.com/gofiber/fiber/v2"
)

func MapSongRoutes(f fiber.Router, h SongHTTPRoutes /*Middlewares*/) {
	f.Post("/add", h.AddNewSong())
	f.Put("/lyrics", h.UpdateSongLyrics())
	f.Get("/list", h.GetAllSongs())
	f.Get(":hash", h.GetSong())
}
