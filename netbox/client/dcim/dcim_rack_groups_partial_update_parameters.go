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

	"github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewDcimRackGroupsPartialUpdateParams creates a new DcimRackGroupsPartialUpdateParams object
// with the default values initialized.
func NewDcimRackGroupsPartialUpdateParams() *DcimRackGroupsPartialUpdateParams {
	var ()
	return &DcimRackGroupsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsPartialUpdateParamsWithTimeout creates a new DcimRackGroupsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackGroupsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRackGroupsPartialUpdateParams {
	var ()
	return &DcimRackGroupsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimRackGroupsPartialUpdateParamsWithContext creates a new DcimRackGroupsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackGroupsPartialUpdateParamsWithContext(ctx context.Context) *DcimRackGroupsPartialUpdateParams {
	var ()
	return &DcimRackGroupsPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimRackGroupsPartialUpdateParamsWithHTTPClient creates a new DcimRackGroupsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackGroupsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRackGroupsPartialUpdateParams {
	var ()
	return &DcimRackGroupsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimRackGroupsPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim rack groups partial update operation typically these are written to a http.Request
*/
type DcimRackGroupsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableRackGroup
	/*ID
	  A unique integer value identifying this rack group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRackGroupsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithContext(ctx context.Context) *DcimRackGroupsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRackGroupsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithData(data *models.WritableRackGroup) *DcimRackGroupsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetData(data *models.WritableRackGroup) {
	o.Data = data
}

// WithID adds the id to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithID(id int64) *DcimRackGroupsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
