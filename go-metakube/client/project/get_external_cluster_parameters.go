// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetExternalClusterParams creates a new GetExternalClusterParams object
// with the default values initialized.
func NewGetExternalClusterParams() *GetExternalClusterParams {
	var ()
	return &GetExternalClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalClusterParamsWithTimeout creates a new GetExternalClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExternalClusterParamsWithTimeout(timeout time.Duration) *GetExternalClusterParams {
	var ()
	return &GetExternalClusterParams{

		timeout: timeout,
	}
}

// NewGetExternalClusterParamsWithContext creates a new GetExternalClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExternalClusterParamsWithContext(ctx context.Context) *GetExternalClusterParams {
	var ()
	return &GetExternalClusterParams{

		Context: ctx,
	}
}

// NewGetExternalClusterParamsWithHTTPClient creates a new GetExternalClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExternalClusterParamsWithHTTPClient(client *http.Client) *GetExternalClusterParams {
	var ()
	return &GetExternalClusterParams{
		HTTPClient: client,
	}
}

/*GetExternalClusterParams contains all the parameters to send to the API endpoint
for the get external cluster operation typically these are written to a http.Request
*/
type GetExternalClusterParams struct {

	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get external cluster params
func (o *GetExternalClusterParams) WithTimeout(timeout time.Duration) *GetExternalClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external cluster params
func (o *GetExternalClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external cluster params
func (o *GetExternalClusterParams) WithContext(ctx context.Context) *GetExternalClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external cluster params
func (o *GetExternalClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external cluster params
func (o *GetExternalClusterParams) WithHTTPClient(client *http.Client) *GetExternalClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external cluster params
func (o *GetExternalClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get external cluster params
func (o *GetExternalClusterParams) WithClusterID(clusterID string) *GetExternalClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get external cluster params
func (o *GetExternalClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the get external cluster params
func (o *GetExternalClusterParams) WithProjectID(projectID string) *GetExternalClusterParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get external cluster params
func (o *GetExternalClusterParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
