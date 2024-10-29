// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTenantsTenantIDMetricsAPIV1QueryRangeParams creates a new GetTenantsTenantIDMetricsAPIV1QueryRangeParams object
// with the default values initialized.
func NewGetTenantsTenantIDMetricsAPIV1QueryRangeParams() *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	var ()
	return &GetTenantsTenantIDMetricsAPIV1QueryRangeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantsTenantIDMetricsAPIV1QueryRangeParamsWithTimeout creates a new GetTenantsTenantIDMetricsAPIV1QueryRangeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantsTenantIDMetricsAPIV1QueryRangeParamsWithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	var ()
	return &GetTenantsTenantIDMetricsAPIV1QueryRangeParams{

		timeout: timeout,
	}
}

// NewGetTenantsTenantIDMetricsAPIV1QueryRangeParamsWithContext creates a new GetTenantsTenantIDMetricsAPIV1QueryRangeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantsTenantIDMetricsAPIV1QueryRangeParamsWithContext(ctx context.Context) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	var ()
	return &GetTenantsTenantIDMetricsAPIV1QueryRangeParams{

		Context: ctx,
	}
}

// NewGetTenantsTenantIDMetricsAPIV1QueryRangeParamsWithHTTPClient creates a new GetTenantsTenantIDMetricsAPIV1QueryRangeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantsTenantIDMetricsAPIV1QueryRangeParamsWithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	var ()
	return &GetTenantsTenantIDMetricsAPIV1QueryRangeParams{
		HTTPClient: client,
	}
}

/*GetTenantsTenantIDMetricsAPIV1QueryRangeParams contains all the parameters to send to the API endpoint
for the get tenants tenant ID metrics API v1 query range operation typically these are written to a http.Request
*/
type GetTenantsTenantIDMetricsAPIV1QueryRangeParams struct {

	/*End
	  end time of the requested range (UnixTime or RFC3339)

	*/
	End *string
	/*Query
	  PromQL query to proxy to prometheus

	*/
	Query string
	/*Start
	  start time of the requested range (UnixTime or RFC3339)

	*/
	Start string
	/*Step
	  query range resolution step width

	*/
	Step *string
	/*TenantID
	  Tenant ID

	*/
	TenantID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) WithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) WithContext(ctx context.Context) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) WithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) WithEnd(end *string) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) SetEnd(end *string) {
	o.End = end
}

// WithQuery adds the query to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) WithQuery(query string) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) SetQuery(query string) {
	o.Query = query
}

// WithStart adds the start to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) WithStart(start string) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) SetStart(start string) {
	o.Start = start
}

// WithStep adds the step to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) WithStep(step *string) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	o.SetStep(step)
	return o
}

// SetStep adds the step to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) SetStep(step *string) {
	o.Step = step
}

// WithTenantID adds the tenantID to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) WithTenantID(tenantID int64) *GetTenantsTenantIDMetricsAPIV1QueryRangeParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get tenants tenant ID metrics API v1 query range params
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) SetTenantID(tenantID int64) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantsTenantIDMetricsAPIV1QueryRangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.End != nil {

		// query param end
		var qrEnd string
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := qrEnd
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {
		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	// query param start
	qrStart := o.Start
	qStart := qrStart
	if qStart != "" {
		if err := r.SetQueryParam("start", qStart); err != nil {
			return err
		}
	}

	if o.Step != nil {

		// query param step
		var qrStep string
		if o.Step != nil {
			qrStep = *o.Step
		}
		qStep := qrStep
		if qStep != "" {
			if err := r.SetQueryParam("step", qStep); err != nil {
				return err
			}
		}

	}

	// path param tenant_id
	if err := r.SetPathParam("tenant_id", swag.FormatInt64(o.TenantID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
