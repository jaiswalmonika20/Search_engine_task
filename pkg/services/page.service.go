package services

import "github.com/module_page/pkg/models"

type PageService interface {
	GetAllPages() ([]*models.Page, error)
	AddPage(*models.Page) error
}
