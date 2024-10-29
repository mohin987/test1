// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetCwfNetworkIDGatewaysGatewayIDMagmadReader is a Reader for the GetCwfNetworkIDGatewaysGatewayIDMagmad structure.
type GetCwfNetworkIDGatewaysGatewayIDMagmadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDGatewaysGatewayIDMagmadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDGatewaysGatewayIDMagmadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDGatewaysGatewayIDMagmadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDMagmadOK creates a GetCwfNetworkIDGatewaysGatewayIDMagmadOK with default headers values
func NewGetCwfNetworkIDGatewaysGatewayIDMagmadOK() *GetCwfNetworkIDGatewaysGatewayIDMagmadOK {
	return &GetCwfNetworkIDGatewaysGatewayIDMagmadOK{}
}

/*GetCwfNetworkIDGatewaysGatewayIDMagmadOK handles this case with default header values.

Magmad agent configuration
*/
type GetCwfNetworkIDGatewaysGatewayIDMagmadOK struct {
	Payload *models.MagmadGatewayConfigs
}

func (o *GetCwfNetworkIDGatewaysGatewayIDMagmadOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways/{gateway_id}/magmad][%d] getCwfNetworkIdGatewaysGatewayIdMagmadOK  %+v", 200, o.Payload)
}

func (o *GetCwfNetworkIDGatewaysGatewayIDMagmadOK) GetPayload() *models.MagmadGatewayConfigs {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysGatewayIDMagmadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MagmadGatewayConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDGatewaysGatewayIDMagmadDefault creates a GetCwfNetworkIDGatewaysGatewayIDMagmadDefault with default headers values
func NewGetCwfNetworkIDGatewaysGatewayIDMagmadDefault(code int) *GetCwfNetworkIDGatewaysGatewayIDMagmadDefault {
	return &GetCwfNetworkIDGatewaysGatewayIDMagmadDefault{
		_statusCode: code,
	}
}

/*GetCwfNetworkIDGatewaysGatewayIDMagmadDefault handles this case with default header values.

Unexpected Error
*/
type GetCwfNetworkIDGatewaysGatewayIDMagmadDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID gateways gateway ID magmad default response
func (o *GetCwfNetworkIDGatewaysGatewayIDMagmadDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDGatewaysGatewayIDMagmadDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways/{gateway_id}/magmad][%d] GetCwfNetworkIDGatewaysGatewayIDMagmad default  %+v", o._statusCode, o.Payload)
}

func (o *GetCwfNetworkIDGatewaysGatewayIDMagmadDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysGatewayIDMagmadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
