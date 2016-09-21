package mux

import "net/http"

type tree interface {
	Get()
	Set()
}

type Handler http.Handler

type Middleware func(w http.ResponseWriter, r *http.Request, next HandlerFunc)

func (mw *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw(w, r, nil)
}

type Mux struct {
	tree *node
}

func New() *Mux {
	return &Mux{}
}

func (m *Mux) Handle(path string, mw []Middleware, h Handler) {
}
