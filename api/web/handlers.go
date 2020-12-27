package main

import (
	"encoding/json"
	"github.com/Edrudo/otvoreno_lab/service"
	"github.com/Edrudo/otvoreno_lab/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

func (app *application) GetMyOpenData(w http.ResponseWriter, r *http.Request) {
	limit, ok := r.URL.Query()["limit"]

	if !ok || len(limit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	limitInt, err := strconv.Atoi(limit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	data, err := service.GetAllCarsWithLimit(app.repo, int64(limitInt))
	if err != nil {
		app.serverError(w, err)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(res)
}

func (app *application) GetMyOpenDataWithYearFilter(w http.ResponseWriter, r *http.Request) {
	year, ok := r.URL.Query()["year"]

	if !ok || len(year[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	yearInt, err := strconv.Atoi(year[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	data, err := service.GetAllCarsFromSomeYear(app.repo, int64(yearInt))
	if err != nil {
		app.serverError(w, err)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(res)
}

func (app *application) GetMyOpenDataWithSpeedFilter(w http.ResponseWriter, r *http.Request) {
	lowerLimit, ok := r.URL.Query()["lowerLimit"]

	if !ok || len(lowerLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	lowerLimitInt, err := strconv.Atoi(lowerLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	upperLimit, ok := r.URL.Query()["upperLimit"]

	if !ok || len(upperLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	upperLimitInt, err := strconv.Atoi(upperLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	data, err := service.GetAllCarsWithSpeedLimit(app.repo, lowerLimitInt, upperLimitInt)
	if err != nil {
		app.serverError(w, err)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(res)
}

func (app *application) GetMyOpenDataWithPowerOutputFilter(w http.ResponseWriter, r *http.Request) {
	lowerLimit, ok := r.URL.Query()["lowerLimit"]

	if !ok || len(lowerLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	lowerLimitInt, err := strconv.Atoi(lowerLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	upperLimit, ok := r.URL.Query()["upperLimit"]

	if !ok || len(upperLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	upperLimitInt, err := strconv.Atoi(upperLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	data, err := service.GetAllCarsWithPowerOutputLimit(app.repo, lowerLimitInt, upperLimitInt)
	if err != nil {
		app.serverError(w, err)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(res)
}

func (app *application) GetMyOpenDataWithBootSpaceFilter(w http.ResponseWriter, r *http.Request) {
	lowerLimit, ok := r.URL.Query()["lowerLimit"]

	if !ok || len(lowerLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	lowerLimitInt, err := strconv.Atoi(lowerLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	upperLimit, ok := r.URL.Query()["upperLimit"]

	if !ok || len(upperLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	upperLimitInt, err := strconv.Atoi(upperLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	data, err := service.GetAllCarsWithBootSpaceLimit(app.repo, lowerLimitInt, upperLimitInt)
	if err != nil {
		app.serverError(w, err)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(res)
}

func (app *application) CreateNewCar(w http.ResponseWriter, r *http.Request) {

	var data storage.MyOpenDataStructDIO

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = storage.InsertCar(app.repo, data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func (app *application) DeleteCar(w http.ResponseWriter, r *http.Request) {

	id, ok := r.URL.Query()["id"]

	if !ok || len(id[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	oid, err := primitive.ObjectIDFromHex(id[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = storage.DeleteCar(app.repo, oid)
	if err != nil {
		app.serverError(w, err)
		return
	}

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func (app *application) UpdateCar(w http.ResponseWriter, r *http.Request) {

	id, ok := r.URL.Query()["id"]

	if !ok || len(id[0]) < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	oid, err := primitive.ObjectIDFromHex(id[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var data storage.MyOpenDataStructDIO

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = storage.UpdateCar(app.repo, oid, data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
}
