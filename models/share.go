// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Share share
// swagger:model Share
type Share struct {

	// ajxp mime
	AjxpMime string `json:"ajxp_mime,omitempty"`

	// ajxp shared
	AjxpShared bool `json:"ajxp_shared,omitempty"`

	// ajxp shared minisite
	AjxpSharedMinisite string `json:"ajxp_shared_minisite,omitempty"`

	// fonticon
	Fonticon string `json:"fonticon,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// openicon
	Openicon string `json:"openicon,omitempty"`

	// original path
	OriginalPath string `json:"original_path,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// share data
	ShareData string `json:"share_data,omitempty"`

	// share element parent repository label
	ShareElementParentRepositoryLabel string `json:"share_element_parent_repository_label,omitempty"`

	// share type
	ShareType string `json:"share_type,omitempty"`

	// share type readable
	ShareTypeReadable string `json:"share_type_readable,omitempty"`

	// shared element hash
	SharedElementHash string `json:"shared_element_hash,omitempty"`

	// shared element parent repository
	SharedElementParentRepository string `json:"shared_element_parent_repository,omitempty"`
}

// Validate validates this share
func (m *Share) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Share) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Share) UnmarshalBinary(b []byte) error {
	var res Share
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
