// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutFegLTENetworkIDReader is a Reader for the PutFegLTENetworkID structure.
type PutFegLTENetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFegLTENetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutFegLTENetworkIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutFegLTENetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutFegLTENetworkIDNoContent creates a PutFegLTENetworkIDNoContent with default headers values
func NewPutFegLTENetworkIDNoContent() *PutFegLTENetworkIDNoContent {
	return &PutFegLTENetworkIDNoContent{}
}

/*PutFegLTENetworkIDNoContent handles this case with default header values.

Success
*/
type PutFegLTENetworkIDNoContent struct {
}

func (o *PutFegLTENetworkIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /feg_lte/{network_id}][%d] putFegLteNetworkIdNoContent ", 204)
}

func (o *PutFegLTENetworkIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFegLTENetworkIDDefault creates a PutFegLTENetworkIDDefault with default headers values
func NewPutFegLTENetworkIDDefault(code int) *PutFegLTENetworkIDDefault {
	return &PutFegLTENetworkIDDefault{
		_statusCode: code,
	}
}

/*PutFegLTENetworkIDDefault handles this case with default header values.

Unexpected Error
*/
type PutFegLTENetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put feg LTE network ID default response
func (o *PutFegLTENetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *PutFegLTENetworkIDDefault) Error() string {
	return fmt.Sprintf("[PUT /feg_lte/{network_id}][%d] PutFegLTENetworkID default  %+v", o._statusCode, o.Payload)
}

func (o *PutFegLTENetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutFegLTENetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
