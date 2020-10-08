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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Circuit circuit
//
// swagger:model Circuit
type Circuit struct {

	// Circuit ID
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Cid *string `json:"cid"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Commit rate (Kbps)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	CommitRate *int64 `json:"commit_rate,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Date installed
	// Format: date
	InstallDate *strfmt.Date `json:"install_date,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// provider
	// Required: true
	Provider *NestedProvider `json:"provider"`

	// status
	Status *CircuitStatus `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// termination a
	Terminationa *CircuitCircuitTermination `json:"termination_a,omitempty"`

	// termination z
	Terminationz *CircuitCircuitTermination `json:"termination_z,omitempty"`

	// type
	// Required: true
	Type *NestedCircuitType `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this circuit
func (m *Circuit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Circuit) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	if err := validate.MinLength("cid", "body", string(*m.Cid), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("cid", "body", string(*m.Cid), 50); err != nil {
		return err
	}

	return nil
}

func (m *Circuit) validateCommitRate(formats strfmt.Registry) error {

	if swag.IsZero(m.CommitRate) { // not required
		return nil
	}

	if err := validate.MinimumInt("commit_rate", "body", int64(*m.CommitRate), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("commit_rate", "body", int64(*m.CommitRate), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *Circuit) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Circuit) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *Circuit) validateInstallDate(formats strfmt.Registry) error {

	if swag.IsZero(m.InstallDate) { // not required
		return nil
	}

	if err := validate.FormatOf("install_date", "body", "date", m.InstallDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Circuit) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Circuit) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	if m.Provider != nil {
		if err := m.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

func (m *Circuit) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Circuit) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Circuit) validateTenant(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Circuit) validateTerminationa(formats strfmt.Registry) error {

	if swag.IsZero(m.Terminationa) { // not required
		return nil
	}

	if m.Terminationa != nil {
		if err := m.Terminationa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("termination_a")
			}
			return err
		}
	}

	return nil
}

func (m *Circuit) validateTerminationz(formats strfmt.Registry) error {

	if swag.IsZero(m.Terminationz) { // not required
		return nil
	}

	if m.Terminationz != nil {
		if err := m.Terminationz.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("termination_z")
			}
			return err
		}
	}

	return nil
}

func (m *Circuit) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *Circuit) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Circuit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Circuit) UnmarshalBinary(b []byte) error {
	var res Circuit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CircuitStatus Status
//
// swagger:model CircuitStatus
type CircuitStatus struct {

	// label
	// Required: true
	// Enum: [Planned Provisioning Active Offline Deprovisioning Decommissioned]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [planned provisioning active offline deprovisioning decommissioned]
	Value *string `json:"value"`
}

// Validate validates this circuit status
func (m *CircuitStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var circuitStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Planned","Provisioning","Active","Offline","Deprovisioning","Decommissioned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		circuitStatusTypeLabelPropEnum = append(circuitStatusTypeLabelPropEnum, v)
	}
}

const (

	// CircuitStatusLabelPlanned captures enum value "Planned"
	CircuitStatusLabelPlanned string = "Planned"

	// CircuitStatusLabelProvisioning captures enum value "Provisioning"
	CircuitStatusLabelProvisioning string = "Provisioning"

	// CircuitStatusLabelActive captures enum value "Active"
	CircuitStatusLabelActive string = "Active"

	// CircuitStatusLabelOffline captures enum value "Offline"
	CircuitStatusLabelOffline string = "Offline"

	// CircuitStatusLabelDeprovisioning captures enum value "Deprovisioning"
	CircuitStatusLabelDeprovisioning string = "Deprovisioning"

	// CircuitStatusLabelDecommissioned captures enum value "Decommissioned"
	CircuitStatusLabelDecommissioned string = "Decommissioned"
)

// prop value enum
func (m *CircuitStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, circuitStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CircuitStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var circuitStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["planned","provisioning","active","offline","deprovisioning","decommissioned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		circuitStatusTypeValuePropEnum = append(circuitStatusTypeValuePropEnum, v)
	}
}

const (

	// CircuitStatusValuePlanned captures enum value "planned"
	CircuitStatusValuePlanned string = "planned"

	// CircuitStatusValueProvisioning captures enum value "provisioning"
	CircuitStatusValueProvisioning string = "provisioning"

	// CircuitStatusValueActive captures enum value "active"
	CircuitStatusValueActive string = "active"

	// CircuitStatusValueOffline captures enum value "offline"
	CircuitStatusValueOffline string = "offline"

	// CircuitStatusValueDeprovisioning captures enum value "deprovisioning"
	CircuitStatusValueDeprovisioning string = "deprovisioning"

	// CircuitStatusValueDecommissioned captures enum value "decommissioned"
	CircuitStatusValueDecommissioned string = "decommissioned"
)

// prop value enum
func (m *CircuitStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, circuitStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CircuitStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CircuitStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CircuitStatus) UnmarshalBinary(b []byte) error {
	var res CircuitStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
