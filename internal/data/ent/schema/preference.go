package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Preference holds the schema definition for the Preference entity.
type Preference struct {
	ent.Schema
}

// Fields of the Preference.
func (Preference) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Comment("关键值").Immutable(),
		field.String("value").Comment("值"),
	}
}

func (Preference) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Preference.
func (Preference) Edges() []ent.Edge {
	return nil
}
