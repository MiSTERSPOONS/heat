package main

import "net/http"

func (app *Config) fire(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to the fire service"))
}
