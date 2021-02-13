package adaptor

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Adaptor struct {
	TechAdaptor
}

// NewAdaptor is a constructor
func NewAdaptor(db *mongo.Database) *Adaptor {
	return &Adaptor{
		TechAdaptor: NewTechMongo(db),
	}
}

type TechAdaptor interface {
	CreateTech() (int, error)
	UpdateTech(id int) (string, error)
	GetAllTechniques(ctx context.Context) (string, error)
	GetTechByComponentName(component string, resource string) (string, bool)
}
