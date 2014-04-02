package api

import (
	"github.com/codegangsta/inject"
	"github.com/go-martini/martini"
	"net/http"
)

type GiAPI interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type giAPI struct {
	inject.Injector
	router martini.Router
}

func (api *giAPI) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	api.Invoke(api.router.Handle)
}

func New(path string) GiAPI {
	r := martini.NewRouter()
	return &giAPI{Injector: inject.New(), router: r}
}
