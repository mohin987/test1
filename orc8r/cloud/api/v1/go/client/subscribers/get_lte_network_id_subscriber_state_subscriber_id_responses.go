// Code generated by go-swagger; DO NOT EDIT.

package subscribers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDSubscriberStateSubscriberIDReader is a Reader for the GetLTENetworkIDSubscriberStateSubscriberID structure.
type GetLTENetworkIDSubscriberStateSubscriberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDSubscriberStateSubscriberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDSubscriberStateSubscriberIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDSubscriberStateSubscriberIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDSubscriberStateSubscriberIDOK creates a GetLTENetworkIDSubscriberStateSubscriberIDOK with default headers values
func NewGetLTENetworkIDSubscriberStateSubscriberIDOK() *GetLTENetworkIDSubscriberStateSubscriberIDOK {
	return &GetLTENetworkIDSubscriberStateSubscriberIDOK{}
}

/*GetLTENetworkIDSubscriberStateSubscriberIDOK handles this case with default header values.

Subscriber state
*/
type GetLTENetworkIDSubscriberStateSubscriberIDOK struct {
	Payload *models.SubscriberState
}

func (o *GetLTENetworkIDSubscriberStateSubscriberIDOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/subscriber_state/{subscriber_id}][%d] getLteNetworkIdSubscriberStateSubscriberIdOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDSubscriberStateSubscriberIDOK) GetPayload() *models.SubscriberState {
	return o.Payload
}

func (o *GetLTENetworkIDSubscriberStateSubscriberIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriberState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDSubscriberStateSubscriberIDDefault creates a GetLTENetworkIDSubscriberStateSubscriberIDDefault with default headers values
func NewGetLTENetworkIDSubscriberStateSubscriberIDDefault(code int) *GetLTENetworkIDSubscriberStateSubscriberIDDefault {
	return &GetLTENetworkIDSubscriberStateSubscriberIDDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDSubscriberStateSubscriberIDDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDSubscriberStateSubscriberIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID subscriber state subscriber ID default response
func (o *GetLTENetworkIDSubscriberStateSubscriberIDDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDSubscriberStateSubscriberIDDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/subscriber_state/{subscriber_id}][%d] GetLTENetworkIDSubscriberStateSubscriberID default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDSubscriberStateSubscriberIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDSubscriberStateSubscriberIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
