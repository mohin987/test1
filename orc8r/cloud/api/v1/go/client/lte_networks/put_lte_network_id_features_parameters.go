// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutLTENetworkIDFeaturesParams creates a new PutLTENetworkIDFeaturesParams object
// with the default values initialized.
func NewPutLTENetworkIDFeaturesParams() *PutLTENetworkIDFeaturesParams {
	var ()
	return &PutLTENetworkIDFeaturesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDFeaturesParamsWithTimeout creates a new PutLTENetworkIDFeaturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDFeaturesParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDFeaturesParams {
	var ()
	return &PutLTENetworkIDFeaturesParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDFeaturesParamsWithContext creates a new PutLTENetworkIDFeaturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDFeaturesParamsWithContext(ctx context.Context) *PutLTENetworkIDFeaturesParams {
	var ()
	return &PutLTENetworkIDFeaturesParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDFeaturesParamsWithHTTPClient creates a new PutLTENetworkIDFeaturesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDFeaturesParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDFeaturesParams {
	var ()
	return &PutLTENetworkIDFeaturesParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDFeaturesParams contains all the parameters to send to the API endpoint
for the put LTE network ID features operation typically these are written to a http.Request
*/
type PutLTENetworkIDFeaturesParams struct {

	/*Config
	  New feature flags for the network

	*/
	Config *models.NetworkFeatures
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) WithContext(ctx context.Context) *PutLTENetworkIDFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDFeaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) WithConfig(config *models.NetworkFeatures) *PutLTENetworkIDFeaturesParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) SetConfig(config *models.NetworkFeatures) {
	o.Config = config
}

// WithNetworkID adds the networkID to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) WithNetworkID(networkID string) *PutLTENetworkIDFeaturesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID features params
func (o *PutLTENetworkIDFeaturesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
