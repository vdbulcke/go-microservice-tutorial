package tenantshandler

import (
	"net/http"

	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/handlers/api"

	"github.com/gorilla/mux"
)

// swagger:route GET /tenants Tenants listTenants
// Return a list of tenants from the database
// responses:
//	200: tenantsResponse
//  500: errorResponse

// ListAll handles GET requests and returns all current tenants
func (p *Tenants) ListAll(rw http.ResponseWriter, r *http.Request) *api.APIError {
	p.l.Println("[DEBUG] get all records")

	prods, getErr := data.GetTenants(p.db)
	if getErr != nil {
		p.l.Println("[ERROR] getting tenant", getErr)
		return &api.APIError{Err: getErr,
			Code:    http.StatusInternalServerError,
			Type:    &api.InternalServerError{},
			Message: "error  getting tenant",
		}
	}

	// Write Status code
	rw.WriteHeader(http.StatusOK)
	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing tenant", err)
		return &api.APIError{Err: err,
			Code:    http.StatusInternalServerError,
			Type:    &api.InternalServerError{},
			Message: "error formating json",
		}
	}

	return nil
}

// swagger:route GET /tenants/{id} Tenants getTenant
// Return a list of tenants from the database
// responses:
//	200: tenantResponse
//	404: errorResponse
//  401: errorResponse

// ListSingle handles GET requests
func (p *Tenants) ListSingle(rw http.ResponseWriter, r *http.Request) *api.APIError {

	id, gettenanterr := getTenantID(r)
	if gettenanterr != nil {
		// parse the tenant id from the url
		vars := mux.Vars(r)

		return &api.APIError{Err: gettenanterr,
			Code:    http.StatusBadRequest,
			Type:    &api.ValidationError{},
			Message: "Invalid UUID " + vars["id"],
		}
	}

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetTenantByID(id, p.db)

	switch err {
	case nil:

	case data.ErrTenantNotFound:
		p.l.Println("[ERROR] fetching tenant", err)

		return &api.APIError{Err: err,
			Code:    http.StatusNotFound,
			Type:    &api.ObjectNotFoundError{},
			Message: "Could not find tenants with id " + id.String(),
		}
	default:
		p.l.Println("[ERROR] fetching tenant", err)

		return &api.APIError{Err: err,
			Code:    http.StatusInternalServerError,
			Type:    &api.InternalServerError{},
			Message: "error getting tenant json",
		}
	}

	// Write Status code
	rw.WriteHeader(http.StatusOK)
	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing tenant", err)
		return &api.APIError{Err: err,
			Code:    http.StatusInternalServerError,
			Type:    &api.InternalServerError{},
			Message: "error formating json",
		}
	}

	return nil
}
