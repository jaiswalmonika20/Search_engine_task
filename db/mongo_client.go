package db

import (
	"context"

	"github.com/module_page/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	pagescollection *mongo.Collection
}

func NewMongoDB(pagescollection *mongo.Collection) *MongoClient {
	return &MongoClient{
		pagescollection: pagescollection,
	}
}

func (mc *MongoClient) GetAllPages(ctx context.Context) (pages []*models.Page, err error) {

	cursor, err := mc.pagescollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var page models.Page
		err := cursor.Decode(&page)
		if err != nil {
			return nil, err
		}
		pages = append(pages, &page)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)
	return

}

func (mc *MongoClient) AddPage(ctx context.Context, pages *models.Page) error {
	_, err := mc.pagescollection.InsertOne(ctx, pages)
	return err

}
