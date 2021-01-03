package tenantshandler

import (
	"net/http"

	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/handlers/api"
)

// swagger:route POST /tenants Tenants createTenant
// Create a new tenant
//
// responses:
//	200: tenantResponse
//  409: tenantResponse
//  400: errorResponse
//  500: errorResponse

// Create handles POST requests to add new tenants
func (p *Tenants) Create(rw http.ResponseWriter, r *http.Request) *api.APIError {
	// fetch the tenant from the context
	// TODO: error handling
	tenantCreate := r.Context().Value(KeyTenant{}).(data.TenantCreate)

	// create Tenant Object from tenantCreate
	tenant := &data.Tenant{
		Name:        tenantCreate.Name,
		Description: tenantCreate.Description,
	}

	p.l.Printf("[DEBUG] Inserting tenant: %#v\n", tenant)
	createdTenant, err := data.AddTenant(*tenant, p.db)
	statusCode := http.StatusCreated
	if err != nil {
		switch err.(type) {
		case *data.TenantAlreadyExist:

			// overwrite the status code
			// to 409 conflict
			statusCode = http.StatusConflict

		// return &api.APIError{Err: err,
		// 	Code:    http.StatusConflict,
		// 	Type:    &api.TenantAlreadyExistError{},
		// 	Message: err.Error(),
		// }
		default:
			return &api.APIError{Err: err,
				Code:    http.StatusInternalServerError,
				Type:    &api.InternalServerError{},
				Message: "error creating tenant",
			}
		}

	}

	// set status code (need to be before the write)
	rw.WriteHeader(statusCode)
	err = data.ToJSON(createdTenant, rw)
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
