package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// XProviderCursor holds the schema definition for the XProviderCursor entity.
type XProviderCursor struct {
	ent.Schema
}

// Fields of the XProviderCursor.
func (XProviderCursor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.Uint64("ChainId"),
		field.Uint64("Height"),
		field.Time("CreatedAt").
			Default(time.Now()),
		field.Time("UpdatedAt").
			Default(time.Now()),
	}
}

// Edges of the XProviderCursor.
func (XProviderCursor) Edges() []ent.Edge {
	return []ent.Edge{}
}
