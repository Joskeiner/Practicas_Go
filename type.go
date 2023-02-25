package main

import "net/http"

type Middlewere func(http.HandlerFunc) http.HandlerFunc
