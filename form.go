package api

import (
	"encoding/json"
	"github.com/goincremental/dal"
	"github.com/goincremental/web"
	"net/http"
)

type API interface {
	http.Handler
}
type api struct {
	server web.Server
}

type PostResult struct {
	ID dal.ObjectID `bson:"_id" json:"id"`
}

func New() (API, error) {

	r := web.NewRouter()
	r.HandleFunc("/v1/form", FormPostHandler).Methods("POST")

	s := web.NewServer()
	s.UseHandler(r)

	a := &api{
		server: s,
	}

	return a, nil

}

func (a *api) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	a.server.ServeHTTP(w, req)
}

func FormPostHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := &PostResult{
		ID: dal.NewObjectID(),
	}

	w.WriteHeader(http.StatusCreated)
	j, _ := json.Marshal(result)

	w.Write(j)

}
