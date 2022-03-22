package delivery

import "com.thebeachmaster/entexample/internal/song"

type Delivery struct {
	Songs interface{ song.SongHTTPRoutes }
}
