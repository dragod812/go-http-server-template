package pages

import (
	"github.com/dragod812/go-http-server-template/entity"
	"github.com/dragod812/go-http-server-template/repository"
	"github.com/dragod812/go-http-server-template/repository/files"
)

type PagesComponent struct {
	pageRepository repository.Pages
}

func NewPageComponent() PagesComponent {
	return PagesComponent{
		pageRepository: files.NewPageFileRepository(),
	}
}

func (c PagesComponent) ReadPage(pageName string) (*entity.Page, error) {
	return c.pageRepository.Read(pageName)
}

func (c PagesComponent) WritePage(page *entity.Page) error {
	return c.pageRepository.Write(*page)
}
