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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIPAMIPAddressesListParams creates a new IPAMIPAddressesListParams object
// with the default values initialized.
func NewIPAMIPAddressesListParams() *IPAMIPAddressesListParams {
	var ()
	return &IPAMIPAddressesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMIPAddressesListParamsWithTimeout creates a new IPAMIPAddressesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMIPAddressesListParamsWithTimeout(timeout time.Duration) *IPAMIPAddressesListParams {
	var ()
	return &IPAMIPAddressesListParams{

		timeout: timeout,
	}
}

// NewIPAMIPAddressesListParamsWithContext creates a new IPAMIPAddressesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMIPAddressesListParamsWithContext(ctx context.Context) *IPAMIPAddressesListParams {
	var ()
	return &IPAMIPAddressesListParams{

		Context: ctx,
	}
}

// NewIPAMIPAddressesListParamsWithHTTPClient creates a new IPAMIPAddressesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMIPAddressesListParamsWithHTTPClient(client *http.Client) *IPAMIPAddressesListParams {
	var ()
	return &IPAMIPAddressesListParams{
		HTTPClient: client,
	}
}

/*IPAMIPAddressesListParams contains all the parameters to send to the API endpoint
for the ipam ip addresses list operation typically these are written to a http.Request
*/
type IPAMIPAddressesListParams struct {

	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *int64
	/*Family*/
	Family *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*InterfaceID*/
	InterfaceID *int64
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MaskLength*/
	MaskLength *int64
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Parent*/
	Parent *string
	/*Q*/
	Q *string
	/*Role*/
	Role *string
	/*Status*/
	Status *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *int64
	/*VirtualMachine*/
	VirtualMachine *string
	/*VirtualMachineID*/
	VirtualMachineID *int64
	/*Vrf*/
	Vrf *string
	/*VrfID*/
	VrfID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithTimeout(timeout time.Duration) *IPAMIPAddressesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithContext(ctx context.Context) *IPAMIPAddressesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithHTTPClient(client *http.Client) *IPAMIPAddressesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithDevice(device *string) *IPAMIPAddressesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithDeviceID(deviceID *int64) *IPAMIPAddressesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithFamily adds the family to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithFamily(family *string) *IPAMIPAddressesListParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetFamily(family *string) {
	o.Family = family
}

// WithIDIn adds the iDIn to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithIDIn(iDIn *string) *IPAMIPAddressesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithInterfaceID adds the interfaceID to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithInterfaceID(interfaceID *int64) *IPAMIPAddressesListParams {
	o.SetInterfaceID(interfaceID)
	return o
}

// SetInterfaceID adds the interfaceId to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetInterfaceID(interfaceID *int64) {
	o.InterfaceID = interfaceID
}

// WithLimit adds the limit to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithLimit(limit *int64) *IPAMIPAddressesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaskLength adds the maskLength to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithMaskLength(maskLength *int64) *IPAMIPAddressesListParams {
	o.SetMaskLength(maskLength)
	return o
}

// SetMaskLength adds the maskLength to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetMaskLength(maskLength *int64) {
	o.MaskLength = maskLength
}

// WithOffset adds the offset to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithOffset(offset *int64) *IPAMIPAddressesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithParent adds the parent to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithParent(parent *string) *IPAMIPAddressesListParams {
	o.SetParent(parent)
	return o
}

// SetParent adds the parent to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetParent(parent *string) {
	o.Parent = parent
}

// WithQ adds the q to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithQ(q *string) *IPAMIPAddressesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithRole(role *string) *IPAMIPAddressesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetRole(role *string) {
	o.Role = role
}

// WithStatus adds the status to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithStatus(status *string) *IPAMIPAddressesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTenant adds the tenant to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithTenant(tenant *string) *IPAMIPAddressesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithTenantID(tenantID *int64) *IPAMIPAddressesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetTenantID(tenantID *int64) {
	o.TenantID = tenantID
}

// WithVirtualMachine adds the virtualMachine to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithVirtualMachine(virtualMachine *string) *IPAMIPAddressesListParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachineID adds the virtualMachineID to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithVirtualMachineID(virtualMachineID *int64) *IPAMIPAddressesListParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetVirtualMachineID(virtualMachineID *int64) {
	o.VirtualMachineID = virtualMachineID
}

// WithVrf adds the vrf to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithVrf(vrf *string) *IPAMIPAddressesListParams {
	o.SetVrf(vrf)
	return o
}

// SetVrf adds the vrf to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetVrf(vrf *string) {
	o.Vrf = vrf
}

// WithVrfID adds the vrfID to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) WithVrfID(vrfID *int64) *IPAMIPAddressesListParams {
	o.SetVrfID(vrfID)
	return o
}

// SetVrfID adds the vrfId to the ipam ip addresses list params
func (o *IPAMIPAddressesListParams) SetVrfID(vrfID *int64) {
	o.VrfID = vrfID
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMIPAddressesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID int64
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := swag.FormatInt64(qrDeviceID)
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.Family != nil {

		// query param family
		var qrFamily string
		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := qrFamily
		if qFamily != "" {
			if err := r.SetQueryParam("family", qFamily); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn string
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := qrIDIn
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

	if o.InterfaceID != nil {

		// query param interface_id
		var qrInterfaceID int64
		if o.InterfaceID != nil {
			qrInterfaceID = *o.InterfaceID
		}
		qInterfaceID := swag.FormatInt64(qrInterfaceID)
		if qInterfaceID != "" {
			if err := r.SetQueryParam("interface_id", qInterfaceID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.MaskLength != nil {

		// query param mask_length
		var qrMaskLength int64
		if o.MaskLength != nil {
			qrMaskLength = *o.MaskLength
		}
		qMaskLength := swag.FormatInt64(qrMaskLength)
		if qMaskLength != "" {
			if err := r.SetQueryParam("mask_length", qMaskLength); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Parent != nil {

		// query param parent
		var qrParent string
		if o.Parent != nil {
			qrParent = *o.Parent
		}
		qParent := qrParent
		if qParent != "" {
			if err := r.SetQueryParam("parent", qParent); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID int64
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := swag.FormatInt64(qrTenantID)
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}

	}

	if o.VirtualMachine != nil {

		// query param virtual_machine
		var qrVirtualMachine string
		if o.VirtualMachine != nil {
			qrVirtualMachine = *o.VirtualMachine
		}
		qVirtualMachine := qrVirtualMachine
		if qVirtualMachine != "" {
			if err := r.SetQueryParam("virtual_machine", qVirtualMachine); err != nil {
				return err
			}
		}

	}

	if o.VirtualMachineID != nil {

		// query param virtual_machine_id
		var qrVirtualMachineID int64
		if o.VirtualMachineID != nil {
			qrVirtualMachineID = *o.VirtualMachineID
		}
		qVirtualMachineID := swag.FormatInt64(qrVirtualMachineID)
		if qVirtualMachineID != "" {
			if err := r.SetQueryParam("virtual_machine_id", qVirtualMachineID); err != nil {
				return err
			}
		}

	}

	if o.Vrf != nil {

		// query param vrf
		var qrVrf string
		if o.Vrf != nil {
			qrVrf = *o.Vrf
		}
		qVrf := qrVrf
		if qVrf != "" {
			if err := r.SetQueryParam("vrf", qVrf); err != nil {
				return err
			}
		}

	}

	if o.VrfID != nil {

		// query param vrf_id
		var qrVrfID int64
		if o.VrfID != nil {
			qrVrfID = *o.VrfID
		}
		qVrfID := swag.FormatInt64(qrVrfID)
		if qVrfID != "" {
			if err := r.SetQueryParam("vrf_id", qVrfID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
