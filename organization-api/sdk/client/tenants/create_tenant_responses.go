// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"go-microservice-tutorial/organization-api/sdk/models"
)

// CreateTenantReader is a Reader for the CreateTenant structure.
type CreateTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateTenantUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewCreateTenantNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTenantOK creates a CreateTenantOK with default headers values
func NewCreateTenantOK() *CreateTenantOK {
	return &CreateTenantOK{}
}

/*CreateTenantOK handles this case with default header values.

Data structure representing a single tenant
*/
type CreateTenantOK struct {
	Payload *models.Tenant
}

func (o *CreateTenantOK) Error() string {
	return fmt.Sprintf("[POST /tenants][%d] createTenantOK  %+v", 200, o.Payload)
}

func (o *CreateTenantOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *CreateTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTenantUnprocessableEntity creates a CreateTenantUnprocessableEntity with default headers values
func NewCreateTenantUnprocessableEntity() *CreateTenantUnprocessableEntity {
	return &CreateTenantUnprocessableEntity{}
}

/*CreateTenantUnprocessableEntity handles this case with default header values.

Validation errors defined as an array of strings
*/
type CreateTenantUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateTenantUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /tenants][%d] createTenantUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateTenantUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateTenantUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTenantNotImplemented creates a CreateTenantNotImplemented with default headers values
func NewCreateTenantNotImplemented() *CreateTenantNotImplemented {
	return &CreateTenantNotImplemented{}
}

/*CreateTenantNotImplemented handles this case with default header values.

Generic error message returned as a string
*/
type CreateTenantNotImplemented struct {
	Payload *models.GenericError
}

func (o *CreateTenantNotImplemented) Error() string {
	return fmt.Sprintf("[POST /tenants][%d] createTenantNotImplemented  %+v", 501, o.Payload)
}

func (o *CreateTenantNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateTenantNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}