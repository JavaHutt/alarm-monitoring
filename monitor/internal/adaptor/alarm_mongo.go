package adaptor

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TechMongo struct {
	db *mongo.Database
}

// NewTechMongo is a constructor
func NewTechMongo(db *mongo.Database) *TechMongo {
	return &TechMongo{db: db}
}

func (a *TechMongo) CreateTech() (int, error) {
	return 4, nil
}

func (a *TechMongo) UpdateTech(id int) (string, error) {
	return "4", nil
}

func (a *TechMongo) GetAllTechniques(ctx context.Context) (string, error) {
	fmt.Println(a.db.ListCollectionNames(ctx, bson.M{}))
	return "3", nil
}

// GetTechByComponentName returns technique by component/resource couple
// if technique is not found returns false as second value
func (a *TechMongo) GetTechByComponentName(component string, resource string) (string, bool) {
	return "4", false
}
