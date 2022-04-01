package handler

import "github.com/dragod812/go-http-server-template/entity"

func (h handlerImpl) WritePage(page *entity.Page) error {
	return h.pageComponent.WritePage(page)
}
