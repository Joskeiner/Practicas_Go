package main

import (
	"net/http"
)

// manejador de rutas
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path, metodo string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, metodoExsist := r.rules[path][metodo]
	return handler, exist, metodoExsist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, resquest *http.Request) {
	handler, metodoExist, exist := r.FindHandler(resquest.URL.Path, resquest.Method)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !metodoExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	handler(w, resquest)
}
