package schema

import (
	"net/url"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Song holds the schema definition for the Song entity.
type Song struct {
	ent.Schema
}

// Fields of the Song.
func (Song) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("hash").NotEmpty(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").Default(time.Now),
		field.String("name").NotEmpty(),
		field.String("file_url"),
		field.String("album_name").Optional(),
		field.JSON("lyrics_url", &url.URL{}).Optional(),
		field.String("thumbnail_url").Optional(),
	}
}

// Edges of the Song.
func (Song) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playlist", Playlist.Type).Ref("song"),
		edge.From("artist", Artist.Type).Ref("song"),
		edge.From("album", Album.Type).Ref("song_list").Unique(),
		edge.From("label", Label.Type).Ref("single").Unique(),
	}
}
