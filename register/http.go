package register

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/dragod812/go-http-server-template/entity"
	"github.com/dragod812/go-http-server-template/handler"
	"github.com/dragod812/go-http-server-template/lib/internalerrors"
	"github.com/dragod812/go-http-server-template/lib/templates"
	"github.com/dragod812/go-http-server-template/mapper"
)

const (
	ReadPath  = "/page/read/"
	WritePath = "/page/write/"
)

var pageHandler = handler.NewPageHandler()

var regExpURLValidator = regexp.MustCompile("^/page/(read|write)/([a-zA-Z0-9]+)$")

func ReadPageHttpHandler(response http.ResponseWriter, request *http.Request) {
	// cpuprofile, _ := os.Create("readcpuprofile")
	// pprof.StartCPUProfile(cpuprofile)
	// defer pprof.StopCPUProfile()
	pageName, err := getPageNameFromURL(request.URL.Path)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	readPageResponse, err := pageHandler.ReadPage(mapper.AdaptPageNameToReadPageRequest(pageName))

	if err != nil {
		if _, ok := err.(*internalerrors.PageNotFoundError); ok {
			http.Error(response, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(response, fmt.Sprintf("Read page operation failed with %s", err.Error()), http.StatusInternalServerError)
		return
	}
	// fmt.Printf("Read page response: %s", readPageResponse.Page)

	err = templates.RenderTemplate("page", readPageResponse.Page, response)
	if err != nil {
		http.Error(response, fmt.Sprintf("Error rendering template: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}
func WritePageHttpHandler(response http.ResponseWriter, request *http.Request) {
	// cpuprofile, _ := os.Create("writecpuprofile")
	// pprof.StartCPUProfile(cpuprofile)
	// defer pprof.StopCPUProfile()

	pageName, err := getPageNameFromURL(request.URL.Path)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	var page entity.Page
	err = json.NewDecoder(request.Body).Decode(&page)
	// fmt.Printf("Page: %s\n", page)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if pageName != page.PageName {
		http.Error(response, "page name not correct", http.StatusBadRequest)
		return
	}

	err = pageHandler.WritePage(&page)

	if err != nil {
		http.Error(response, "Server Error: Unable to write the page", http.StatusInternalServerError)
		return
	}
}

func getPageNameFromURL(url string) (string, error) {
	// regExpURLValidator.
	m := regExpURLValidator.FindStringSubmatch(url)
	if m == nil {
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil
}

func RegisterHTTPHandlers() {

	templates.NewTemplateCache(templates.TemplateFile)

	http.HandleFunc(ReadPath, ReadPageHttpHandler)
	log.Default().Print("Registered ReadPageHttpHandler")
	http.HandleFunc(WritePath, WritePageHttpHandler)
	log.Default().Print("Registered WritePageHttpHandler")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
