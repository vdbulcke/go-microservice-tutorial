// Package classification of Tenant API
//
// Documentation for Tenant API
//
//	Schemes: http
//	BasePath: /api/beta2/
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package tenantshandler

import (
	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/handlers/api"

	"github.com/google/uuid"
)

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body api.GenericAPIError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of tenants
// swagger:response tenantsResponse
type tenantsResponseWrapper struct {
	// All current tenants
	// in: body
	Body []data.Tenant
}

// Data structure representing a single tenant
// swagger:response tenantResponse
type tenantResponseWrapper struct {
	// Newly created tenant
	// in: body
	Body data.Tenant
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateTenant
type tenantUpdateParamsWrapper struct {
	// Tenant data structure to Update or Create
	// in: body
	// required: true
	Body data.TenantUpdate
}

// swagger:parameters createTenant
type tenantCreateParamsWrapper struct {
	// Tenant data structure to Update or Create
	// in: body
	// required: true
	Body data.TenantCreate
}

// swagger:parameters getTenant deleteTenant
type tenantIDParamsWrapper struct {
	// The id of the tenant for which the operation relates
	// in: path
	// required: true
	ID uuid.UUID `json:"id"`
}
