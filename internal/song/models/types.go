package models

import "time"

type (
	CreateSong struct {
		Name          string `json:"name"`
		FileURL       string `json:"fileUrl"`
		Album         string `json:"album"`
		LyricsFileURL string `json:"lyricsFile"`
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
