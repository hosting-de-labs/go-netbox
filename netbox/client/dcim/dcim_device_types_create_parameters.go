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

	"github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewDcimDeviceTypesCreateParams creates a new DcimDeviceTypesCreateParams object
// with the default values initialized.
func NewDcimDeviceTypesCreateParams() *DcimDeviceTypesCreateParams {
	var ()
	return &DcimDeviceTypesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesCreateParamsWithTimeout creates a new DcimDeviceTypesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceTypesCreateParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesCreateParams {
	var ()
	return &DcimDeviceTypesCreateParams{

		timeout: timeout,
	}
}

// NewDcimDeviceTypesCreateParamsWithContext creates a new DcimDeviceTypesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceTypesCreateParamsWithContext(ctx context.Context) *DcimDeviceTypesCreateParams {
	var ()
	return &DcimDeviceTypesCreateParams{

		Context: ctx,
	}
}

// NewDcimDeviceTypesCreateParamsWithHTTPClient creates a new DcimDeviceTypesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceTypesCreateParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesCreateParams {
	var ()
	return &DcimDeviceTypesCreateParams{
		HTTPClient: client,
	}
}

/*DcimDeviceTypesCreateParams contains all the parameters to send to the API endpoint
for the dcim device types create operation typically these are written to a http.Request
*/
type DcimDeviceTypesCreateParams struct {

	/*Data*/
	Data *models.WritableDeviceType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device types create params
func (o *DcimDeviceTypesCreateParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types create params
func (o *DcimDeviceTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types create params
func (o *DcimDeviceTypesCreateParams) WithContext(ctx context.Context) *DcimDeviceTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types create params
func (o *DcimDeviceTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types create params
func (o *DcimDeviceTypesCreateParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types create params
func (o *DcimDeviceTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device types create params
func (o *DcimDeviceTypesCreateParams) WithData(data *models.WritableDeviceType) *DcimDeviceTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device types create params
func (o *DcimDeviceTypesCreateParams) SetData(data *models.WritableDeviceType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
