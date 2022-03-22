package server

import (
	songsroutemap "com.thebeachmaster/entexample/internal/song"
	routedelivery "com.thebeachmaster/entexample/pkg/delivery"
	"github.com/gofiber/fiber/v2"
)

func (srv *Server) MapHTTPHandlers(app *fiber.App) error {

	controller := routedelivery.New(srv.entClient, srv.redis, srv.cfg)

	//TODO: Throw in a bunch of Fiber MW
	apiV1 := app.Group("/api/v1")
	songGroup := apiV1.Group("/song")

	songsroutemap.MapSongRoutes(songGroup, controller.NewDelivery().Songs)

	return nil
}
