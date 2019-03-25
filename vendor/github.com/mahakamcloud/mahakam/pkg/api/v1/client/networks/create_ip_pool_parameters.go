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

	models "github.com/mahakamcloud/mahakam/pkg/api/v1/models"
)

// NewCreateIPPoolParams creates a new CreateIPPoolParams object
// with the default values initialized.
func NewCreateIPPoolParams() *CreateIPPoolParams {
	var ()
	return &CreateIPPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIPPoolParamsWithTimeout creates a new CreateIPPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateIPPoolParamsWithTimeout(timeout time.Duration) *CreateIPPoolParams {
	var ()
	return &CreateIPPoolParams{

		timeout: timeout,
	}
}

// NewCreateIPPoolParamsWithContext creates a new CreateIPPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateIPPoolParamsWithContext(ctx context.Context) *CreateIPPoolParams {
	var ()
	return &CreateIPPoolParams{

		Context: ctx,
	}
}

// NewCreateIPPoolParamsWithHTTPClient creates a new CreateIPPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateIPPoolParamsWithHTTPClient(client *http.Client) *CreateIPPoolParams {
	var ()
	return &CreateIPPoolParams{
		HTTPClient: client,
	}
}

/*CreateIPPoolParams contains all the parameters to send to the API endpoint
for the create Ip pool operation typically these are written to a http.Request
*/
type CreateIPPoolParams struct {

	/*IPPool*/
	IPPool *models.IPPool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create Ip pool params
func (o *CreateIPPoolParams) WithTimeout(timeout time.Duration) *CreateIPPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Ip pool params
func (o *CreateIPPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Ip pool params
func (o *CreateIPPoolParams) WithContext(ctx context.Context) *CreateIPPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Ip pool params
func (o *CreateIPPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Ip pool params
func (o *CreateIPPoolParams) WithHTTPClient(client *http.Client) *CreateIPPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Ip pool params
func (o *CreateIPPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPPool adds the iPPool to the create Ip pool params
func (o *CreateIPPoolParams) WithIPPool(iPPool *models.IPPool) *CreateIPPoolParams {
	o.SetIPPool(iPPool)
	return o
}

// SetIPPool adds the ipPool to the create Ip pool params
func (o *CreateIPPoolParams) SetIPPool(iPPool *models.IPPool) {
	o.IPPool = iPPool
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIPPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IPPool != nil {
		if err := r.SetBodyParam(o.IPPool); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
