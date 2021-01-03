package tenantshandler

import (
	"net/http"

	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/handlers/api"
)

// swagger:route PUT /tenants Tenants updateTenant
// Update a tenants details
//
// responses:
//	204: noContentResponse
//  404: errorResponse
//  400: errorResponse
//  500: errorResponse

// Update handles PUT requests to update tenants
func (p *Tenants) Update(rw http.ResponseWriter, r *http.Request) *api.APIError {

	// fetch the tenant from the context
	tenantUpdate := r.Context().Value(KeyTenant{}).(data.TenantUpdate)
	p.l.Println("[DEBUG] updating record id", tenantUpdate.ID)

	// create tenant  from tenantUpdate
	// create Tenant Object from tenantCreate
	tenant := &data.Tenant{
		ID:          tenantUpdate.ID,
		Name:        tenantUpdate.Name,
		Description: tenantUpdate.Description,
	}

	err := data.UpdateTenant(*tenant, p.db)
	if err == data.ErrTenantNotFound {
		p.l.Println("[ERROR] tenant not found", err)

		return &api.APIError{Err: err,
			Code:    http.StatusNotFound,
			Type:    &api.ObjectNotFoundError{},
			Message: "Could not find tenants with id " + tenant.ID.String(),
		}
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
	return nil
}
