// Code generated by go-docs; DO NOT EDIT.

package models

// This file was generated by the docs tool.
// Editing this file might prove futile when you re-run the docs generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Page page
// docs:model Page
type Page struct {

	// books ID
	BooksID int32 `json:"BooksID,omitempty"`

	// ID
	ID int32 `json:"ID,omitempty"`

	// number
	Number int32 `json:"Number,omitempty"`

	// text
	Text string `json:"Text,omitempty"`
}

// Validate validates this page
func (m *Page) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Page) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Page) UnmarshalBinary(b []byte) error {
	var res Page
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
