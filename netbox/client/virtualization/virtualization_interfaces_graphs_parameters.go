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

package virtualization

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

// NewVirtualizationInterfacesGraphsParams creates a new VirtualizationInterfacesGraphsParams object
// with the default values initialized.
func NewVirtualizationInterfacesGraphsParams() *VirtualizationInterfacesGraphsParams {
	var ()
	return &VirtualizationInterfacesGraphsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesGraphsParamsWithTimeout creates a new VirtualizationInterfacesGraphsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationInterfacesGraphsParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesGraphsParams {
	var ()
	return &VirtualizationInterfacesGraphsParams{

		timeout: timeout,
	}
}

// NewVirtualizationInterfacesGraphsParamsWithContext creates a new VirtualizationInterfacesGraphsParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationInterfacesGraphsParamsWithContext(ctx context.Context) *VirtualizationInterfacesGraphsParams {
	var ()
	return &VirtualizationInterfacesGraphsParams{

		Context: ctx,
	}
}

// NewVirtualizationInterfacesGraphsParamsWithHTTPClient creates a new VirtualizationInterfacesGraphsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationInterfacesGraphsParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesGraphsParams {
	var ()
	return &VirtualizationInterfacesGraphsParams{
		HTTPClient: client,
	}
}

/*VirtualizationInterfacesGraphsParams contains all the parameters to send to the API endpoint
for the virtualization interfaces graphs operation typically these are written to a http.Request
*/
type VirtualizationInterfacesGraphsParams struct {

	/*ID
	  A unique integer value identifying this interface.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization interfaces graphs params
func (o *VirtualizationInterfacesGraphsParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesGraphsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces graphs params
func (o *VirtualizationInterfacesGraphsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces graphs params
func (o *VirtualizationInterfacesGraphsParams) WithContext(ctx context.Context) *VirtualizationInterfacesGraphsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces graphs params
func (o *VirtualizationInterfacesGraphsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces graphs params
func (o *VirtualizationInterfacesGraphsParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesGraphsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces graphs params
func (o *VirtualizationInterfacesGraphsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization interfaces graphs params
func (o *VirtualizationInterfacesGraphsParams) WithID(id int64) *VirtualizationInterfacesGraphsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization interfaces graphs params
func (o *VirtualizationInterfacesGraphsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesGraphsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
