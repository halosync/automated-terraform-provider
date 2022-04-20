// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetDeviceListParams creates a new GetDeviceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceListParams() *GetDeviceListParams {
	return &GetDeviceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceListParamsWithTimeout creates a new GetDeviceListParams object
// with the ability to set a timeout on a request.
func NewGetDeviceListParamsWithTimeout(timeout time.Duration) *GetDeviceListParams {
	return &GetDeviceListParams{
		timeout: timeout,
	}
}

// NewGetDeviceListParamsWithContext creates a new GetDeviceListParams object
// with the ability to set a context for a request.
func NewGetDeviceListParamsWithContext(ctx context.Context) *GetDeviceListParams {
	return &GetDeviceListParams{
		Context: ctx,
	}
}

// NewGetDeviceListParamsWithHTTPClient creates a new GetDeviceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceListParamsWithHTTPClient(client *http.Client) *GetDeviceListParams {
	return &GetDeviceListParams{
		HTTPClient: client,
	}
}

/* GetDeviceListParams contains all the parameters to send to the API endpoint
   for the get device list operation.

   Typically these are written to a http.Request.
*/
type GetDeviceListParams struct {

	// End.
	//
	// Format: int64
	End *int64

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// NetflowFilter.
	NetflowFilter *string

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceListParams) WithDefaults() *GetDeviceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceListParams) SetDefaults() {
	var (
		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetDeviceListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device list params
func (o *GetDeviceListParams) WithTimeout(timeout time.Duration) *GetDeviceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device list params
func (o *GetDeviceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device list params
func (o *GetDeviceListParams) WithContext(ctx context.Context) *GetDeviceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device list params
func (o *GetDeviceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device list params
func (o *GetDeviceListParams) WithHTTPClient(client *http.Client) *GetDeviceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device list params
func (o *GetDeviceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get device list params
func (o *GetDeviceListParams) WithEnd(end *int64) *GetDeviceListParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get device list params
func (o *GetDeviceListParams) SetEnd(end *int64) {
	o.End = end
}

// WithFields adds the fields to the get device list params
func (o *GetDeviceListParams) WithFields(fields *string) *GetDeviceListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device list params
func (o *GetDeviceListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device list params
func (o *GetDeviceListParams) WithFilter(filter *string) *GetDeviceListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device list params
func (o *GetDeviceListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithNetflowFilter adds the netflowFilter to the get device list params
func (o *GetDeviceListParams) WithNetflowFilter(netflowFilter *string) *GetDeviceListParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the get device list params
func (o *GetDeviceListParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithOffset adds the offset to the get device list params
func (o *GetDeviceListParams) WithOffset(offset *int32) *GetDeviceListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device list params
func (o *GetDeviceListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device list params
func (o *GetDeviceListParams) WithSize(size *int32) *GetDeviceListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device list params
func (o *GetDeviceListParams) SetSize(size *int32) {
	o.Size = size
}

// WithStart adds the start to the get device list params
func (o *GetDeviceListParams) WithStart(start *int64) *GetDeviceListParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get device list params
func (o *GetDeviceListParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.NetflowFilter != nil {

		// query param netflowFilter
		var qrNetflowFilter string

		if o.NetflowFilter != nil {
			qrNetflowFilter = *o.NetflowFilter
		}
		qNetflowFilter := qrNetflowFilter
		if qNetflowFilter != "" {

			if err := r.SetQueryParam("netflowFilter", qNetflowFilter); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
