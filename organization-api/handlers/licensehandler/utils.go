package licensehandler

import (
	"net/http"

	"go-microservice-tutorial/organization-api/handlers/api"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// getID returns the ID from the URL
func getID(r *http.Request) (uuid.UUID, *api.APIError) {
	// parse the id from the url
	vars := mux.Vars(r)

	// convert the id into an uuid and return
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		return id, &api.APIError{Err: err,
			Code:    http.StatusBadRequest,
			Type:    &api.ValidationError{},
			Message: "error getting tenant",
		}
	}

	return id, nil
}
