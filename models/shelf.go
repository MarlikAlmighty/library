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

// Shelf shelf
// swagger:model Shelf
type Shelf struct {

	// books
	Books []*Book `json:"Books"`

	// ID
	ID int32 `json:"ID,omitempty"`

	// name shelfs
	NameShelfs string `json:"NameShelfs,omitempty"`

	// number
	Number int32 `json:"Number,omitempty"`
}

// Validate validates this shelf
func (m *Shelf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBooks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Shelf) validateBooks(formats strfmt.Registry) error {

	if swag.IsZero(m.Books) { // not required
		return nil
	}

	for i := 0; i < len(m.Books); i++ {
		if swag.IsZero(m.Books[i]) { // not required
			continue
		}

		if m.Books[i] != nil {
			if err := m.Books[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Books" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Shelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Shelf) UnmarshalBinary(b []byte) error {
	var res Shelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}