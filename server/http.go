package server

import (
	songRepo "com.thebeachmaster/entexample/internal/song/repository"
	songHTTPHandler "com.thebeachmaster/entexample/internal/song/routes"
	songUseCase "com.thebeachmaster/entexample/internal/song/usecase"
	"github.com/gofiber/fiber/v2"
)

func (srv *Server) MapHTTPHandlers(app *fiber.App) error {
	songRepo := songRepo.NewSongRespository(srv.entClient)
	songUsecase := songUseCase.NewSongUseCase(srv.cfg, songRepo)
	songHandler := songHTTPHandler.NewSongHTTPHandler(srv.cfg, songUsecase)

	//TODO: Throw in a bunch of Fiber MW
	apiV1 := app.Group("/api/v1")     // https://domain.tld/api/v1
	songGroup := apiV1.Group("/song") // https://domain.tld/api/v1/song

	songHTTPHandler.MapSongRoutes(songGroup, songHandler)

	return nil
}
