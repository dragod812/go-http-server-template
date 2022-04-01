package handler

import "github.com/dragod812/go-http-server-template/entity"

type ReadPageRequest struct {
	PageName string
}

type ReadPageResponse struct {
	Page *entity.Page
}

func (h *handlerImpl) ReadPage(request *ReadPageRequest) (*ReadPageResponse, error) {

	page, err := h.pageComponent.ReadPage(request.PageName)
	if err != nil {
		return nil, err
	}

	return &ReadPageResponse{
		Page: page,
	}, nil
}
