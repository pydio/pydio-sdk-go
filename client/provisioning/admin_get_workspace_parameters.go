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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAdminGetWorkspaceParams creates a new AdminGetWorkspaceParams object
// with the default values initialized.
func NewAdminGetWorkspaceParams() *AdminGetWorkspaceParams {
	var ()
	return &AdminGetWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetWorkspaceParamsWithTimeout creates a new AdminGetWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetWorkspaceParamsWithTimeout(timeout time.Duration) *AdminGetWorkspaceParams {
	var ()
	return &AdminGetWorkspaceParams{

		timeout: timeout,
	}
}

// NewAdminGetWorkspaceParamsWithContext creates a new AdminGetWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetWorkspaceParamsWithContext(ctx context.Context) *AdminGetWorkspaceParams {
	var ()
	return &AdminGetWorkspaceParams{

		Context: ctx,
	}
}

// NewAdminGetWorkspaceParamsWithHTTPClient creates a new AdminGetWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetWorkspaceParamsWithHTTPClient(client *http.Client) *AdminGetWorkspaceParams {
	var ()
	return &AdminGetWorkspaceParams{
		HTTPClient: client,
	}
}

/*AdminGetWorkspaceParams contains all the parameters to send to the API endpoint
for the admin get workspace operation typically these are written to a http.Request
*/
type AdminGetWorkspaceParams struct {

	/*Format
	  Format produced in output (defaults to xml)

	*/
	Format *string
	/*LoadFillValues
	  Load additional data to build a form for editing this role

	*/
	LoadFillValues *bool
	/*WorkspaceID
	  Id or Alias / Load detail of this workspace

	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin get workspace params
func (o *AdminGetWorkspaceParams) WithTimeout(timeout time.Duration) *AdminGetWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get workspace params
func (o *AdminGetWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get workspace params
func (o *AdminGetWorkspaceParams) WithContext(ctx context.Context) *AdminGetWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get workspace params
func (o *AdminGetWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get workspace params
func (o *AdminGetWorkspaceParams) WithHTTPClient(client *http.Client) *AdminGetWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get workspace params
func (o *AdminGetWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the admin get workspace params
func (o *AdminGetWorkspaceParams) WithFormat(format *string) *AdminGetWorkspaceParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the admin get workspace params
func (o *AdminGetWorkspaceParams) SetFormat(format *string) {
	o.Format = format
}

// WithLoadFillValues adds the loadFillValues to the admin get workspace params
func (o *AdminGetWorkspaceParams) WithLoadFillValues(loadFillValues *bool) *AdminGetWorkspaceParams {
	o.SetLoadFillValues(loadFillValues)
	return o
}

// SetLoadFillValues adds the loadFillValues to the admin get workspace params
func (o *AdminGetWorkspaceParams) SetLoadFillValues(loadFillValues *bool) {
	o.LoadFillValues = loadFillValues
}

// WithWorkspaceID adds the workspaceID to the admin get workspace params
func (o *AdminGetWorkspaceParams) WithWorkspaceID(workspaceID string) *AdminGetWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the admin get workspace params
func (o *AdminGetWorkspaceParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	if o.LoadFillValues != nil {

		// query param load_fill_values
		var qrLoadFillValues bool
		if o.LoadFillValues != nil {
			qrLoadFillValues = *o.LoadFillValues
		}
		qLoadFillValues := swag.FormatBool(qrLoadFillValues)
		if qLoadFillValues != "" {
			if err := r.SetQueryParam("load_fill_values", qLoadFillValues); err != nil {
				return err
			}
		}

	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
