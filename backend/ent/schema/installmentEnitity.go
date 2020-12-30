package schema

import "github.com/facebookincubator/ent"

// Room holds the schema definition for the Room entity.
type InstallmentEntity struct {
	ent.Schema
}

// Fields of the Room.
func (InstallmentEntity) Fields() []ent.Field {
	return nil
}

// Edges of the Room.
func (InstallmentEntity) Edges() []ent.Edge {
	return nil
}
