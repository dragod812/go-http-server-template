package handler

import (
	"github.com/dragod812/go-http-server-template/component/pages"
	"github.com/dragod812/go-http-server-template/entity"
)

type PageHandler interface {
	WritePage(page *entity.Page) error
	ReadPage(request *ReadPageRequest) (*ReadPageResponse, error)
}

type handlerImpl struct {
	pageComponent pages.PagesComponent
}

func NewPageHandler() PageHandler {
	return &handlerImpl{
		pageComponent: pages.NewPageComponent(),
	}
}
