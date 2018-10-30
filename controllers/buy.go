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
		MinimumRole: model.RoleUser,
		Handler:     b.(http.Handler),
	}
}

func (b Buy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO
}
