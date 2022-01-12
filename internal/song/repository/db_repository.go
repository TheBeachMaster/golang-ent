package repository

import (
	"context"

	"com.thebeachmaster/entexample/ent"
	"com.thebeachmaster/entexample/internal/song"
	"com.thebeachmaster/entexample/internal/song/models"
)

type songRepository struct {
	client *ent.SongClient
}

func NewSongRespository(client *ent.Client) song.SongDBRepository {
	return &songRepository{client: client.Song}
}

func (s *songRepository) Create(context context.Context, songInfo *models.CreateSong) (*models.CreateSongResponse, error) {
	song, err := s.client.Create().SetName(songInfo.Name).SetFileURL(songInfo.FileURL).SetLyricsURL(songInfo.LyricsFileURL).Save(context)
	if err != nil {
		return nil, err
	}
	res := &models.CreateSongResponse{}
	res.SongID = string(song.ID.String())
	res.CreatedAt = song.CreatedAt
	return res, nil
}

func (s *songRepository) UpdateLyricsURL(context context.Context, songInfo *models.UpdateSongLyricsFileURL) (*models.UpdateSongResponse, error) {
	updateSong, err := s.client.Update().SetLyricsURL(songInfo.FileURL).Save(context)
	if err != nil {
		return nil, err
	}
	res := &models.UpdateSongResponse{
		SongsUpdated: updateSong,
	}
	return res, nil
}
