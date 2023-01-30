package services

import (
	"context"
	"errors"

	"github.com/module_page/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PageServiceImpl struct {
	pagescollection *mongo.Collection
	context         context.Context
}

func NewPageService(pagescollection *mongo.Collection, context context.Context) PageService {
	return &PageServiceImpl{
		pagescollection: pagescollection,
		context:         context,
	}
}

func (u *PageServiceImpl) AddPage(page *models.Page) error {
	_, err := u.pagescollection.InsertOne(u.context, page)
	return err
}

func (psi *PageServiceImpl) GetAllPages() ([]*models.Page, error) {
	var pages []*models.Page
	cursor, err := psi.pagescollection.Find(psi.context, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(psi.context) {
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

	cursor.Close(psi.context)

	if len(pages) == 0 {
		return nil, errors.New("pages not found")
	}
	return pages, nil
}
