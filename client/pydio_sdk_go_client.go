// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pydio/pydio-sdk-go/client/file"
	"github.com/pydio/pydio-sdk-go/client/provisioning"
	"github.com/pydio/pydio-sdk-go/client/task"
	"github.com/pydio/pydio-sdk-go/client/user_account"
	"github.com/pydio/pydio-sdk-go/client/workspace"
)

// Default pydio sdk go HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v2"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new pydio sdk go HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PydioSdkGo {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new pydio sdk go HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PydioSdkGo {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new pydio sdk go client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PydioSdkGo {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PydioSdkGo)
	cli.Transport = transport

	cli.File = file.New(transport, formats)

	cli.Provisioning = provisioning.New(transport, formats)

	cli.Task = task.New(transport, formats)

	cli.UserAccount = user_account.New(transport, formats)

	cli.Workspace = workspace.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PydioSdkGo is a client for pydio sdk go
type PydioSdkGo struct {
	File *file.Client

	Provisioning *provisioning.Client

	Task *task.Client

	UserAccount *user_account.Client

	Workspace *workspace.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PydioSdkGo) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.File.SetTransport(transport)

	c.Provisioning.SetTransport(transport)

	c.Task.SetTransport(transport)

	c.UserAccount.SetTransport(transport)

	c.Workspace.SetTransport(transport)

}
