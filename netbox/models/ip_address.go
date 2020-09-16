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
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPAddress IP address
// swagger:model IPAddress
type IPAddress struct {

	// Address
	//
	// IPv4 or IPv6 address (with mask)
	// Required: true
	Address *string `json:"address"`

	// Assigned object
	// Read Only: true
	AssignedObject map[string]string `json:"assigned_object,omitempty"`

	// Assigned object id
	// Maximum: 2.147483647e+09
	// Minimum: 0
	AssignedObjectID *int64 `json:"assigned_object_id,omitempty"`

	// Assigned object type
	AssignedObjectType string `json:"assigned_object_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// DNS Name
	//
	// Hostname or FQDN (not case-sensitive)
	// Max Length: 255
	// Pattern: ^[0-9A-Za-z._-]+$
	DNSName string `json:"dns_name,omitempty"`

	// family
	Family *IPAddressFamily `json:"family,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// nat inside
	NatInside *NestedIPAddress `json:"nat_inside,omitempty"`

	// nat outside
	NatOutside *NestedIPAddress `json:"nat_outside,omitempty"`

	// role
	Role *IPAddressRole `json:"role,omitempty"`

	// status
	Status *IPAddressStatus `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// vrf
	Vrf *NestedVRF `json:"vrf,omitempty"`
}

// Validate validates this IP address
func (m *IPAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatInside(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatOutside(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
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

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVrf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAddress) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *IPAddress) validateAssignedObjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignedObjectID) { // not required
		return nil
	}

	if err := validate.MinimumInt("assigned_object_id", "body", int64(*m.AssignedObjectID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("assigned_object_id", "body", int64(*m.AssignedObjectID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *IPAddress) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPAddress) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *IPAddress) validateDNSName(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSName) { // not required
		return nil
	}

	if err := validate.MaxLength("dns_name", "body", string(m.DNSName), 255); err != nil {
		return err
	}

	if err := validate.Pattern("dns_name", "body", string(m.DNSName), `^[0-9A-Za-z._-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *IPAddress) validateFamily(formats strfmt.Registry) error {

	if swag.IsZero(m.Family) { // not required
		return nil
	}

	if m.Family != nil {
		if err := m.Family.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("family")
			}
			return err
		}
	}

	return nil
}

func (m *IPAddress) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPAddress) validateNatInside(formats strfmt.Registry) error {

	if swag.IsZero(m.NatInside) { // not required
		return nil
	}

	if m.NatInside != nil {
		if err := m.NatInside.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nat_inside")
			}
			return err
		}
	}

	return nil
}

func (m *IPAddress) validateNatOutside(formats strfmt.Registry) error {

	if swag.IsZero(m.NatOutside) { // not required
		return nil
	}

	if m.NatOutside != nil {
		if err := m.NatOutside.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nat_outside")
			}
			return err
		}
	}

	return nil
}

func (m *IPAddress) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *IPAddress) validateStatus(formats strfmt.Registry) error {

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

func (m *IPAddress) validateTags(formats strfmt.Registry) error {

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

func (m *IPAddress) validateTenant(formats strfmt.Registry) error {

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

func (m *IPAddress) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPAddress) validateVrf(formats strfmt.Registry) error {

	if swag.IsZero(m.Vrf) { // not required
		return nil
	}

	if m.Vrf != nil {
		if err := m.Vrf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vrf")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddress) UnmarshalBinary(b []byte) error {
	var res IPAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPAddressFamily Family
// swagger:model IPAddressFamily
type IPAddressFamily struct {

	// label
	// Required: true
	// Enum: [IPv4 IPv6]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [4 6]
	Value *int64 `json:"value"`
}

// Validate validates this IP address family
func (m *IPAddressFamily) Validate(formats strfmt.Registry) error {
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

var ipAddressFamilyTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressFamilyTypeLabelPropEnum = append(ipAddressFamilyTypeLabelPropEnum, v)
	}
}

const (

	// IPAddressFamilyLabelIPV4 captures enum value "IPv4"
	IPAddressFamilyLabelIPV4 string = "IPv4"

	// IPAddressFamilyLabelIPV6 captures enum value "IPv6"
	IPAddressFamilyLabelIPV6 string = "IPv6"
)

// prop value enum
func (m *IPAddressFamily) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipAddressFamilyTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPAddressFamily) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("family"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("family"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var ipAddressFamilyTypeValuePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[4,6]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressFamilyTypeValuePropEnum = append(ipAddressFamilyTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *IPAddressFamily) validateValueEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, ipAddressFamilyTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPAddressFamily) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("family"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("family"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAddressFamily) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddressFamily) UnmarshalBinary(b []byte) error {
	var res IPAddressFamily
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPAddressRole Role
// swagger:model IPAddressRole
type IPAddressRole struct {

	// label
	// Required: true
	// Enum: [Loopback Secondary Anycast VIP VRRP HSRP GLBP CARP]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [loopback secondary anycast vip vrrp hsrp glbp carp]
	Value *string `json:"value"`
}

// Validate validates this IP address role
func (m *IPAddressRole) Validate(formats strfmt.Registry) error {
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

var ipAddressRoleTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Loopback","Secondary","Anycast","VIP","VRRP","HSRP","GLBP","CARP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressRoleTypeLabelPropEnum = append(ipAddressRoleTypeLabelPropEnum, v)
	}
}

const (

	// IPAddressRoleLabelLoopback captures enum value "Loopback"
	IPAddressRoleLabelLoopback string = "Loopback"

	// IPAddressRoleLabelSecondary captures enum value "Secondary"
	IPAddressRoleLabelSecondary string = "Secondary"

	// IPAddressRoleLabelAnycast captures enum value "Anycast"
	IPAddressRoleLabelAnycast string = "Anycast"

	// IPAddressRoleLabelVIP captures enum value "VIP"
	IPAddressRoleLabelVIP string = "VIP"

	// IPAddressRoleLabelVRRP captures enum value "VRRP"
	IPAddressRoleLabelVRRP string = "VRRP"

	// IPAddressRoleLabelHSRP captures enum value "HSRP"
	IPAddressRoleLabelHSRP string = "HSRP"

	// IPAddressRoleLabelGLBP captures enum value "GLBP"
	IPAddressRoleLabelGLBP string = "GLBP"

	// IPAddressRoleLabelCARP captures enum value "CARP"
	IPAddressRoleLabelCARP string = "CARP"
)

// prop value enum
func (m *IPAddressRole) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipAddressRoleTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPAddressRole) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("role"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("role"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var ipAddressRoleTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["loopback","secondary","anycast","vip","vrrp","hsrp","glbp","carp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressRoleTypeValuePropEnum = append(ipAddressRoleTypeValuePropEnum, v)
	}
}

const (

	// IPAddressRoleValueLoopback captures enum value "loopback"
	IPAddressRoleValueLoopback string = "loopback"

	// IPAddressRoleValueSecondary captures enum value "secondary"
	IPAddressRoleValueSecondary string = "secondary"

	// IPAddressRoleValueAnycast captures enum value "anycast"
	IPAddressRoleValueAnycast string = "anycast"

	// IPAddressRoleValueVip captures enum value "vip"
	IPAddressRoleValueVip string = "vip"

	// IPAddressRoleValueVrrp captures enum value "vrrp"
	IPAddressRoleValueVrrp string = "vrrp"

	// IPAddressRoleValueHsrp captures enum value "hsrp"
	IPAddressRoleValueHsrp string = "hsrp"

	// IPAddressRoleValueGlbp captures enum value "glbp"
	IPAddressRoleValueGlbp string = "glbp"

	// IPAddressRoleValueCarp captures enum value "carp"
	IPAddressRoleValueCarp string = "carp"
)

// prop value enum
func (m *IPAddressRole) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipAddressRoleTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPAddressRole) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("role"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("role"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAddressRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddressRole) UnmarshalBinary(b []byte) error {
	var res IPAddressRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPAddressStatus Status
// swagger:model IPAddressStatus
type IPAddressStatus struct {

	// label
	// Required: true
	// Enum: [Active Reserved Deprecated DHCP SLAAC]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [active reserved deprecated dhcp slaac]
	Value *string `json:"value"`
}

// Validate validates this IP address status
func (m *IPAddressStatus) Validate(formats strfmt.Registry) error {
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

var ipAddressStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Reserved","Deprecated","DHCP","SLAAC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressStatusTypeLabelPropEnum = append(ipAddressStatusTypeLabelPropEnum, v)
	}
}

const (

	// IPAddressStatusLabelActive captures enum value "Active"
	IPAddressStatusLabelActive string = "Active"

	// IPAddressStatusLabelReserved captures enum value "Reserved"
	IPAddressStatusLabelReserved string = "Reserved"

	// IPAddressStatusLabelDeprecated captures enum value "Deprecated"
	IPAddressStatusLabelDeprecated string = "Deprecated"

	// IPAddressStatusLabelDHCP captures enum value "DHCP"
	IPAddressStatusLabelDHCP string = "DHCP"

	// IPAddressStatusLabelSLAAC captures enum value "SLAAC"
	IPAddressStatusLabelSLAAC string = "SLAAC"
)

// prop value enum
func (m *IPAddressStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipAddressStatusTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPAddressStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var ipAddressStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","reserved","deprecated","dhcp","slaac"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressStatusTypeValuePropEnum = append(ipAddressStatusTypeValuePropEnum, v)
	}
}

const (

	// IPAddressStatusValueActive captures enum value "active"
	IPAddressStatusValueActive string = "active"

	// IPAddressStatusValueReserved captures enum value "reserved"
	IPAddressStatusValueReserved string = "reserved"

	// IPAddressStatusValueDeprecated captures enum value "deprecated"
	IPAddressStatusValueDeprecated string = "deprecated"

	// IPAddressStatusValueDhcp captures enum value "dhcp"
	IPAddressStatusValueDhcp string = "dhcp"

	// IPAddressStatusValueSlaac captures enum value "slaac"
	IPAddressStatusValueSlaac string = "slaac"
)

// prop value enum
func (m *IPAddressStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipAddressStatusTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPAddressStatus) validateValue(formats strfmt.Registry) error {

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
func (m *IPAddressStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddressStatus) UnmarshalBinary(b []byte) error {
	var res IPAddressStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
