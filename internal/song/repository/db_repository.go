package repository

import (
	"context"

	"com.thebeachmaster/entexample/common"
	"com.thebeachmaster/entexample/ent"
	songpredicate "com.thebeachmaster/entexample/ent/song"
	"com.thebeachmaster/entexample/internal/song"
	"com.thebeachmaster/entexample/internal/song/models"
	"com.thebeachmaster/entexample/utils"
)

type songRepository struct {
	client *ent.SongClient
}

func NewSongRespository(client *ent.Client) song.SongDBRepository {
	return &songRepository{client: client.Song}
}

func (s *songRepository) Create(context context.Context, songInfo *models.CreateSong) (*models.CreateSongResponse, error) {
	// let's create the entity hash first, that we'll use to query - this is "IDEALLY" unique
	hashInput := &common.HashInput{
		Name:     songInfo.Name,
		Metadata: songInfo.FileURL,
	}
	hashValue, err := utils.GenerateEntityHash(hashInput)
	if err != nil {
		return nil, err
	}
	song, err := s.client.Create().SetName(songInfo.Name).SetFileURL(songInfo.FileURL).SetThumbnailURL(songInfo.ThumbnailURL).SetLyricsURL(songInfo.LyricsFileURL).SetHash(hashValue).Save(context)
	if err != nil {
		return nil, err
	}
	res := &models.CreateSongResponse{}
	res.SongHash = song.Hash
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

func (s *songRepository) GetAllSongs(context context.Context) (*models.SongListResponse, error) {
	list, err := s.client.Query().All(context)
	if err != nil {
		return nil, err
	}
	songs := make([]models.SongInfo, 0)
	for _, v := range list {
		song := &models.SongInfo{
			Hash:      v.Hash,
			Album:     v.AlbumName,
			Thumbnail: v.ThumbnailURL,
			Name:      v.Name,
			Created:   v.CreatedAt,
		}
		songs = append(songs, *song)
	}

	res := &models.SongListResponse{
		Songs: songs,
	}
	return res, nil
}

func (s *songRepository) GetSong(context context.Context, hash string) (*models.SongDetailsResponse, error) {
	song, err := s.client.Query().Where(songpredicate.HashEQ(hash)).Only(context)
	if err != nil {
		return nil, err
	}

	res := &models.SongDetailsResponse{
		Hash:      song.Hash,
		Album:     song.AlbumName,
		Thumbnail: song.ThumbnailURL,
		Lyrics:    song.LyricsURL,
		Name:      song.Name,
		Created:   song.CreatedAt,
		Updated:   song.UpdatedAt,
	}
	return res, nil
}
