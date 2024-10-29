// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDPrometheusQueryReader is a Reader for the GetNetworksNetworkIDPrometheusQuery structure.
type GetNetworksNetworkIDPrometheusQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDPrometheusQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDPrometheusQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDPrometheusQueryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDPrometheusQueryOK creates a GetNetworksNetworkIDPrometheusQueryOK with default headers values
func NewGetNetworksNetworkIDPrometheusQueryOK() *GetNetworksNetworkIDPrometheusQueryOK {
	return &GetNetworksNetworkIDPrometheusQueryOK{}
}

/*GetNetworksNetworkIDPrometheusQueryOK handles this case with default header values.

List of PromQL metrics results
*/
type GetNetworksNetworkIDPrometheusQueryOK struct {
	Payload *models.PromqlReturnObject
}

func (o *GetNetworksNetworkIDPrometheusQueryOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/prometheus/query][%d] getNetworksNetworkIdPrometheusQueryOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDPrometheusQueryOK) GetPayload() *models.PromqlReturnObject {
	return o.Payload
}

func (o *GetNetworksNetworkIDPrometheusQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromqlReturnObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDPrometheusQueryDefault creates a GetNetworksNetworkIDPrometheusQueryDefault with default headers values
func NewGetNetworksNetworkIDPrometheusQueryDefault(code int) *GetNetworksNetworkIDPrometheusQueryDefault {
	return &GetNetworksNetworkIDPrometheusQueryDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDPrometheusQueryDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDPrometheusQueryDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID prometheus query default response
func (o *GetNetworksNetworkIDPrometheusQueryDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDPrometheusQueryDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/prometheus/query][%d] GetNetworksNetworkIDPrometheusQuery default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDPrometheusQueryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDPrometheusQueryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
