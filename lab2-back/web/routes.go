package main

import (
	"github.com/gorilla/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/data", http.HandlerFunc(app.GetMyOpenData))
	mux.Get("/filteredJSON", http.HandlerFunc(app.GetFilteredOpenDataInJSON))
	mux.Get("/filteredCSV", http.HandlerFunc(app.GetFilteredOpenDataInCSV))
	mux.Get("/JSON", http.HandlerFunc(app.GetMyOpenDataJSON))
	mux.Get("/CSV", http.HandlerFunc(app.GetMyOpenDataCSV))
	mux.Get("/filteredData", http.HandlerFunc(app.GetFilteredMyOpenData))
	return mux
}
