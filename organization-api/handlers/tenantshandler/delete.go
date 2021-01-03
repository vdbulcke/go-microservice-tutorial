package tenantshandler

import (
	"net/http"

	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/handlers/api"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /tenants/{id} Tenants deleteTenant
// Delete a tenant details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  500: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Tenants) Delete(rw http.ResponseWriter, r *http.Request) *api.APIError {

	id, gettenanterr := getTenantID(r)
	if gettenanterr != nil {
		// parse the tenant id from the url
		vars := mux.Vars(r)

		return &api.APIError{Err: gettenanterr,
			Code:    http.StatusInternalServerError,
			Type:    &api.ValidationError{},
			Message: "Invalid UUID " + vars["id"],
		}
	}

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteTenant(id, p.db)
	if err == data.ErrTenantNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		return &api.APIError{Err: err,
			Code:    http.StatusNotFound,
			Type:    &api.ObjectNotFoundError{},
			Message: "Could not find tenants with id " + id.String(),
		}
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)
		return &api.APIError{Err: err,
			Code:    http.StatusInternalServerError,
			Type:    &api.InternalServerError{},
			Message: "Error deleting record with id " + id.String(),
		}
	}

	// Write Status code
	rw.WriteHeader(http.StatusNoContent)
	return nil
}
