package tenantshandler

import (
	"context"
	"net/http"

	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/handlers/api"
)

// MiddlewareValidateTenant validates the tenant in the request and calls next if ok
func (p *Tenants) MiddlewareValidateTenant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Tenant{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing tenant", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the tenant
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating tenant", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the tenant to the context
		ctx := context.WithValue(r.Context(), KeyTenant{}, *prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)

	})
}

// MiddlewareValidateTenantCreate validates the tenant in the request and calls next if ok
func (p *Tenants) MiddlewareValidateTenantCreate(next http.Handler) http.Handler {
	return http.Handler(api.APIHandler{Handler: func(rw http.ResponseWriter, r *http.Request) *api.APIError {
		tenantCreate := &data.TenantCreate{}

		err := data.FromJSON(tenantCreate, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing tenant", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return &api.APIError{Err: err,
				Code:    http.StatusBadRequest,
				Type:    &api.ValidationError{},
				Message: "error deserializing json",
			}
		}

		// validate the tenant
		errs := p.v.Validate(tenantCreate)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating tenant", errs)

			return &api.APIError{Err: err,
				Code:    http.StatusBadRequest,
				Type:    &api.ValidationError{},
				Message: "error validating tenant",
			}
		}

		// add the tenant to the context
		ctx := context.WithValue(r.Context(), KeyTenant{}, *tenantCreate)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
		return nil

	}})
}

// MiddlewareValidateTenantUpdate validates the tenant in the request and calls next if ok
func (p *Tenants) MiddlewareValidateTenantUpdate(next http.Handler) http.Handler {
	return http.Handler(api.APIHandler{Handler: func(rw http.ResponseWriter, r *http.Request) *api.APIError {
		tenantUpdate := &data.TenantUpdate{}

		err := data.FromJSON(tenantUpdate, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing tenant", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return &api.APIError{Err: err,
				Code:    http.StatusBadRequest,
				Type:    &api.ValidationError{},
				Message: "error deserializing json",
			}
		}

		// validate the tenant
		errs := p.v.Validate(tenantUpdate)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating tenant", errs)

			return &api.APIError{Err: err,
				Code:    http.StatusBadRequest,
				Type:    &api.ValidationError{},
				Message: "error validating tenant",
			}
		}

		// add the tenant to the context
		ctx := context.WithValue(r.Context(), KeyTenant{}, *tenantUpdate)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
		return nil

	}})
}
