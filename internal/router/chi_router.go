package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	r = chi.NewRouter()
)

func init() {
	r.Use(middleware.Logger)
}

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Get(uri, handler)
}

func (*chiRouter) POST(uri string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Post(uri, handler)
}

func (*chiRouter) PUT(uri string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Put(uri, handler)
}

func (*chiRouter) PATCH(uri string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Patch(uri, handler)
}

func (*chiRouter) DELETE(uri string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Delete(uri, handler)
}

func (*chiRouter) GetParam(req *http.Request, key string) any {
	return chi.URLParam(req, key)
}
