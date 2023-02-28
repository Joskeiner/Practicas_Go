package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "COME HOME ")
}

func PostResquest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error:%v", err)
		return
	}
	fmt.Fprintf(w, "payload %v \n", metadata)
}

func UserPostResquest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error:%v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(response)
}
