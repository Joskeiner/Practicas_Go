package main

import (
	"net/http"
)

// manejador de rutas
type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]

	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, resquest *http.Request) {
	handler, exist := r.FindHandler(resquest.URL.Path)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, resquest)
}