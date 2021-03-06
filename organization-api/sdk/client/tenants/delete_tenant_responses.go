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

// DeleteTenantReader is a Reader for the DeleteTenant structure.
type DeleteTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeleteTenantCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTenantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDeleteTenantNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTenantCreated creates a DeleteTenantCreated with default headers values
func NewDeleteTenantCreated() *DeleteTenantCreated {
	return &DeleteTenantCreated{}
}

/*DeleteTenantCreated handles this case with default header values.

No content is returned by this API endpoint
*/
type DeleteTenantCreated struct {
}

func (o *DeleteTenantCreated) Error() string {
	return fmt.Sprintf("[DELETE /tenants/{id}][%d] deleteTenantCreated ", 201)
}

func (o *DeleteTenantCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTenantNotFound creates a DeleteTenantNotFound with default headers values
func NewDeleteTenantNotFound() *DeleteTenantNotFound {
	return &DeleteTenantNotFound{}
}

/*DeleteTenantNotFound handles this case with default header values.

Generic error message returned as a string
*/
type DeleteTenantNotFound struct {
	Payload *models.GenericError
}

func (o *DeleteTenantNotFound) Error() string {
	return fmt.Sprintf("[DELETE /tenants/{id}][%d] deleteTenantNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTenantNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteTenantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTenantNotImplemented creates a DeleteTenantNotImplemented with default headers values
func NewDeleteTenantNotImplemented() *DeleteTenantNotImplemented {
	return &DeleteTenantNotImplemented{}
}

/*DeleteTenantNotImplemented handles this case with default header values.

Generic error message returned as a string
*/
type DeleteTenantNotImplemented struct {
	Payload *models.GenericError
}

func (o *DeleteTenantNotImplemented) Error() string {
	return fmt.Sprintf("[DELETE /tenants/{id}][%d] deleteTenantNotImplemented  %+v", 501, o.Payload)
}

func (o *DeleteTenantNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteTenantNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
