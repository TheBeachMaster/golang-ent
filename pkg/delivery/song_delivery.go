package delivery

import (
	"com.thebeachmaster/entexample/internal/song"
	songRepo "com.thebeachmaster/entexample/internal/song/repository"
	songHTTPHandler "com.thebeachmaster/entexample/internal/song/routes"
	songUseCase "com.thebeachmaster/entexample/internal/song/usecase"
)

func (d *deliveryRegistry) NewSongDelivery() song.SongHTTPRoutes {
	songRepo := songRepo.NewSongRespository(d.client)
	songUsecase := songUseCase.NewSongUseCase(d.cfg, songRepo)
	return songHTTPHandler.NewSongHTTPHandler(d.cfg, songUsecase)
}
