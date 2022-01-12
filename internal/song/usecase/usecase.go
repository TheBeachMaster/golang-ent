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
	res, err := s.repo.Create(context, songInfo)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *songusecase) AddSongLyrics(context context.Context, lyricsInfo *models.UpdateSongLyricsFileURL) (*models.UpdateSongResponse, error) {
	res, err := s.repo.UpdateLyricsURL(context, lyricsInfo)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *songusecase) ListAllSongs(context context.Context) (*models.SongListResponse, error) {
	res, err := s.repo.GetAllSongs(context)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *songusecase) FetchSong(context context.Context, hash string) (*models.SongDetailsResponse, error) {
	res, err := s.repo.GetSong(context, hash)
	if err != nil {
		return nil, err
	}
	return res, nil
}
