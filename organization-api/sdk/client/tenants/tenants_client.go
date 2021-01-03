// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tenants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTenant(params *CreateTenantParams) (*CreateTenantOK, error)

	DeleteTenant(params *DeleteTenantParams) (*DeleteTenantCreated, error)

	GetTenant(params *GetTenantParams) (*GetTenantOK, error)

	ListTenants(params *ListTenantsParams) (*ListTenantsOK, error)

	UpdateTenant(params *UpdateTenantParams) (*UpdateTenantCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTenant Create a new tenant
*/
func (a *Client) CreateTenant(params *CreateTenantParams) (*CreateTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTenant",
		Method:             "POST",
		PathPattern:        "/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTenantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTenant Delete a tenant details
*/
func (a *Client) DeleteTenant(params *DeleteTenantParams) (*DeleteTenantCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTenant",
		Method:             "DELETE",
		PathPattern:        "/tenants/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTenantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTenantCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTenant Return a list of tenants from the database
*/
func (a *Client) GetTenant(params *GetTenantParams) (*GetTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTenant",
		Method:             "GET",
		PathPattern:        "/tenants/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTenantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListTenants Return a list of tenants from the database
*/
func (a *Client) ListTenants(params *ListTenantsParams) (*ListTenantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTenantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listTenants",
		Method:             "GET",
		PathPattern:        "/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListTenantsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListTenantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listTenants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateTenant Update a tenants details
*/
func (a *Client) UpdateTenant(params *UpdateTenantParams) (*UpdateTenantCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateTenant",
		Method:             "PUT",
		PathPattern:        "/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTenantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTenantCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}