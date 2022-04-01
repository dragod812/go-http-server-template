package main

import (
	_ "net/http/pprof"

	"github.com/dragod812/go-http-server-template/register"
)

func main() {
	register.RegisterHTTPHandlers()
}
