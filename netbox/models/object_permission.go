// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ObjectPermission object permission
//
// swagger:model ObjectPermission
type ObjectPermission struct {

	// The list of actions granted by this permission
	// Required: true
	Actions []string `json:"actions"`

	// Constraints
	//
	// Queryset filter matching the applicable objects of the selected type(s)
	Constraints *string `json:"constraints,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// groups
	// Unique: true
	Groups []*NestedGroup `json:"groups"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// object types
	// Required: true
	// Unique: true
	ObjectTypes []string `json:"object_types"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// users
	// Unique: true
	Users []*NestedUser `json:"users"`
}

// Validate validates this object permission
func (m *ObjectPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectPermission) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("actions", "body", m.Actions); err != nil {
		return err
	}

	for i := 0; i < len(m.Actions); i++ {

		if err := validate.MinLength("actions"+"."+strconv.Itoa(i), "body", string(m.Actions[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("actions"+"."+strconv.Itoa(i), "body", string(m.Actions[i]), 30); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectPermission) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *ObjectPermission) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	if err := validate.UniqueItems("groups", "body", m.Groups); err != nil {
		return err
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectPermission) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *ObjectPermission) validateObjectTypes(formats strfmt.Registry) error {

	if err := validate.Required("object_types", "body", m.ObjectTypes); err != nil {
		return err
	}

	if err := validate.UniqueItems("object_types", "body", m.ObjectTypes); err != nil {
		return err
	}

	return nil
}

func (m *ObjectPermission) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectPermission) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.Users) { // not required
		return nil
	}

	if err := validate.UniqueItems("users", "body", m.Users); err != nil {
		return err
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectPermission) UnmarshalBinary(b []byte) error {
	var res ObjectPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
