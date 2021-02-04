// Code generated by go-swagger; DO NOT EDIT.

package datacenter

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
)

// NewDeleteDCParams creates a new DeleteDCParams object
// with the default values initialized.
func NewDeleteDCParams() *DeleteDCParams {
	var ()
	return &DeleteDCParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDCParamsWithTimeout creates a new DeleteDCParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDCParamsWithTimeout(timeout time.Duration) *DeleteDCParams {
	var ()
	return &DeleteDCParams{

		timeout: timeout,
	}
}

// NewDeleteDCParamsWithContext creates a new DeleteDCParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDCParamsWithContext(ctx context.Context) *DeleteDCParams {
	var ()
	return &DeleteDCParams{

		Context: ctx,
	}
}

// NewDeleteDCParamsWithHTTPClient creates a new DeleteDCParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDCParamsWithHTTPClient(client *http.Client) *DeleteDCParams {
	var ()
	return &DeleteDCParams{
		HTTPClient: client,
	}
}

/*DeleteDCParams contains all the parameters to send to the API endpoint
for the delete d c operation typically these are written to a http.Request
*/
type DeleteDCParams struct {

	/*Dc*/
	DC string
	/*SeedName*/
	Seed string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete d c params
func (o *DeleteDCParams) WithTimeout(timeout time.Duration) *DeleteDCParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete d c params
func (o *DeleteDCParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete d c params
func (o *DeleteDCParams) WithContext(ctx context.Context) *DeleteDCParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete d c params
func (o *DeleteDCParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete d c params
func (o *DeleteDCParams) WithHTTPClient(client *http.Client) *DeleteDCParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete d c params
func (o *DeleteDCParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDC adds the dc to the delete d c params
func (o *DeleteDCParams) WithDC(dc string) *DeleteDCParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the delete d c params
func (o *DeleteDCParams) SetDC(dc string) {
	o.DC = dc
}

// WithSeed adds the seedName to the delete d c params
func (o *DeleteDCParams) WithSeed(seedName string) *DeleteDCParams {
	o.SetSeed(seedName)
	return o
}

// SetSeed adds the seedName to the delete d c params
func (o *DeleteDCParams) SetSeed(seedName string) {
	o.Seed = seedName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDCParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param seed_name
	if err := r.SetPathParam("seed_name", o.Seed); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
