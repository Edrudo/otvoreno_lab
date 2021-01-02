package main

import (
	"encoding/json"
	"github.com/Edrudo/otvoreno_lab/service"
	"github.com/Edrudo/otvoreno_lab/storage"
	"net/http"
)

func (app *application) GetMyOpenData(w http.ResponseWriter, r *http.Request) {
	data, err := storage.GetMyOpenData(app.repo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	res, _ := json.Marshal(data)

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(res)

}

func (app *application) GetMyOpenDataJSON(w http.ResponseWriter, r *http.Request) {
	data, err := storage.GetMyOpenData(app.repo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Disposition", "attachment;filename=cars.json")
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteData)

}

func (app *application) GetMyOpenDataCSV(w http.ResponseWriter, r *http.Request) {
	data, err := storage.GetMyOpenData(app.repo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	csvFile, err := service.ExportDataAsCSV(data)
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer csvFile.Close()

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Disposition", "attachment;filename=cars.csv")
	http.ServeFile(w, r, "cars.csv")
}

func (app *application) GetFilteredMyOpenData(w http.ResponseWriter, r *http.Request) {
	fieldName, ok := r.URL.Query()["fieldName"]
	if !ok {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	filterValue, ok := r.URL.Query()["filterValue"]

	data, err := storage.GetMyOpenData(app.repo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	filteredData, err := service.FilterData(data, service.DataFilter{
		FieldName: fieldName[0],
		Value:     filterValue[0],
	})

	res, _ := json.Marshal(filteredData)

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(res)

}

func (app *application) GetFilteredOpenDataInCSV(w http.ResponseWriter, r *http.Request) {
	fieldName, ok := r.URL.Query()["fieldName"]
	if !ok {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	filterValue, ok := r.URL.Query()["filterValue"]

	data, err := storage.GetMyOpenData(app.repo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	filteredData, err := service.FilterData(data, service.DataFilter{
		FieldName: fieldName[0],
		Value:     filterValue[0],
	})

	csvFile, err := service.ExportDataAsCSV(filteredData)
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer csvFile.Close()

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Content-Transfer-Encoding", "binary")
	(w).Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=cars.csv")
	http.ServeFile(w, r, "cars.csv")
}

func (app *application) GetFilteredOpenDataInJSON(w http.ResponseWriter, r *http.Request) {
	fieldName, ok := r.URL.Query()["fieldName"]
	if !ok {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	filterValue, ok := r.URL.Query()["filterValue"]

	data, err := storage.GetMyOpenData(app.repo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	filteredData, err := service.FilterData(data, service.DataFilter{
		FieldName: fieldName[0],
		Value:     filterValue[0],
	})

	byteData, err := json.Marshal(filteredData)
	if err != nil {
		app.serverError(w, err)
		return
	}

	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Disposition", "attachment;filename=cars.json")
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteData)
}
