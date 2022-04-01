package files

import (
	"fmt"
	"os"

	"github.com/dragod812/go-http-server-template/entity"
	"github.com/dragod812/go-http-server-template/lib/internalerrors"
	"github.com/dragod812/go-http-server-template/repository"
)

type pagesFileImpl struct{}

func NewPageFileRepository() repository.Pages {
	return &pagesFileImpl{}
}

func (p *pagesFileImpl) Read(pageName string) (*entity.Page, error) {
	fileName := pageName + ".txt"
	pageData, err := os.ReadFile(fileName)
	if err != nil {
		_, ok := err.(*os.PathError)
		if ok {
			fmt.Printf("Error: No file found with the file name %s", fileName)
			return nil, &internalerrors.PageNotFoundError{
				PageName: pageName,
			}
		}
		return nil, err
	}
	return &entity.Page{
		PageName: pageName,
		PageData: pageData,
	}, nil
}

func (p *pagesFileImpl) Write(page entity.Page) error {
	fileName := page.PageName + ".txt"
	return os.WriteFile(fileName, page.PageData, 0600)
}
