package licensehandler

import (
	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/handlers/api"
	"net/http"
)

// swagger:route GET /license/get_license_by_id/{id} License getLicense
// Return a license from the database
//	200: licenseResponse
//	404: errorResponse
//  401: errorResponse
//  500: errorResponse

// GetLicenseByID handles Get requests and return a license
func (handler *LicenseHandler) GetLicenseByID(rw http.ResponseWriter, r *http.Request) *api.APIError {

	// get UUID from request
	id, uuidErr := getID(r)
	if uuidErr != nil {
		return uuidErr
	}

	// get license from DB
	license, err := handler.db.GetLicenceByUUID(id)
	if err != nil {
		handler.logger.Error("searching license ", "id", id.String(), "error", err)
		return &api.APIError{
			Err:     err,
			Code:    http.StatusNotFound,
			Type:    &api.ObjectNotFoundError{},
			Message: "Could not find license with id " + id.String(),
		}
	}

	// Setting Status code before writing body
	rw.WriteHeader(http.StatusOK)

	err = data.ToJSON(license, rw)
	if err != nil {
		// we should never be here but log the error just incase
		handler.logger.Error("writing json", "error", err)
		return &api.APIError{Err: err,
			Code:    http.StatusInternalServerError,
			Type:    &api.InternalServerError{},
			Message: "error formating json",
		}
	}

	// return nil if everything is ok
	return nil
}

// swagger:route GET /license/get_licenses_by_tenant_id/{id} License listLicenses
// Return a list of licenses from the database
// responses:
//	200: licensesResponse
//	404: errorResponse
//  401: errorResponse
//  500: errorResponse

// GetLicensesByTenantID handles GET request, returns a list of license of a tenant
func (handler *LicenseHandler) GetLicensesByTenantID(rw http.ResponseWriter, r *http.Request) *api.APIError {

	// get UUID from request
	id, uuidErr := getID(r)
	if uuidErr != nil {
		return uuidErr
	}

	// get licenses from DB
	licenses, err := handler.db.GetLicensesByTenantID(id)
	if err != nil {
		handler.logger.Error("searching licenses for tenant ", "id", id.String(), "error", err)
		return &api.APIError{
			Err:     err,
			Code:    http.StatusNotFound,
			Type:    &api.ObjectNotFoundError{},
			Message: "Could not find licenses for tenant with id " + id.String(),
		}
	}

	// Setting Status code before writing body
	rw.WriteHeader(http.StatusOK)

	err = data.ToJSON(licenses, rw)
	if err != nil {
		// we should never be here but log the error just incase
		handler.logger.Error("writing json", "error", err)
		return &api.APIError{Err: err,
			Code:    http.StatusInternalServerError,
			Type:    &api.InternalServerError{},
			Message: "error formating json",
		}
	}

	// return nil if everything is ok
	return nil
}

// swagger:route GET /license/generate_license_for_tenant_id/{id} License generateLicense
// Return a license newly created license
// responses:
//	200: licenseResponse
//	404: errorResponse
//  401: errorResponse
//  500: errorResponse

// GenerateLicenceForTenant handles GET requests and return a new license
func (handler *LicenseHandler) GenerateLicenceForTenant(rw http.ResponseWriter, r *http.Request) *api.APIError {

	// get UUID from request
	id, uuidErr := getID(r)
	if uuidErr != nil {
		return uuidErr
	}

	license, err := handler.db.CreateLicenseForTenant(id)
	if err != nil {
		return &api.APIError{Err: err,
			Code:    http.StatusInternalServerError,
			Type:    &api.InternalServerError{},
			Message: "error creating license for tenant " + id.String(),
		}
	}

	// Setting Status code before writing body
	rw.WriteHeader(http.StatusOK)

	err = data.ToJSON(license, rw)
	if err != nil {
		// we should never be here but log the error just incase
		handler.logger.Error("writing json", "error", err)
		return &api.APIError{Err: err,
			Code:    http.StatusInternalServerError,
			Type:    &api.InternalServerError{},
			Message: "error formating json",
		}
	}

	// return nil if everything is ok
	return nil

}
