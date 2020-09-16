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

// WritableInterface writable interface
// swagger:model WritableInterface
type WritableInterface struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Connection status
	// Enum: [false true]
	ConnectionStatus *bool `json:"connection_status,omitempty"`

	// Count ipaddresses
	// Read Only: true
	CountIpaddresses int64 `json:"count_ipaddresses,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Parent LAG
	Lag *int64 `json:"lag,omitempty"`

	// MAC Address
	MacAddress *string `json:"mac_address,omitempty"`

	// OOB Management
	//
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// Mode
	// Enum: [access tagged tagged-all]
	Mode string `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tagged vlans
	// Unique: true
	TaggedVlans []int64 `json:"tagged_vlans"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Type
	// Required: true
	// Enum: [virtual lag 100base-tx 1000base-t 2.5gbase-t 5gbase-t 10gbase-t 10gbase-cx4 1000base-x-gbic 1000base-x-sfp 10gbase-x-sfpp 10gbase-x-xfp 10gbase-x-xenpak 10gbase-x-x2 25gbase-x-sfp28 40gbase-x-qsfpp 50gbase-x-sfp28 100gbase-x-cfp 100gbase-x-cfp2 200gbase-x-cfp2 100gbase-x-cfp4 100gbase-x-cpak 100gbase-x-qsfp28 200gbase-x-qsfp56 400gbase-x-qsfpdd 400gbase-x-osfp ieee802.11a ieee802.11g ieee802.11n ieee802.11ac ieee802.11ad ieee802.11ax gsm cdma lte sonet-oc3 sonet-oc12 sonet-oc48 sonet-oc192 sonet-oc768 sonet-oc1920 sonet-oc3840 1gfc-sfp 2gfc-sfp 4gfc-sfp 8gfc-sfpp 16gfc-sfpp 32gfc-sfp28 128gfc-sfp28 infiniband-sdr infiniband-ddr infiniband-qdr infiniband-fdr10 infiniband-fdr infiniband-edr infiniband-hdr infiniband-ndr infiniband-xdr t1 e1 t3 e3 cisco-stackwise cisco-stackwise-plus cisco-flexstack cisco-flexstack-plus juniper-vcp extreme-summitstack extreme-summitstack-128 extreme-summitstack-256 extreme-summitstack-512 other]
	Type *string `json:"type"`

	// Untagged VLAN
	UntaggedVlan *int64 `json:"untagged_vlan,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable interface
func (m *WritableInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *WritableInterface) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

var writableInterfaceTypeConnectionStatusPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTypeConnectionStatusPropEnum = append(writableInterfaceTypeConnectionStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableInterface) validateConnectionStatusEnum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, writableInterfaceTypeConnectionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterface) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("connection_status", "body", *m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", string(m.Label), 64); err != nil {
		return err
	}

	return nil
}

var writableInterfaceTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access","tagged","tagged-all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTypeModePropEnum = append(writableInterfaceTypeModePropEnum, v)
	}
}

const (

	// WritableInterfaceModeAccess captures enum value "access"
	WritableInterfaceModeAccess string = "access"

	// WritableInterfaceModeTagged captures enum value "tagged"
	WritableInterfaceModeTagged string = "tagged"

	// WritableInterfaceModeTaggedAll captures enum value "tagged-all"
	WritableInterfaceModeTaggedAll string = "tagged-all"
)

// prop value enum
func (m *WritableInterface) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableInterfaceTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterface) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateMtu(formats strfmt.Registry) error {

	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", int64(*m.Mtu), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", int64(*m.Mtu), 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateTaggedVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateTags(formats strfmt.Registry) error {

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

var writableInterfaceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","lag","100base-tx","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","128gfc-sfp28","infiniband-sdr","infiniband-ddr","infiniband-qdr","infiniband-fdr10","infiniband-fdr","infiniband-edr","infiniband-hdr","infiniband-ndr","infiniband-xdr","t1","e1","t3","e3","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTypeTypePropEnum = append(writableInterfaceTypeTypePropEnum, v)
	}
}

const (

	// WritableInterfaceTypeVirtual captures enum value "virtual"
	WritableInterfaceTypeVirtual string = "virtual"

	// WritableInterfaceTypeLag captures enum value "lag"
	WritableInterfaceTypeLag string = "lag"

	// WritableInterfaceTypeNr100baseTx captures enum value "100base-tx"
	WritableInterfaceTypeNr100baseTx string = "100base-tx"

	// WritableInterfaceTypeNr1000baset captures enum value "1000base-t"
	WritableInterfaceTypeNr1000baset string = "1000base-t"

	// WritableInterfaceTypeNr25gbaset captures enum value "2.5gbase-t"
	WritableInterfaceTypeNr25gbaset string = "2.5gbase-t"

	// WritableInterfaceTypeNr5gbaset captures enum value "5gbase-t"
	WritableInterfaceTypeNr5gbaset string = "5gbase-t"

	// WritableInterfaceTypeNr10gbaset captures enum value "10gbase-t"
	WritableInterfaceTypeNr10gbaset string = "10gbase-t"

	// WritableInterfaceTypeNr10gbaseCx4 captures enum value "10gbase-cx4"
	WritableInterfaceTypeNr10gbaseCx4 string = "10gbase-cx4"

	// WritableInterfaceTypeNr1000basexGbic captures enum value "1000base-x-gbic"
	WritableInterfaceTypeNr1000basexGbic string = "1000base-x-gbic"

	// WritableInterfaceTypeNr1000basexSfp captures enum value "1000base-x-sfp"
	WritableInterfaceTypeNr1000basexSfp string = "1000base-x-sfp"

	// WritableInterfaceTypeNr10gbasexSfpp captures enum value "10gbase-x-sfpp"
	WritableInterfaceTypeNr10gbasexSfpp string = "10gbase-x-sfpp"

	// WritableInterfaceTypeNr10gbasexXfp captures enum value "10gbase-x-xfp"
	WritableInterfaceTypeNr10gbasexXfp string = "10gbase-x-xfp"

	// WritableInterfaceTypeNr10gbasexXenpak captures enum value "10gbase-x-xenpak"
	WritableInterfaceTypeNr10gbasexXenpak string = "10gbase-x-xenpak"

	// WritableInterfaceTypeNr10gbasexX2 captures enum value "10gbase-x-x2"
	WritableInterfaceTypeNr10gbasexX2 string = "10gbase-x-x2"

	// WritableInterfaceTypeNr25gbasexSfp28 captures enum value "25gbase-x-sfp28"
	WritableInterfaceTypeNr25gbasexSfp28 string = "25gbase-x-sfp28"

	// WritableInterfaceTypeNr40gbasexQsfpp captures enum value "40gbase-x-qsfpp"
	WritableInterfaceTypeNr40gbasexQsfpp string = "40gbase-x-qsfpp"

	// WritableInterfaceTypeNr50gbasexSfp28 captures enum value "50gbase-x-sfp28"
	WritableInterfaceTypeNr50gbasexSfp28 string = "50gbase-x-sfp28"

	// WritableInterfaceTypeNr100gbasexCfp captures enum value "100gbase-x-cfp"
	WritableInterfaceTypeNr100gbasexCfp string = "100gbase-x-cfp"

	// WritableInterfaceTypeNr100gbasexCfp2 captures enum value "100gbase-x-cfp2"
	WritableInterfaceTypeNr100gbasexCfp2 string = "100gbase-x-cfp2"

	// WritableInterfaceTypeNr200gbasexCfp2 captures enum value "200gbase-x-cfp2"
	WritableInterfaceTypeNr200gbasexCfp2 string = "200gbase-x-cfp2"

	// WritableInterfaceTypeNr100gbasexCfp4 captures enum value "100gbase-x-cfp4"
	WritableInterfaceTypeNr100gbasexCfp4 string = "100gbase-x-cfp4"

	// WritableInterfaceTypeNr100gbasexCpak captures enum value "100gbase-x-cpak"
	WritableInterfaceTypeNr100gbasexCpak string = "100gbase-x-cpak"

	// WritableInterfaceTypeNr100gbasexQsfp28 captures enum value "100gbase-x-qsfp28"
	WritableInterfaceTypeNr100gbasexQsfp28 string = "100gbase-x-qsfp28"

	// WritableInterfaceTypeNr200gbasexQsfp56 captures enum value "200gbase-x-qsfp56"
	WritableInterfaceTypeNr200gbasexQsfp56 string = "200gbase-x-qsfp56"

	// WritableInterfaceTypeNr400gbasexQsfpdd captures enum value "400gbase-x-qsfpdd"
	WritableInterfaceTypeNr400gbasexQsfpdd string = "400gbase-x-qsfpdd"

	// WritableInterfaceTypeNr400gbasexOsfp captures enum value "400gbase-x-osfp"
	WritableInterfaceTypeNr400gbasexOsfp string = "400gbase-x-osfp"

	// WritableInterfaceTypeIeee80211a captures enum value "ieee802.11a"
	WritableInterfaceTypeIeee80211a string = "ieee802.11a"

	// WritableInterfaceTypeIeee80211g captures enum value "ieee802.11g"
	WritableInterfaceTypeIeee80211g string = "ieee802.11g"

	// WritableInterfaceTypeIeee80211n captures enum value "ieee802.11n"
	WritableInterfaceTypeIeee80211n string = "ieee802.11n"

	// WritableInterfaceTypeIeee80211ac captures enum value "ieee802.11ac"
	WritableInterfaceTypeIeee80211ac string = "ieee802.11ac"

	// WritableInterfaceTypeIeee80211ad captures enum value "ieee802.11ad"
	WritableInterfaceTypeIeee80211ad string = "ieee802.11ad"

	// WritableInterfaceTypeIeee80211ax captures enum value "ieee802.11ax"
	WritableInterfaceTypeIeee80211ax string = "ieee802.11ax"

	// WritableInterfaceTypeGsm captures enum value "gsm"
	WritableInterfaceTypeGsm string = "gsm"

	// WritableInterfaceTypeCdma captures enum value "cdma"
	WritableInterfaceTypeCdma string = "cdma"

	// WritableInterfaceTypeLte captures enum value "lte"
	WritableInterfaceTypeLte string = "lte"

	// WritableInterfaceTypeSonetOc3 captures enum value "sonet-oc3"
	WritableInterfaceTypeSonetOc3 string = "sonet-oc3"

	// WritableInterfaceTypeSonetOc12 captures enum value "sonet-oc12"
	WritableInterfaceTypeSonetOc12 string = "sonet-oc12"

	// WritableInterfaceTypeSonetOc48 captures enum value "sonet-oc48"
	WritableInterfaceTypeSonetOc48 string = "sonet-oc48"

	// WritableInterfaceTypeSonetOc192 captures enum value "sonet-oc192"
	WritableInterfaceTypeSonetOc192 string = "sonet-oc192"

	// WritableInterfaceTypeSonetOc768 captures enum value "sonet-oc768"
	WritableInterfaceTypeSonetOc768 string = "sonet-oc768"

	// WritableInterfaceTypeSonetOc1920 captures enum value "sonet-oc1920"
	WritableInterfaceTypeSonetOc1920 string = "sonet-oc1920"

	// WritableInterfaceTypeSonetOc3840 captures enum value "sonet-oc3840"
	WritableInterfaceTypeSonetOc3840 string = "sonet-oc3840"

	// WritableInterfaceTypeNr1gfcSfp captures enum value "1gfc-sfp"
	WritableInterfaceTypeNr1gfcSfp string = "1gfc-sfp"

	// WritableInterfaceTypeNr2gfcSfp captures enum value "2gfc-sfp"
	WritableInterfaceTypeNr2gfcSfp string = "2gfc-sfp"

	// WritableInterfaceTypeNr4gfcSfp captures enum value "4gfc-sfp"
	WritableInterfaceTypeNr4gfcSfp string = "4gfc-sfp"

	// WritableInterfaceTypeNr8gfcSfpp captures enum value "8gfc-sfpp"
	WritableInterfaceTypeNr8gfcSfpp string = "8gfc-sfpp"

	// WritableInterfaceTypeNr16gfcSfpp captures enum value "16gfc-sfpp"
	WritableInterfaceTypeNr16gfcSfpp string = "16gfc-sfpp"

	// WritableInterfaceTypeNr32gfcSfp28 captures enum value "32gfc-sfp28"
	WritableInterfaceTypeNr32gfcSfp28 string = "32gfc-sfp28"

	// WritableInterfaceTypeNr128gfcSfp28 captures enum value "128gfc-sfp28"
	WritableInterfaceTypeNr128gfcSfp28 string = "128gfc-sfp28"

	// WritableInterfaceTypeInfinibandSdr captures enum value "infiniband-sdr"
	WritableInterfaceTypeInfinibandSdr string = "infiniband-sdr"

	// WritableInterfaceTypeInfinibandDdr captures enum value "infiniband-ddr"
	WritableInterfaceTypeInfinibandDdr string = "infiniband-ddr"

	// WritableInterfaceTypeInfinibandQdr captures enum value "infiniband-qdr"
	WritableInterfaceTypeInfinibandQdr string = "infiniband-qdr"

	// WritableInterfaceTypeInfinibandFdr10 captures enum value "infiniband-fdr10"
	WritableInterfaceTypeInfinibandFdr10 string = "infiniband-fdr10"

	// WritableInterfaceTypeInfinibandFdr captures enum value "infiniband-fdr"
	WritableInterfaceTypeInfinibandFdr string = "infiniband-fdr"

	// WritableInterfaceTypeInfinibandEdr captures enum value "infiniband-edr"
	WritableInterfaceTypeInfinibandEdr string = "infiniband-edr"

	// WritableInterfaceTypeInfinibandHdr captures enum value "infiniband-hdr"
	WritableInterfaceTypeInfinibandHdr string = "infiniband-hdr"

	// WritableInterfaceTypeInfinibandNdr captures enum value "infiniband-ndr"
	WritableInterfaceTypeInfinibandNdr string = "infiniband-ndr"

	// WritableInterfaceTypeInfinibandXdr captures enum value "infiniband-xdr"
	WritableInterfaceTypeInfinibandXdr string = "infiniband-xdr"

	// WritableInterfaceTypeT1 captures enum value "t1"
	WritableInterfaceTypeT1 string = "t1"

	// WritableInterfaceTypeE1 captures enum value "e1"
	WritableInterfaceTypeE1 string = "e1"

	// WritableInterfaceTypeT3 captures enum value "t3"
	WritableInterfaceTypeT3 string = "t3"

	// WritableInterfaceTypeE3 captures enum value "e3"
	WritableInterfaceTypeE3 string = "e3"

	// WritableInterfaceTypeCiscoStackwise captures enum value "cisco-stackwise"
	WritableInterfaceTypeCiscoStackwise string = "cisco-stackwise"

	// WritableInterfaceTypeCiscoStackwisePlus captures enum value "cisco-stackwise-plus"
	WritableInterfaceTypeCiscoStackwisePlus string = "cisco-stackwise-plus"

	// WritableInterfaceTypeCiscoFlexstack captures enum value "cisco-flexstack"
	WritableInterfaceTypeCiscoFlexstack string = "cisco-flexstack"

	// WritableInterfaceTypeCiscoFlexstackPlus captures enum value "cisco-flexstack-plus"
	WritableInterfaceTypeCiscoFlexstackPlus string = "cisco-flexstack-plus"

	// WritableInterfaceTypeJuniperVcp captures enum value "juniper-vcp"
	WritableInterfaceTypeJuniperVcp string = "juniper-vcp"

	// WritableInterfaceTypeExtremeSummitstack captures enum value "extreme-summitstack"
	WritableInterfaceTypeExtremeSummitstack string = "extreme-summitstack"

	// WritableInterfaceTypeExtremeSummitstack128 captures enum value "extreme-summitstack-128"
	WritableInterfaceTypeExtremeSummitstack128 string = "extreme-summitstack-128"

	// WritableInterfaceTypeExtremeSummitstack256 captures enum value "extreme-summitstack-256"
	WritableInterfaceTypeExtremeSummitstack256 string = "extreme-summitstack-256"

	// WritableInterfaceTypeExtremeSummitstack512 captures enum value "extreme-summitstack-512"
	WritableInterfaceTypeExtremeSummitstack512 string = "extreme-summitstack-512"

	// WritableInterfaceTypeOther captures enum value "other"
	WritableInterfaceTypeOther string = "other"
)

// prop value enum
func (m *WritableInterface) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableInterfaceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterface) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableInterface) UnmarshalBinary(b []byte) error {
	var res WritableInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
