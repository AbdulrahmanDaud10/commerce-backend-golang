package repository

import (
	"context"

	"github.com/AbdulrahmanDaud10/commerce-backend-golang/pkg/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductStore struct {
	db   *mongo.Database
	call string
}

func NewMongoProductStore(db *mongo.Database) *MongoProductStore {
	return &MongoProductStore{
		db:   db,
		call: "products",
	}
}

func (s *MongoProductStore) Insert(ctx context.Context, p *api.Product) error {
	res, err := s.db.Collection(s.call).InsertOne(ctx, p)
	if err != nil {
		return err
	}

	p.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return err
}

func (s *MongoProductStore) GetAll(ctx context.Context) ([]*api.Product, error) {
	cursor, err := s.db.Collection(s.call).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}

	products := []*api.Product{}
	err = cursor.All(ctx, &products)

	return products, err
}

func (s *MongoProductStore) GetByID(ctx context.Context, id string) (*api.Product, error) {
	var (
		ObjectID, _ = primitive.ObjectIDFromHex(id)
		res         = s.db.Collection(s.call).FindOne(ctx, bson.M{"_id": ObjectID})
		p           = &api.Product{}
		err         = res.Decode(p)
	)
	return p, err
}
