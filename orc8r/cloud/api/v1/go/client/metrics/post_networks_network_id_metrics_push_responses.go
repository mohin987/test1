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

// PostNetworksNetworkIDMetricsPushReader is a Reader for the PostNetworksNetworkIDMetricsPush structure.
type PostNetworksNetworkIDMetricsPushReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksNetworkIDMetricsPushReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNetworksNetworkIDMetricsPushOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksNetworkIDMetricsPushDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksNetworkIDMetricsPushOK creates a PostNetworksNetworkIDMetricsPushOK with default headers values
func NewPostNetworksNetworkIDMetricsPushOK() *PostNetworksNetworkIDMetricsPushOK {
	return &PostNetworksNetworkIDMetricsPushOK{}
}

/*PostNetworksNetworkIDMetricsPushOK handles this case with default header values.

Submitted
*/
type PostNetworksNetworkIDMetricsPushOK struct {
}

func (o *PostNetworksNetworkIDMetricsPushOK) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/metrics/push][%d] postNetworksNetworkIdMetricsPushOK ", 200)
}

func (o *PostNetworksNetworkIDMetricsPushOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostNetworksNetworkIDMetricsPushDefault creates a PostNetworksNetworkIDMetricsPushDefault with default headers values
func NewPostNetworksNetworkIDMetricsPushDefault(code int) *PostNetworksNetworkIDMetricsPushDefault {
	return &PostNetworksNetworkIDMetricsPushDefault{
		_statusCode: code,
	}
}

/*PostNetworksNetworkIDMetricsPushDefault handles this case with default header values.

Unexpected Error
*/
type PostNetworksNetworkIDMetricsPushDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks network ID metrics push default response
func (o *PostNetworksNetworkIDMetricsPushDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksNetworkIDMetricsPushDefault) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/metrics/push][%d] PostNetworksNetworkIDMetricsPush default  %+v", o._statusCode, o.Payload)
}

func (o *PostNetworksNetworkIDMetricsPushDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksNetworkIDMetricsPushDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
