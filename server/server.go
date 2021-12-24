package server

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"com.thebeachmaster/entexample/config"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{app: fiber.New(fiber.Config{
		Prefork:      cfg.Server.Prefork,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout),
		AppName:      cfg.Server.AppName,
		ServerHeader: cfg.Server.ServerHeader,
	})}
}

func (srv *Server) Run() error {
	go func() {
		log.Printf("Server is listening on PORT: %s", strconv.Itoa(srv.cfg.Server.Port))
		addr := ":" + strconv.Itoa(srv.cfg.Server.Port)
		if err := srv.app.Listen(addr); err != nil {
			log.Panicf("[CRIT] Unable to start server. Reason: %v", err)
		}
	}()

	quitServer := make(chan struct{})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	close(quitServer)

	<-quitServer

	log.Printf("Server shutdown")
	return srv.app.Shutdown()

}
