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

// GetNetworksNetworkIDPoliciesRulesViewFullReader is a Reader for the GetNetworksNetworkIDPoliciesRulesViewFull structure.
type GetNetworksNetworkIDPoliciesRulesViewFullReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDPoliciesRulesViewFullReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDPoliciesRulesViewFullOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDPoliciesRulesViewFullDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDPoliciesRulesViewFullOK creates a GetNetworksNetworkIDPoliciesRulesViewFullOK with default headers values
func NewGetNetworksNetworkIDPoliciesRulesViewFullOK() *GetNetworksNetworkIDPoliciesRulesViewFullOK {
	return &GetNetworksNetworkIDPoliciesRulesViewFullOK{}
}

/*GetNetworksNetworkIDPoliciesRulesViewFullOK handles this case with default header values.

Map of all policy rules for the network by rule ID
*/
type GetNetworksNetworkIDPoliciesRulesViewFullOK struct {
	Payload map[string]models.PolicyRule
}

func (o *GetNetworksNetworkIDPoliciesRulesViewFullOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/policies/rules?view=full][%d] getNetworksNetworkIdPoliciesRulesViewFullOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDPoliciesRulesViewFullOK) GetPayload() map[string]models.PolicyRule {
	return o.Payload
}

func (o *GetNetworksNetworkIDPoliciesRulesViewFullOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDPoliciesRulesViewFullDefault creates a GetNetworksNetworkIDPoliciesRulesViewFullDefault with default headers values
func NewGetNetworksNetworkIDPoliciesRulesViewFullDefault(code int) *GetNetworksNetworkIDPoliciesRulesViewFullDefault {
	return &GetNetworksNetworkIDPoliciesRulesViewFullDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDPoliciesRulesViewFullDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDPoliciesRulesViewFullDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID policies rules view full default response
func (o *GetNetworksNetworkIDPoliciesRulesViewFullDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDPoliciesRulesViewFullDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/policies/rules?view=full][%d] GetNetworksNetworkIDPoliciesRulesViewFull default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDPoliciesRulesViewFullDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDPoliciesRulesViewFullDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
