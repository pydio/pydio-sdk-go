// Code generated by go-swagger; DO NOT EDIT.

package file

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

// NewCreateNodeParams creates a new CreateNodeParams object
// with the default values initialized.
func NewCreateNodeParams() *CreateNodeParams {
	var (
		overrideDefault  = bool(false)
		recursiveDefault = bool(false)
	)
	return &CreateNodeParams{
		Override:  &overrideDefault,
		Recursive: &recursiveDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNodeParamsWithTimeout creates a new CreateNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNodeParamsWithTimeout(timeout time.Duration) *CreateNodeParams {
	var (
		overrideDefault  = bool(false)
		recursiveDefault = bool(false)
	)
	return &CreateNodeParams{
		Override:  &overrideDefault,
		Recursive: &recursiveDefault,

		timeout: timeout,
	}
}

// NewCreateNodeParamsWithContext creates a new CreateNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNodeParamsWithContext(ctx context.Context) *CreateNodeParams {
	var (
		overrideDefault  = bool(false)
		recursiveDefault = bool(false)
	)
	return &CreateNodeParams{
		Override:  &overrideDefault,
		Recursive: &recursiveDefault,

		Context: ctx,
	}
}

// NewCreateNodeParamsWithHTTPClient creates a new CreateNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNodeParamsWithHTTPClient(client *http.Client) *CreateNodeParams {
	var (
		overrideDefault  = bool(false)
		recursiveDefault = bool(false)
	)
	return &CreateNodeParams{
		Override:   &overrideDefault,
		Recursive:  &recursiveDefault,
		HTTPClient: client,
	}
}

/*CreateNodeParams contains all the parameters to send to the API endpoint
for the create node operation typically these are written to a http.Request
*/
type CreateNodeParams struct {

	/*CopySource
	  If it's a move or a copy, indicated the path of the original node. Path must contain the original workspace Id, as it can be used for cross repository copy as well.

	*/
	CopySource *string
	/*DeleteSource
	  If it's a move/rename, will remove original after copy operation

	*/
	DeleteSource *bool
	/*Override
	  Ignore existing resource and override it

	*/
	Override *bool
	/*Path
	  Workspace id or alias + full path to the node, e.g. "/my-files/path/to/node"

	*/
	Path string
	/*Recursive
	  For directories, create parents if necessary

	*/
	Recursive *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create node params
func (o *CreateNodeParams) WithTimeout(timeout time.Duration) *CreateNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create node params
func (o *CreateNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create node params
func (o *CreateNodeParams) WithContext(ctx context.Context) *CreateNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create node params
func (o *CreateNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create node params
func (o *CreateNodeParams) WithHTTPClient(client *http.Client) *CreateNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create node params
func (o *CreateNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCopySource adds the copySource to the create node params
func (o *CreateNodeParams) WithCopySource(copySource *string) *CreateNodeParams {
	o.SetCopySource(copySource)
	return o
}

// SetCopySource adds the copySource to the create node params
func (o *CreateNodeParams) SetCopySource(copySource *string) {
	o.CopySource = copySource
}

// WithDeleteSource adds the deleteSource to the create node params
func (o *CreateNodeParams) WithDeleteSource(deleteSource *bool) *CreateNodeParams {
	o.SetDeleteSource(deleteSource)
	return o
}

// SetDeleteSource adds the deleteSource to the create node params
func (o *CreateNodeParams) SetDeleteSource(deleteSource *bool) {
	o.DeleteSource = deleteSource
}

// WithOverride adds the override to the create node params
func (o *CreateNodeParams) WithOverride(override *bool) *CreateNodeParams {
	o.SetOverride(override)
	return o
}

// SetOverride adds the override to the create node params
func (o *CreateNodeParams) SetOverride(override *bool) {
	o.Override = override
}

// WithPath adds the path to the create node params
func (o *CreateNodeParams) WithPath(path string) *CreateNodeParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the create node params
func (o *CreateNodeParams) SetPath(path string) {
	o.Path = path
}

// WithRecursive adds the recursive to the create node params
func (o *CreateNodeParams) WithRecursive(recursive *bool) *CreateNodeParams {
	o.SetRecursive(recursive)
	return o
}

// SetRecursive adds the recursive to the create node params
func (o *CreateNodeParams) SetRecursive(recursive *bool) {
	o.Recursive = recursive
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CopySource != nil {

		// query param copy_source
		var qrCopySource string
		if o.CopySource != nil {
			qrCopySource = *o.CopySource
		}
		qCopySource := qrCopySource
		if qCopySource != "" {
			if err := r.SetQueryParam("copy_source", qCopySource); err != nil {
				return err
			}
		}

	}

	if o.DeleteSource != nil {

		// query param delete_source
		var qrDeleteSource bool
		if o.DeleteSource != nil {
			qrDeleteSource = *o.DeleteSource
		}
		qDeleteSource := swag.FormatBool(qrDeleteSource)
		if qDeleteSource != "" {
			if err := r.SetQueryParam("delete_source", qDeleteSource); err != nil {
				return err
			}
		}

	}

	if o.Override != nil {

		// query param override
		var qrOverride bool
		if o.Override != nil {
			qrOverride = *o.Override
		}
		qOverride := swag.FormatBool(qrOverride)
		if qOverride != "" {
			if err := r.SetQueryParam("override", qOverride); err != nil {
				return err
			}
		}

	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if o.Recursive != nil {

		// query param recursive
		var qrRecursive bool
		if o.Recursive != nil {
			qrRecursive = *o.Recursive
		}
		qRecursive := swag.FormatBool(qrRecursive)
		if qRecursive != "" {
			if err := r.SetQueryParam("recursive", qRecursive); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
