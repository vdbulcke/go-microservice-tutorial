package tenantshandler

import (
	"fmt"
	"log"
	"net/http"

	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/data/database"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// KeyTenant is a key used for the Tenant object in the context
type KeyTenant struct{}

// Tenants handler for getting and updating tenants
type Tenants struct {
	l  *log.Logger
	v  *data.Validation
	db *database.DB
}

// NewTenants returns a new tenant handler with the given logger
func NewTenants(l *log.Logger, v *data.Validation, db *database.DB) *Tenants {
	return &Tenants{l, v, db}
}

// ErrInvalidTenantPath is an error message when the terant path is not valid
var ErrInvalidTenantPath = fmt.Errorf("Invalid Path, path should be /tenants/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getTenantID returns the tenant ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getTenantID(r *http.Request) (uuid.UUID, error) {
	// parse the tenant id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		// should never happen
		return id, err
	}

	return id, nil
}
