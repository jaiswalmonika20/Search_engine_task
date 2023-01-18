package services

import (
	"context"
	"errors"

	"github.com/module_page/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PageServiceImpl struct {
	pagescollection *mongo.Collection
	ctx             context.Context
}

func NewPageService(pagescollection *mongo.Collection, ctx context.Context) PageService {
	return &PageServiceImpl{
		pagescollection: pagescollection,
		ctx:             ctx,
	}
}

func (psi *PageServiceImpl) GetAllPages() ([]*models.Page, error) {
	var pages []*models.Page
	cursor, err := psi.pagescollection.Find(psi.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(psi.ctx) {
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

	cursor.Close(psi.ctx)

	if len(pages) == 0 {
		return nil, errors.New("pages not found")
	}
	return pages, nil
}
