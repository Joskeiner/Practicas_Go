package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func AuthCheck() Middlewere {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := false
			fmt.Println("checking Authentication")
			if !flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middlewere {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
