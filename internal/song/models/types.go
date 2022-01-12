package models

import "time"

type (
	CreateSong struct {
		Name          string
		FileURL       string
		Album         string
		LyricsFileURL string
	}
	CreateSongResponse struct {
		SongID    string
		CreatedAt time.Time
	}
	UpdateSongResponse struct {
		SongsUpdated int
	}
	UpdateSongLyricsFileURL struct {
		FileURL string
	}
)
