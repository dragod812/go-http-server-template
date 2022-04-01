package internalerrors

import (
	"fmt"
)

type PageNotFoundError struct {
	PageName string
}

func (p *PageNotFoundError) Error() string {
	return fmt.Sprintf("page not found with name %s", p.PageName)
}
