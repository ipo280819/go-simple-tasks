package router

import (
	"go-tasks/constants"
	"net/http"
)

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	PUT(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

func NewRouter(typeRouter string) Router {
	switch typeRouter {
	case constants.MUX:
		return newMuxRouter()
	}
	return nil
}
