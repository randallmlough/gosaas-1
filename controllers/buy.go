package controllers

import (
	"net/http"

	"github.com/dstpierre/gosaas/data/model"
	"github.com/dstpierre/gosaas/engine"
)

// User handles everything related to the /user requests
type Buy struct{}

func newBuy() *engine.Route {
	var b interface{} = Buy{}
	return &engine.Route{
		Logger:      true,
		EnforceRate: false,
		MinimumRole: model.RolePublic,
		Handler:     b.(http.Handler),
	}
}

func (b Buy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "" {
		b.home(w, r)
	}
}

func (b Buy) home(w http.ResponseWriter, r *http.Request) {
	var data = new(struct {
		Msg string
	})
	data.Msg = "this is much cleaner"

	engine.ServePage(w, r, "index-dev.html", data)
}
