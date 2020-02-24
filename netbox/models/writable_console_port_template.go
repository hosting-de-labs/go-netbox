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
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableConsolePortTemplate writable console port template
// swagger:model WritableConsolePortTemplate
type WritableConsolePortTemplate struct {

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Type
	// Enum: [de-9 db-25 rj-12 rj-45 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b other]
	Type string `json:"type,omitempty"`
}

// Validate validates this writable console port template
func (m *WritableConsolePortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableConsolePortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

var writableConsolePortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-12","rj-45","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableConsolePortTemplateTypeTypePropEnum = append(writableConsolePortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableConsolePortTemplateTypeDe9 captures enum value "de-9"
	WritableConsolePortTemplateTypeDe9 string = "de-9"

	// WritableConsolePortTemplateTypeDb25 captures enum value "db-25"
	WritableConsolePortTemplateTypeDb25 string = "db-25"

	// WritableConsolePortTemplateTypeRj12 captures enum value "rj-12"
	WritableConsolePortTemplateTypeRj12 string = "rj-12"

	// WritableConsolePortTemplateTypeRj45 captures enum value "rj-45"
	WritableConsolePortTemplateTypeRj45 string = "rj-45"

	// WritableConsolePortTemplateTypeUsba captures enum value "usb-a"
	WritableConsolePortTemplateTypeUsba string = "usb-a"

	// WritableConsolePortTemplateTypeUsbb captures enum value "usb-b"
	WritableConsolePortTemplateTypeUsbb string = "usb-b"

	// WritableConsolePortTemplateTypeUsbc captures enum value "usb-c"
	WritableConsolePortTemplateTypeUsbc string = "usb-c"

	// WritableConsolePortTemplateTypeUsbMinia captures enum value "usb-mini-a"
	WritableConsolePortTemplateTypeUsbMinia string = "usb-mini-a"

	// WritableConsolePortTemplateTypeUsbMinib captures enum value "usb-mini-b"
	WritableConsolePortTemplateTypeUsbMinib string = "usb-mini-b"

	// WritableConsolePortTemplateTypeUsbMicroa captures enum value "usb-micro-a"
	WritableConsolePortTemplateTypeUsbMicroa string = "usb-micro-a"

	// WritableConsolePortTemplateTypeUsbMicrob captures enum value "usb-micro-b"
	WritableConsolePortTemplateTypeUsbMicrob string = "usb-micro-b"

	// WritableConsolePortTemplateTypeOther captures enum value "other"
	WritableConsolePortTemplateTypeOther string = "other"
)

// prop value enum
func (m *WritableConsolePortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableConsolePortTemplateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableConsolePortTemplate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableConsolePortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableConsolePortTemplate) UnmarshalBinary(b []byte) error {
	var res WritableConsolePortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
