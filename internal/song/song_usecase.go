package song

import (
	"context"

	"com.thebeachmaster/entexample/internal/song/models"
)

type SongUseCase interface {
	AddNewSong(context context.Context, songInfo *models.CreateSong) (*models.CreateSongResponse, error)
	AddSongLyrics(context context.Context, lyricsInfo *models.UpdateSongLyricsFileURL) (*models.UpdateSongResponse, error)
	ListAllSongs(context context.Context) (*models.SongListResponse, error)
	FetchSong(context context.Context, songHash string) (*models.SongDetailsResponse, error)
}
