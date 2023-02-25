package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

// ! importante siempre usar * y & para evitar que se esten utilizando copias de las struct
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}
func (e *Server) AddMidleware(f http.HandlerFunc, middlewares ...Middlewere) http.HandlerFunc {
	for _, v := range middlewares {
		f = v(f)
	}
	return f

}
func (e *Server) handle(path string, handler http.HandlerFunc) {
	e.router.rules[path] = handler
}

func (e *Server) Listen() error {
	http.Handle("/", e.router)

	err := http.ListenAndServe(e.port, nil)

	if err != nil {
		return err
	}
	return nil
}
