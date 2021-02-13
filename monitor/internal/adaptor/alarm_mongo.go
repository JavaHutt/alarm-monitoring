package adaptor

import (
	"context"
	"log"
	"monitor/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const techniqueCollection = "test_technique"

// TechMongo is an adaptor for techniques
type TechMongo struct {
	db             *mongo.Database
	collectionName string
}

// NewTechMongo is a constructor
func NewTechMongo(db *mongo.Database) *TechMongo {
	return &TechMongo{
		db:             db,
		collectionName: techniqueCollection,
	}
}

// CreateTech creates new record in technique collection
func (a *TechMongo) CreateTech(ctx context.Context, alarm model.Alarm) error {
	_, err := a.getCollection().InsertOne(ctx, alarm)
	if err != nil {
		return err
	}

	log.Println("Created new technique")
	return nil
}

// UpdateOngoingTech updates record in technique collection
func (a *TechMongo) UpdateOngoingTech(ctx context.Context, id primitive.ObjectID, fields model.Alarm) error {
	_, err := a.getCollection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"crit":      fields.Crit,
		"last_msg":  fields.LastMsg,
		"last_time": fields.LastTime,
	}})
	if err != nil {
		return err
	}
	log.Println("Updated existing ongoing technique")
	return nil
}

// UpdateResolvedTech updates and resolves record in technique collection
func (a *TechMongo) UpdateResolvedTech(ctx context.Context, id primitive.ObjectID, fields model.Alarm) error {
	_, err := a.getCollection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"crit":      fields.Crit,
		"last_msg":  fields.LastMsg,
		"last_time": fields.LastTime,
		"status":    model.Resolved,
	}})
	if err != nil {
		return err
	}
	log.Println("Updated existing resolved technique")
	return nil
}

// GetAllTechniques returns all records in technique collection
func (a *TechMongo) GetAllTechniques(ctx context.Context) ([]model.Alarm, error) {
	result := new([]model.Alarm)

	cursor, err := a.getCollection().Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err = cursor.All(ctx, result); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return *result, nil
}

// GetOngoingTechByComponentName returns technique by component/resource couple
// if technique is not found returns false as second value
func (a *TechMongo) GetOngoingTechByComponentName(
	ctx context.Context,
	component string,
	resource string,
) (*model.Alarm, error) {
	result := new(model.Alarm)

	err := a.getCollection().FindOne(ctx, bson.M{
		"component": component,
		"resource":  resource,
		"status":    model.Ongoing,
	}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Record does not exist")
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	log.Println("Record exists")
	return result, nil
}

func (a *TechMongo) getCollection() *mongo.Collection {
	return a.db.Collection(a.collectionName)
}
