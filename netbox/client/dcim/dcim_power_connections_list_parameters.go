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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDcimPowerConnectionsListParams creates a new DcimPowerConnectionsListParams object
// with the default values initialized.
func NewDcimPowerConnectionsListParams() *DcimPowerConnectionsListParams {
	var ()
	return &DcimPowerConnectionsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerConnectionsListParamsWithTimeout creates a new DcimPowerConnectionsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerConnectionsListParamsWithTimeout(timeout time.Duration) *DcimPowerConnectionsListParams {
	var ()
	return &DcimPowerConnectionsListParams{

		timeout: timeout,
	}
}

// NewDcimPowerConnectionsListParamsWithContext creates a new DcimPowerConnectionsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerConnectionsListParamsWithContext(ctx context.Context) *DcimPowerConnectionsListParams {
	var ()
	return &DcimPowerConnectionsListParams{

		Context: ctx,
	}
}

// NewDcimPowerConnectionsListParamsWithHTTPClient creates a new DcimPowerConnectionsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerConnectionsListParamsWithHTTPClient(client *http.Client) *DcimPowerConnectionsListParams {
	var ()
	return &DcimPowerConnectionsListParams{
		HTTPClient: client,
	}
}

/*DcimPowerConnectionsListParams contains all the parameters to send to the API endpoint
for the dcim power connections list operation typically these are written to a http.Request
*/
type DcimPowerConnectionsListParams struct {

	/*ConnectionStatus*/
	ConnectionStatus *string
	/*ConnectionStatusn*/
	ConnectionStatusn *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *int64
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*NameIc*/
	NameIc *string
	/*NameIe*/
	NameIe *string
	/*NameIew*/
	NameIew *string
	/*NameIsw*/
	NameIsw *string
	/*Namen*/
	Namen *string
	/*NameNic*/
	NameNic *string
	/*NameNie*/
	NameNie *string
	/*NameNiew*/
	NameNiew *string
	/*NameNisw*/
	NameNisw *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Site*/
	Site *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithTimeout(timeout time.Duration) *DcimPowerConnectionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithContext(ctx context.Context) *DcimPowerConnectionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithHTTPClient(client *http.Client) *DcimPowerConnectionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionStatus adds the connectionStatus to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithConnectionStatus(connectionStatus *string) *DcimPowerConnectionsListParams {
	o.SetConnectionStatus(connectionStatus)
	return o
}

// SetConnectionStatus adds the connectionStatus to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetConnectionStatus(connectionStatus *string) {
	o.ConnectionStatus = connectionStatus
}

// WithConnectionStatusn adds the connectionStatusn to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithConnectionStatusn(connectionStatusn *string) *DcimPowerConnectionsListParams {
	o.SetConnectionStatusn(connectionStatusn)
	return o
}

// SetConnectionStatusn adds the connectionStatusN to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetConnectionStatusn(connectionStatusn *string) {
	o.ConnectionStatusn = connectionStatusn
}

// WithDevice adds the device to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithDevice(device *string) *DcimPowerConnectionsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithDeviceID(deviceID *int64) *DcimPowerConnectionsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithLimit adds the limit to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithLimit(limit *int64) *DcimPowerConnectionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithName(name *string) *DcimPowerConnectionsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithNameIc(nameIc *string) *DcimPowerConnectionsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithNameIe(nameIe *string) *DcimPowerConnectionsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithNameIew(nameIew *string) *DcimPowerConnectionsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithNameIsw(nameIsw *string) *DcimPowerConnectionsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithNamen(namen *string) *DcimPowerConnectionsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithNameNic(nameNic *string) *DcimPowerConnectionsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithNameNie(nameNie *string) *DcimPowerConnectionsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithNameNiew(nameNiew *string) *DcimPowerConnectionsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithNameNisw(nameNisw *string) *DcimPowerConnectionsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithOffset(offset *int64) *DcimPowerConnectionsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSite adds the site to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithSite(site *string) *DcimPowerConnectionsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetSite(site *string) {
	o.Site = site
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerConnectionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConnectionStatus != nil {

		// query param connection_status
		var qrConnectionStatus string
		if o.ConnectionStatus != nil {
			qrConnectionStatus = *o.ConnectionStatus
		}
		qConnectionStatus := qrConnectionStatus
		if qConnectionStatus != "" {
			if err := r.SetQueryParam("connection_status", qConnectionStatus); err != nil {
				return err
			}
		}

	}

	if o.ConnectionStatusn != nil {

		// query param connection_status__n
		var qrConnectionStatusn string
		if o.ConnectionStatusn != nil {
			qrConnectionStatusn = *o.ConnectionStatusn
		}
		qConnectionStatusn := qrConnectionStatusn
		if qConnectionStatusn != "" {
			if err := r.SetQueryParam("connection_status__n", qConnectionStatusn); err != nil {
				return err
			}
		}

	}

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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string
		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {
			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}

	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string
		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {
			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}

	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string
		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {
			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}

	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string
		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {
			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}

	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string
		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {
			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}

	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string
		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {
			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}

	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string
		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {
			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}

	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string
		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {
			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}

	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string
		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {
			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
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

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
