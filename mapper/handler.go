package mapper

import "github.com/dragod812/go-http-server-template/handler"

func AdaptPageNameToReadPageRequest(pageName string) *handler.ReadPageRequest {
	return &handler.ReadPageRequest{
		PageName: pageName,
	}
}
