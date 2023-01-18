package services

import "github.com/module_page/models"

type PageService interface {
	GetAllPages() ([]*models.Page, error)
}
