// Code generated by go-swagger; DO NOT EDIT.

package networks

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
)

// NewAllocateOrReleaseFromIPPoolParams creates a new AllocateOrReleaseFromIPPoolParams object
// with the default values initialized.
func NewAllocateOrReleaseFromIPPoolParams() *AllocateOrReleaseFromIPPoolParams {
	var ()
	return &AllocateOrReleaseFromIPPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllocateOrReleaseFromIPPoolParamsWithTimeout creates a new AllocateOrReleaseFromIPPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllocateOrReleaseFromIPPoolParamsWithTimeout(timeout time.Duration) *AllocateOrReleaseFromIPPoolParams {
	var ()
	return &AllocateOrReleaseFromIPPoolParams{

		timeout: timeout,
	}
}

// NewAllocateOrReleaseFromIPPoolParamsWithContext creates a new AllocateOrReleaseFromIPPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllocateOrReleaseFromIPPoolParamsWithContext(ctx context.Context) *AllocateOrReleaseFromIPPoolParams {
	var ()
	return &AllocateOrReleaseFromIPPoolParams{

		Context: ctx,
	}
}

// NewAllocateOrReleaseFromIPPoolParamsWithHTTPClient creates a new AllocateOrReleaseFromIPPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllocateOrReleaseFromIPPoolParamsWithHTTPClient(client *http.Client) *AllocateOrReleaseFromIPPoolParams {
	var ()
	return &AllocateOrReleaseFromIPPoolParams{
		HTTPClient: client,
	}
}

/*AllocateOrReleaseFromIPPoolParams contains all the parameters to send to the API endpoint
for the allocate or release from Ip pool operation typically these are written to a http.Request
*/
type AllocateOrReleaseFromIPPoolParams struct {

	/*Action*/
	Action *string
	/*AllocatedIP*/
	AllocatedIP interface{}
	/*PoolID*/
	PoolID *string
	/*ReleasedIP*/
	ReleasedIP *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) WithTimeout(timeout time.Duration) *AllocateOrReleaseFromIPPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) WithContext(ctx context.Context) *AllocateOrReleaseFromIPPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) WithHTTPClient(client *http.Client) *AllocateOrReleaseFromIPPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) WithAction(action *string) *AllocateOrReleaseFromIPPoolParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) SetAction(action *string) {
	o.Action = action
}

// WithAllocatedIP adds the allocatedIP to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) WithAllocatedIP(allocatedIP interface{}) *AllocateOrReleaseFromIPPoolParams {
	o.SetAllocatedIP(allocatedIP)
	return o
}

// SetAllocatedIP adds the allocatedIp to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) SetAllocatedIP(allocatedIP interface{}) {
	o.AllocatedIP = allocatedIP
}

// WithPoolID adds the poolID to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) WithPoolID(poolID *string) *AllocateOrReleaseFromIPPoolParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) SetPoolID(poolID *string) {
	o.PoolID = poolID
}

// WithReleasedIP adds the releasedIP to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) WithReleasedIP(releasedIP *string) *AllocateOrReleaseFromIPPoolParams {
	o.SetReleasedIP(releasedIP)
	return o
}

// SetReleasedIP adds the releasedIp to the allocate or release from Ip pool params
func (o *AllocateOrReleaseFromIPPoolParams) SetReleasedIP(releasedIP *string) {
	o.ReleasedIP = releasedIP
}

// WriteToRequest writes these params to a swagger request
func (o *AllocateOrReleaseFromIPPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string
		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {
			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}

	}

	if o.AllocatedIP != nil {
		if err := r.SetBodyParam(o.AllocatedIP); err != nil {
			return err
		}
	}

	if o.PoolID != nil {

		// path param poolId
		if err := r.SetPathParam("poolId", *o.PoolID); err != nil {
			return err
		}

	}

	if o.ReleasedIP != nil {

		// query param releasedIP
		var qrReleasedIP string
		if o.ReleasedIP != nil {
			qrReleasedIP = *o.ReleasedIP
		}
		qReleasedIP := qrReleasedIP
		if qReleasedIP != "" {
			if err := r.SetQueryParam("releasedIP", qReleasedIP); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}