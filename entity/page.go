package entity

import "fmt"

type Page struct {
	PageName string
	PageData []byte
}

func (p Page) Render() string {
	return fmt.Sprintf("<h1>%s<\\h1><p>%s<\\p><br>", p.PageName, string(p.PageData))
}
