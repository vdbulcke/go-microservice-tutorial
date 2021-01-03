package api

import (
	"go-microservice-tutorial/organization-api/data"
	"net/http"
)

// INFO: https://blog.golang.org/error-handling-and-go

// APIHandler wrapper for http.handler doing error handling
type APIHandler struct {
	Handler func(http.ResponseWriter, *http.Request) *APIError
}

func (fn APIHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// call handler
	err := fn.Handler(rw, r)
	if err != nil {

		// return http status code from error
		rw.WriteHeader(err.Code)
		switch err.Type.(type) {
		case *InternalServerError:
			// write JSON error message
			data.ToJSON(&GenericAPIError{Message: err.Message, Error: "internal_server_error"}, rw)
			return

		case *TenantAlreadyExistError:
			// write JSON error message
			data.ToJSON(&GenericAPIError{Message: err.Message, Error: "object_already_exists"}, rw)
			return

		case *ObjectNotFoundError:
			// write JSON error message
			data.ToJSON(&GenericAPIError{Message: err.Message, Error: "object_not_found"}, rw)
			return
		case *ValidationError:
			// write JSON error message
			data.ToJSON(&GenericAPIError{Message: err.Message, Error: "validation_error"}, rw)
			return
		default:
			// write JSON error message
			data.ToJSON(&GenericAPIError{Message: "Something went wrong", Error: "internal_server_error"}, rw)
			return
		}

	}

}
