package main

import (
	"github.com/gorilla/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/cars", http.HandlerFunc(app.GetMyOpenData))
	mux.Post("/cars", http.HandlerFunc(app.CreateNewCar))
	mux.Put("/cars", http.HandlerFunc(app.UpdateCar))
	mux.Delete("/cars", http.HandlerFunc(app.DeleteCar))
	mux.Get("/year", http.HandlerFunc(app.GetMyOpenDataWithYearFilter))
	mux.Get("/speed", http.HandlerFunc(app.GetMyOpenDataWithSpeedFilter))
	mux.Get("/powerOutput", http.HandlerFunc(app.GetMyOpenDataWithPowerOutputFilter))
	mux.Get("/bootSpace", http.HandlerFunc(app.GetMyOpenDataWithBootSpaceFilter))
	return mux
}
