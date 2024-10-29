// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PostNetworksNetworkIDPoliciesBaseNamesReader is a Reader for the PostNetworksNetworkIDPoliciesBaseNames structure.
type PostNetworksNetworkIDPoliciesBaseNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksNetworkIDPoliciesBaseNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostNetworksNetworkIDPoliciesBaseNamesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksNetworkIDPoliciesBaseNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksNetworkIDPoliciesBaseNamesCreated creates a PostNetworksNetworkIDPoliciesBaseNamesCreated with default headers values
func NewPostNetworksNetworkIDPoliciesBaseNamesCreated() *PostNetworksNetworkIDPoliciesBaseNamesCreated {
	return &PostNetworksNetworkIDPoliciesBaseNamesCreated{}
}

/*PostNetworksNetworkIDPoliciesBaseNamesCreated handles this case with default header values.

Charging Rule Base Name
*/
type PostNetworksNetworkIDPoliciesBaseNamesCreated struct {
	Payload models.BaseName
}

func (o *PostNetworksNetworkIDPoliciesBaseNamesCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/policies/base_names][%d] postNetworksNetworkIdPoliciesBaseNamesCreated  %+v", 201, o.Payload)
}

func (o *PostNetworksNetworkIDPoliciesBaseNamesCreated) GetPayload() models.BaseName {
	return o.Payload
}

func (o *PostNetworksNetworkIDPoliciesBaseNamesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworksNetworkIDPoliciesBaseNamesDefault creates a PostNetworksNetworkIDPoliciesBaseNamesDefault with default headers values
func NewPostNetworksNetworkIDPoliciesBaseNamesDefault(code int) *PostNetworksNetworkIDPoliciesBaseNamesDefault {
	return &PostNetworksNetworkIDPoliciesBaseNamesDefault{
		_statusCode: code,
	}
}

/*PostNetworksNetworkIDPoliciesBaseNamesDefault handles this case with default header values.

Unexpected Error
*/
type PostNetworksNetworkIDPoliciesBaseNamesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks network ID policies base names default response
func (o *PostNetworksNetworkIDPoliciesBaseNamesDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksNetworkIDPoliciesBaseNamesDefault) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/policies/base_names][%d] PostNetworksNetworkIDPoliciesBaseNames default  %+v", o._statusCode, o.Payload)
}

func (o *PostNetworksNetworkIDPoliciesBaseNamesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksNetworkIDPoliciesBaseNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
