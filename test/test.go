package main

import (
	"errors"
	"fmt"
	"net/http"
)

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		fmt.Errorf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

func viewRecord(w http.ResponseWriter, r *http.Request) *appError {

	err := errors.New("Custom error")
	return &appError{Error: err, Message: "some message", Code: 501}
}

func main() {
	http.Handle("/view", appHandler(viewRecord))

	http.ListenAndServe(":8080", nil)
}
