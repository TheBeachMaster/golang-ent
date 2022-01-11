package models

type (
	CreateSong struct {
		Name          string
		FileURL       string
		Album         string
		LyricsFileURL string
	}
	CreateSongResponse struct {
		SongID string
	}
)
