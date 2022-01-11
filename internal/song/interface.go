package song

import (
	"context"

	"com.thebeachmaster/entexample/internal/song/models"
)

type DBRepository interface {
	Create(ctx context.Context, songinfo *models.CreateSong) (*models.CreateSongResponse, error)
}
