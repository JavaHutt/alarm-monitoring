package adaptor

import (
	"context"
	"monitor/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Adaptor for data converting
type Adaptor struct {
	TechAdaptor
}

// NewAdaptor is a constructor
func NewAdaptor(db *mongo.Database) *Adaptor {
	return &Adaptor{
		TechAdaptor: NewTechMongo(db),
	}
}

// TechAdaptor is an interface for technique collection
type TechAdaptor interface {
	CreateTech(ctx context.Context, alarm model.Alarm) error
	UpdateOngoingTech(ctx context.Context, id primitive.ObjectID, fields model.Alarm) error
	UpdateResolvedTech(ctx context.Context, id primitive.ObjectID, fields model.Alarm) error
	GetAllTechniques(ctx context.Context) ([]model.Alarm, error)
	GetOngoingTechByComponentName(ctx context.Context, component string, resource string) (*model.Alarm, error)
}
