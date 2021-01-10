// Package classification of License API
//
// Documentation for License API
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
package licensehandler

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

// Data structure representing a single license
// swagger:response licenseResponse
type licenseResponseWrapper struct {
	// A License
	// in: body
	Body []data.License
}

// Data structure representing a list of licenses
// swagger:response licensesResponse
type licensesResponseWrapper struct {
	// A List of License
	// in: body
	Body []data.Licenses
}

// swagger:parameters getLicense listLicenses generateLicense
type objectIDParamsWrapper struct {
	// The id of the object for which the operation relates
	// in: path
	// required: true
	ID uuid.UUID `json:"id"`
}
