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

// ListTenantsReader is a Reader for the ListTenants structure.
type ListTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTenantsOK creates a ListTenantsOK with default headers values
func NewListTenantsOK() *ListTenantsOK {
	return &ListTenantsOK{}
}

/*ListTenantsOK handles this case with default header values.

A list of tenants
*/
type ListTenantsOK struct {
	Payload []*models.Tenant
}

func (o *ListTenantsOK) Error() string {
	return fmt.Sprintf("[GET /tenants][%d] listTenantsOK  %+v", 200, o.Payload)
}

func (o *ListTenantsOK) GetPayload() []*models.Tenant {
	return o.Payload
}

func (o *ListTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
