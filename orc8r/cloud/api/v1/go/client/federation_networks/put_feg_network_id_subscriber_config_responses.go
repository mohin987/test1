// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutFegNetworkIDSubscriberConfigReader is a Reader for the PutFegNetworkIDSubscriberConfig structure.
type PutFegNetworkIDSubscriberConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFegNetworkIDSubscriberConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutFegNetworkIDSubscriberConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutFegNetworkIDSubscriberConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutFegNetworkIDSubscriberConfigNoContent creates a PutFegNetworkIDSubscriberConfigNoContent with default headers values
func NewPutFegNetworkIDSubscriberConfigNoContent() *PutFegNetworkIDSubscriberConfigNoContent {
	return &PutFegNetworkIDSubscriberConfigNoContent{}
}

/*PutFegNetworkIDSubscriberConfigNoContent handles this case with default header values.

Success
*/
type PutFegNetworkIDSubscriberConfigNoContent struct {
}

func (o *PutFegNetworkIDSubscriberConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /feg/{network_id}/subscriber_config][%d] putFegNetworkIdSubscriberConfigNoContent ", 204)
}

func (o *PutFegNetworkIDSubscriberConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFegNetworkIDSubscriberConfigDefault creates a PutFegNetworkIDSubscriberConfigDefault with default headers values
func NewPutFegNetworkIDSubscriberConfigDefault(code int) *PutFegNetworkIDSubscriberConfigDefault {
	return &PutFegNetworkIDSubscriberConfigDefault{
		_statusCode: code,
	}
}

/*PutFegNetworkIDSubscriberConfigDefault handles this case with default header values.

Unexpected Error
*/
type PutFegNetworkIDSubscriberConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put feg network ID subscriber config default response
func (o *PutFegNetworkIDSubscriberConfigDefault) Code() int {
	return o._statusCode
}

func (o *PutFegNetworkIDSubscriberConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /feg/{network_id}/subscriber_config][%d] PutFegNetworkIDSubscriberConfig default  %+v", o._statusCode, o.Payload)
}

func (o *PutFegNetworkIDSubscriberConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutFegNetworkIDSubscriberConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
