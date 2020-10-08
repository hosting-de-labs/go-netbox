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

// NewDcimPowerOutletsTraceParams creates a new DcimPowerOutletsTraceParams object
// with the default values initialized.
func NewDcimPowerOutletsTraceParams() *DcimPowerOutletsTraceParams {
	var ()
	return &DcimPowerOutletsTraceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsTraceParamsWithTimeout creates a new DcimPowerOutletsTraceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerOutletsTraceParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsTraceParams {
	var ()
	return &DcimPowerOutletsTraceParams{

		timeout: timeout,
	}
}

// NewDcimPowerOutletsTraceParamsWithContext creates a new DcimPowerOutletsTraceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerOutletsTraceParamsWithContext(ctx context.Context) *DcimPowerOutletsTraceParams {
	var ()
	return &DcimPowerOutletsTraceParams{

		Context: ctx,
	}
}

// NewDcimPowerOutletsTraceParamsWithHTTPClient creates a new DcimPowerOutletsTraceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerOutletsTraceParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsTraceParams {
	var ()
	return &DcimPowerOutletsTraceParams{
		HTTPClient: client,
	}
}

/*DcimPowerOutletsTraceParams contains all the parameters to send to the API endpoint
for the dcim power outlets trace operation typically these are written to a http.Request
*/
type DcimPowerOutletsTraceParams struct {

	/*ID
	  A unique integer value identifying this power outlet.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsTraceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) WithContext(ctx context.Context) *DcimPowerOutletsTraceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsTraceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) WithID(id int64) *DcimPowerOutletsTraceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsTraceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
