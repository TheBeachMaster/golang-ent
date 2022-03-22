package usecase

import (
	"context"

	"com.thebeachmaster/entexample/config"
	"com.thebeachmaster/entexample/internal/song"
	"com.thebeachmaster/entexample/internal/song/models"
)

type songusecase struct {
	cfg  *config.Config
	repo song.SongDBRepository
}

func NewSongUseCase(c *config.Config, r song.SongDBRepository) song.SongUseCase {
	return &songusecase{cfg: c, repo: r}
}

func (s *songusecase) AddNewSong(context context.Context, songInfo *models.CreateSong) (*models.CreateSongResponse, error) {
	return s.repo.Create(context, songInfo)
}

func (s *songusecase) AddSongLyrics(context context.Context, lyricsInfo *models.UpdateSongLyricsFileURL) (*models.UpdateSongResponse, error) {
	return s.repo.UpdateLyricsURL(context, lyricsInfo)
}

func (s *songusecase) ListAllSongs(context context.Context) (*models.SongListResponse, error) {
	return s.repo.GetAllSongs(context)
}

func (s *songusecase) FetchSong(context context.Context, hash string) (*models.SongDetailsResponse, error) {
	return s.repo.GetSong(context, hash)
}
