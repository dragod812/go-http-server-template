package repository

import "github.com/dragod812/go-http-server-template/entity"

type Pages interface {
	Write(page entity.Page) error
	Read(pageName string) (*entity.Page, error)
}
