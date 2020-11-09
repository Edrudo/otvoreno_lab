package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

/*
The serverError helper writes an error message and stack trace to the errorLog,
then sends a generic 500 Internal Server Error response to the employee.
*/
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

/*
The clientError helper sends a specific status code and corresponding description
to the employee
*/
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
