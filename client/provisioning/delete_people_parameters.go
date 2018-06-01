// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePeopleParams creates a new DeletePeopleParams object
// with the default values initialized.
func NewDeletePeopleParams() *DeletePeopleParams {
	var ()
	return &DeletePeopleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePeopleParamsWithTimeout creates a new DeletePeopleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePeopleParamsWithTimeout(timeout time.Duration) *DeletePeopleParams {
	var ()
	return &DeletePeopleParams{

		timeout: timeout,
	}
}

// NewDeletePeopleParamsWithContext creates a new DeletePeopleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePeopleParamsWithContext(ctx context.Context) *DeletePeopleParams {
	var ()
	return &DeletePeopleParams{

		Context: ctx,
	}
}

// NewDeletePeopleParamsWithHTTPClient creates a new DeletePeopleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePeopleParamsWithHTTPClient(client *http.Client) *DeletePeopleParams {
	var ()
	return &DeletePeopleParams{
		HTTPClient: client,
	}
}

/*DeletePeopleParams contains all the parameters to send to the API endpoint
for the delete people operation typically these are written to a http.Request
*/
type DeletePeopleParams struct {

	/*Path
	  User or group identifier, including full group path

	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete people params
func (o *DeletePeopleParams) WithTimeout(timeout time.Duration) *DeletePeopleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete people params
func (o *DeletePeopleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete people params
func (o *DeletePeopleParams) WithContext(ctx context.Context) *DeletePeopleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete people params
func (o *DeletePeopleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete people params
func (o *DeletePeopleParams) WithHTTPClient(client *http.Client) *DeletePeopleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete people params
func (o *DeletePeopleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the delete people params
func (o *DeletePeopleParams) WithPath(path string) *DeletePeopleParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the delete people params
func (o *DeletePeopleParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePeopleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
