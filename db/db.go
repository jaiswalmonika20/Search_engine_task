package db

import (
	"context"

	"github.com/module_page/models"
)

type DBClient interface {
	GetAllPages(ctx context.Context) (resp []*models.Page, err error)
	AddPage(res *models.Page) error
}
