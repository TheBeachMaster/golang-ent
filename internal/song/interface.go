package song

import (
	"context"

	"com.thebeachmaster/entexample/internal/song/models"
)

type SongDBRepository interface {
	Create(ctx context.Context, songInfo *models.CreateSong) (*models.CreateSongResponse, error)
	UpdateLyricsURL(ctx context.Context, songInfo *models.UpdateSongLyricsFileURL) (*models.UpdateSongResponse, error)
}
