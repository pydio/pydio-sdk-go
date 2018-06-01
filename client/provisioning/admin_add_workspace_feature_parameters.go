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

	models "github.com/pydio/pydio-sdk-go/models"
)

// NewAdminAddWorkspaceFeatureParams creates a new AdminAddWorkspaceFeatureParams object
// with the default values initialized.
func NewAdminAddWorkspaceFeatureParams() *AdminAddWorkspaceFeatureParams {
	var ()
	return &AdminAddWorkspaceFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminAddWorkspaceFeatureParamsWithTimeout creates a new AdminAddWorkspaceFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminAddWorkspaceFeatureParamsWithTimeout(timeout time.Duration) *AdminAddWorkspaceFeatureParams {
	var ()
	return &AdminAddWorkspaceFeatureParams{

		timeout: timeout,
	}
}

// NewAdminAddWorkspaceFeatureParamsWithContext creates a new AdminAddWorkspaceFeatureParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminAddWorkspaceFeatureParamsWithContext(ctx context.Context) *AdminAddWorkspaceFeatureParams {
	var ()
	return &AdminAddWorkspaceFeatureParams{

		Context: ctx,
	}
}

// NewAdminAddWorkspaceFeatureParamsWithHTTPClient creates a new AdminAddWorkspaceFeatureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminAddWorkspaceFeatureParamsWithHTTPClient(client *http.Client) *AdminAddWorkspaceFeatureParams {
	var ()
	return &AdminAddWorkspaceFeatureParams{
		HTTPClient: client,
	}
}

/*AdminAddWorkspaceFeatureParams contains all the parameters to send to the API endpoint
for the admin add workspace feature operation typically these are written to a http.Request
*/
type AdminAddWorkspaceFeatureParams struct {

	/*MetaID
	  plugin id for new meta

	*/
	MetaID string
	/*Parameters
	  Meta source parameters

	*/
	Parameters models.MetaSourceParameters
	/*WorkspaceID
	  id or alias of this workspace

	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) WithTimeout(timeout time.Duration) *AdminAddWorkspaceFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) WithContext(ctx context.Context) *AdminAddWorkspaceFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) WithHTTPClient(client *http.Client) *AdminAddWorkspaceFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetaID adds the metaID to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) WithMetaID(metaID string) *AdminAddWorkspaceFeatureParams {
	o.SetMetaID(metaID)
	return o
}

// SetMetaID adds the metaId to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) SetMetaID(metaID string) {
	o.MetaID = metaID
}

// WithParameters adds the parameters to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) WithParameters(parameters models.MetaSourceParameters) *AdminAddWorkspaceFeatureParams {
	o.SetParameters(parameters)
	return o
}

// SetParameters adds the parameters to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) SetParameters(parameters models.MetaSourceParameters) {
	o.Parameters = parameters
}

// WithWorkspaceID adds the workspaceID to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) WithWorkspaceID(workspaceID string) *AdminAddWorkspaceFeatureParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the admin add workspace feature params
func (o *AdminAddWorkspaceFeatureParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminAddWorkspaceFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param metaId
	if err := r.SetPathParam("metaId", o.MetaID); err != nil {
		return err
	}

	if o.Parameters != nil {
		if err := r.SetBodyParam(o.Parameters); err != nil {
			return err
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
