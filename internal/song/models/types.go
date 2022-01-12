package models

import "time"

type (
	CreateSong struct {
		Name          string `json:"name"`
		FileURL       string `json:"fileUrl"`
		Album         string `json:"album"`
		LyricsFileURL string `json:"lyricsFile"`
		ThumbnailURL  string `json:"thumbnaiFile"`
	}
	CreateSongResponse struct {
		SongHash  string
		CreatedAt time.Time
	}
	UpdateSongResponse struct {
		SongsUpdated int
	}
	SongListResponse struct {
		Songs []SongInfo
	}
	SongInfo struct {
		Hash      string
		Album     string
		Thumbnail string
		Name      string
		Created   time.Time
	}
	SongDetailsResponse struct { // Return Richer Data {Edges}
		Hash      string
		Album     string
		Thumbnail string
		Lyrics    string
		Name      string
		Created   time.Time
		Updated   time.Time
	}
	UpdateSongLyricsFileURL struct {
		FileURL string
	}
)
