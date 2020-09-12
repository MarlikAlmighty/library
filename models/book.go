// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Book book
// swagger:model Book
type Book struct {

	// author
	Author string `json:"Author,omitempty"`

	// ID
	ID int32 `json:"ID,omitempty"`

	// num pages
	NumPages int32 `json:"NumPages,omitempty"`

	// pages
	Pages []*Page `json:"Pages"`

	// publisher
	Publisher string `json:"Publisher,omitempty"`

	// shelves ID
	ShelvesID int32 `json:"ShelvesID,omitempty"`
}

// Validate validates this book
func (m *Book) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Book) validatePages(formats strfmt.Registry) error {

	if swag.IsZero(m.Pages) { // not required
		return nil
	}

	for i := 0; i < len(m.Pages); i++ {
		if swag.IsZero(m.Pages[i]) { // not required
			continue
		}

		if m.Pages[i] != nil {
			if err := m.Pages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Pages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Book) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Book) UnmarshalBinary(b []byte) error {
	var res Book
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
