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
		app.clientError(w, http.StatusBadRequest, "Request must have limit query parameter")
		return
	}

	limitInt, err := strconv.Atoi(limit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Limit query parameter has to be integer")
		return
	}

	data, err := service.GetAllCarsWithLimit(app.repo, int64(limitInt))
	if err != nil {
		app.serverError(w, err, "Error occurred while reading data from database", http.StatusInternalServerError)
		return
	}

	(w).Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(data)

	_, err = w.Write(res)
	if err != nil {
		app.serverError(w, err, "Error occurred while writing data to body of the response", http.StatusInternalServerError)
		return
	}
}

func (app *application) GetMyOpenDataWithYearFilter(w http.ResponseWriter, r *http.Request) {
	year, ok := r.URL.Query()["year"]

	if !ok || len(year[0]) < 1 {
		app.clientError(w, http.StatusBadRequest, "Request must have year query parameter")
		return
	}

	yearInt, err := strconv.Atoi(year[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Limit query parameter has to be integer")
		return
	}

	data, err := service.GetAllCarsFromSomeYear(app.repo, int64(yearInt))
	if err != nil {
		app.serverError(w, err, "Error occurred while reading data from database", http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		app.serverError(w, err, "Error occurred while writing data to body of the response", http.StatusBadRequest)
		return
	}
}

func (app *application) GetMyOpenDataWithSpeedFilter(w http.ResponseWriter, r *http.Request) {
	lowerLimit, ok := r.URL.Query()["lowerLimit"]

	if !ok || len(lowerLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest, "Request must have lowerLimit query parameter")
		return
	}

	lowerLimitInt, err := strconv.Atoi(lowerLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Lower limit query parameter has to be integer")
		return
	}

	upperLimit, ok := r.URL.Query()["upperLimit"]

	if !ok || len(upperLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest, "Request must have upperLimit query parameter")
		return
	}

	upperLimitInt, err := strconv.Atoi(upperLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Lower limit query parameter has to be integer")
		return
	}

	data, err := service.GetAllCarsWithSpeedLimit(app.repo, lowerLimitInt, upperLimitInt)
	if err != nil {
		app.serverError(w, err, "Error occurred while reading data from database", http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		app.serverError(w, err, "Error occurred while writing data to body of the response", http.StatusInternalServerError)
		return
	}
}

func (app *application) GetMyOpenDataWithPowerOutputFilter(w http.ResponseWriter, r *http.Request) {
	lowerLimit, ok := r.URL.Query()["lowerLimit"]

	if !ok || len(lowerLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest, "Request must have limit query parameter")
		return
	}

	lowerLimitInt, err := strconv.Atoi(lowerLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Lower lower limit query parameter has to be integer")
		return
	}

	upperLimit, ok := r.URL.Query()["upperLimit"]

	if !ok || len(upperLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest, "Request must have upper limit query parameter")
		return
	}

	upperLimitInt, err := strconv.Atoi(upperLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Upper limit query parameter has to be integer")
		return
	}

	data, err := service.GetAllCarsWithPowerOutputLimit(app.repo, lowerLimitInt, upperLimitInt)
	if err != nil {
		app.serverError(w, err, "Error occurred while reading data from database", http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		app.serverError(w, err, "Error occurred while writing data to body of the response", http.StatusInternalServerError)
		return
	}
}

func (app *application) GetMyOpenDataWithBootSpaceFilter(w http.ResponseWriter, r *http.Request) {
	lowerLimit, ok := r.URL.Query()["lowerLimit"]

	if !ok || len(lowerLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest, "Request must have lower limit query parameter")
		return
	}

	lowerLimitInt, err := strconv.Atoi(lowerLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Lower limit query parameter has to be integer")
		return
	}

	upperLimit, ok := r.URL.Query()["upperLimit"]

	if !ok || len(upperLimit[0]) < 1 {
		app.clientError(w, http.StatusBadRequest, "Request must have upper limit query parameter")
		return
	}

	upperLimitInt, err := strconv.Atoi(upperLimit[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Upper limit query parameter has to be integer")
		return
	}

	data, err := service.GetAllCarsWithBootSpaceLimit(app.repo, lowerLimitInt, upperLimitInt)
	if err != nil {
		app.serverError(w, err, "Error occurred while reading data from database", http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		app.serverError(w, err, "Error occurred while writing data to body of the response", http.StatusInternalServerError)
		return
	}
}

func (app *application) CreateNewCar(w http.ResponseWriter, r *http.Request) {

	var data storage.MyOpenDataStructDIO

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Error occurred while decoding sent data")
		return
	}

	oid, err := storage.InsertCar(app.repo, data)
	if err != nil {
		app.serverError(w, err, "Error occurred while inserting new car", http.StatusInternalServerError)
		return
	}

	insertedCar, err := storage.GetCarWithId(app.repo, oid)

	response := service.ApiResponseCRUD{
		Status:   "OK",
		Message:  "Car inserted",
		Response: service.ModelOneCarDataForResponse(insertedCar),
	}
	res, _ := json.Marshal(response)

	(w).WriteHeader(http.StatusOK)
	(w).Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}

func (app *application) DeleteCar(w http.ResponseWriter, r *http.Request) {

	id, ok := r.URL.Query()["id"]

	if !ok || len(id[0]) < 1 {
		app.clientError(w, http.StatusBadRequest, "Request must have id query parameter")
		return
	}

	oid, err := primitive.ObjectIDFromHex(id[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "id query parameter has to be objectID from MonogoDB")
		return
	}

	err = storage.DeleteCar(app.repo, oid)
	if err != nil {
		app.serverError(w, err, "Could not find car with given id", http.StatusNotFound)
		return
	}

	response := service.ApiResponse{
		Status:  "OK",
		Message: "Car deleted",
	}
	res, _ := json.Marshal(response)
	(w).Header().Set("Content-Type", "application/json")
	(w).WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (app *application) UpdateCar(w http.ResponseWriter, r *http.Request) {

	id, ok := r.URL.Query()["id"]

	if !ok || len(id[0]) < 1 {
		app.clientError(w, http.StatusBadRequest, "Request must have id query parameter")
		return
	}

	oid, err := primitive.ObjectIDFromHex(id[0])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "id query parameter has to be objectID from MonogoDB")
		return
	}

	var data storage.MyOpenDataStructDIO

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Error occurred while decoding sent data")
		return
	}

	err = storage.UpdateCar(app.repo, oid, data)
	if err != nil {
		app.serverError(w, err, "Error occurred while updating car in database"+err.Error(), http.StatusBadRequest)
		return
	}

	response := service.ApiResponseCRUD{
		Status:  "OK",
		Message: "Car updated",
	}

	car, err := storage.GetCarWithId(app.repo, oid)

	if err == nil {
		response.Response = service.ModelOneCarDataForResponse(car)
	}

	(w).WriteHeader(http.StatusOK)
	res, _ := json.Marshal(response)
	(w).Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}
