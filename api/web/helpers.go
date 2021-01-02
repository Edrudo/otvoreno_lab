package main

import (
	"encoding/json"
	"fmt"
	"github.com/Edrudo/otvoreno_lab/service"
	"net/http"
	"runtime/debug"
)

/*
The serverError helper writes an error message and stack trace to the errorLog,
then sends a generic 500 Internal Server Error response to the employee.
*/
func (app *application) serverError(w http.ResponseWriter, err error, message string, status int) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	response := service.ApiResponse{
		Status:  http.StatusText(status),
		Message: message,
	}

	res, _ := json.Marshal(response)
	(w).Header().Set("Content-Type", "application/json")
	(w).WriteHeader(status)
	_, _ = w.Write(res)

}

/*
The clientError helper sends a specific status code and corresponding description
to the employee
*/
func (app *application) clientError(w http.ResponseWriter, status int, message string) {

	response := service.ApiResponse{
		Status:  http.StatusText(status),
		Message: message,
	}
	res, _ := json.Marshal(response)
	(w).Header().Set("Content-Type", "application/json")
	(w).WriteHeader(status)
	_, _ = w.Write(res)
}
