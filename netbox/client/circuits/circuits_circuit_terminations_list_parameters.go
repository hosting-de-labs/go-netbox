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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCircuitsCircuitTerminationsListParams creates a new CircuitsCircuitTerminationsListParams object
// with the default values initialized.
func NewCircuitsCircuitTerminationsListParams() *CircuitsCircuitTerminationsListParams {
	var ()
	return &CircuitsCircuitTerminationsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTerminationsListParamsWithTimeout creates a new CircuitsCircuitTerminationsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitTerminationsListParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsListParams {
	var ()
	return &CircuitsCircuitTerminationsListParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitTerminationsListParamsWithContext creates a new CircuitsCircuitTerminationsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitTerminationsListParamsWithContext(ctx context.Context) *CircuitsCircuitTerminationsListParams {
	var ()
	return &CircuitsCircuitTerminationsListParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitTerminationsListParamsWithHTTPClient creates a new CircuitsCircuitTerminationsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitTerminationsListParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsListParams {
	var ()
	return &CircuitsCircuitTerminationsListParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitTerminationsListParams contains all the parameters to send to the API endpoint
for the circuits circuit terminations list operation typically these are written to a http.Request
*/
type CircuitsCircuitTerminationsListParams struct {

	/*CircuitID*/
	CircuitID *int64
	/*CircuitIDn*/
	CircuitIDn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*PortSpeed*/
	PortSpeed *float64
	/*PortSpeedGt*/
	PortSpeedGt *string
	/*PortSpeedGte*/
	PortSpeedGte *string
	/*PortSpeedLt*/
	PortSpeedLt *string
	/*PortSpeedLte*/
	PortSpeedLte *string
	/*PortSpeedn*/
	PortSpeedn *string
	/*Q*/
	Q *string
	/*Site*/
	Site *string
	/*Siten*/
	Siten *string
	/*SiteID*/
	SiteID *int64
	/*SiteIDn*/
	SiteIDn *string
	/*TermSide*/
	TermSide *string
	/*TermSiden*/
	TermSiden *string
	/*UpstreamSpeed*/
	UpstreamSpeed *float64
	/*UpstreamSpeedGt*/
	UpstreamSpeedGt *string
	/*UpstreamSpeedGte*/
	UpstreamSpeedGte *string
	/*UpstreamSpeedLt*/
	UpstreamSpeedLt *string
	/*UpstreamSpeedLte*/
	UpstreamSpeedLte *string
	/*UpstreamSpeedn*/
	UpstreamSpeedn *string
	/*XconnectID*/
	XconnectID *int64
	/*XconnectIDIc*/
	XconnectIDIc *string
	/*XconnectIDIe*/
	XconnectIDIe *string
	/*XconnectIDIew*/
	XconnectIDIew *string
	/*XconnectIDIsw*/
	XconnectIDIsw *string
	/*XconnectIDn*/
	XconnectIDn *string
	/*XconnectIDNic*/
	XconnectIDNic *string
	/*XconnectIDNie*/
	XconnectIDNie *string
	/*XconnectIDNiew*/
	XconnectIDNiew *string
	/*XconnectIDNisw*/
	XconnectIDNisw *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithContext(ctx context.Context) *CircuitsCircuitTerminationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCircuitID adds the circuitID to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithCircuitID(circuitID *int64) *CircuitsCircuitTerminationsListParams {
	o.SetCircuitID(circuitID)
	return o
}

// SetCircuitID adds the circuitId to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetCircuitID(circuitID *int64) {
	o.CircuitID = circuitID
}

// WithCircuitIDn adds the circuitIDn to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithCircuitIDn(circuitIDn *string) *CircuitsCircuitTerminationsListParams {
	o.SetCircuitIDn(circuitIDn)
	return o
}

// SetCircuitIDn adds the circuitIdN to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetCircuitIDn(circuitIDn *string) {
	o.CircuitIDn = circuitIDn
}

// WithLimit adds the limit to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithLimit(limit *int64) *CircuitsCircuitTerminationsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithOffset(offset *int64) *CircuitsCircuitTerminationsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPortSpeed adds the portSpeed to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithPortSpeed(portSpeed *float64) *CircuitsCircuitTerminationsListParams {
	o.SetPortSpeed(portSpeed)
	return o
}

// SetPortSpeed adds the portSpeed to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetPortSpeed(portSpeed *float64) {
	o.PortSpeed = portSpeed
}

// WithPortSpeedGt adds the portSpeedGt to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithPortSpeedGt(portSpeedGt *string) *CircuitsCircuitTerminationsListParams {
	o.SetPortSpeedGt(portSpeedGt)
	return o
}

// SetPortSpeedGt adds the portSpeedGt to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetPortSpeedGt(portSpeedGt *string) {
	o.PortSpeedGt = portSpeedGt
}

// WithPortSpeedGte adds the portSpeedGte to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithPortSpeedGte(portSpeedGte *string) *CircuitsCircuitTerminationsListParams {
	o.SetPortSpeedGte(portSpeedGte)
	return o
}

// SetPortSpeedGte adds the portSpeedGte to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetPortSpeedGte(portSpeedGte *string) {
	o.PortSpeedGte = portSpeedGte
}

// WithPortSpeedLt adds the portSpeedLt to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithPortSpeedLt(portSpeedLt *string) *CircuitsCircuitTerminationsListParams {
	o.SetPortSpeedLt(portSpeedLt)
	return o
}

// SetPortSpeedLt adds the portSpeedLt to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetPortSpeedLt(portSpeedLt *string) {
	o.PortSpeedLt = portSpeedLt
}

// WithPortSpeedLte adds the portSpeedLte to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithPortSpeedLte(portSpeedLte *string) *CircuitsCircuitTerminationsListParams {
	o.SetPortSpeedLte(portSpeedLte)
	return o
}

// SetPortSpeedLte adds the portSpeedLte to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetPortSpeedLte(portSpeedLte *string) {
	o.PortSpeedLte = portSpeedLte
}

// WithPortSpeedn adds the portSpeedn to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithPortSpeedn(portSpeedn *string) *CircuitsCircuitTerminationsListParams {
	o.SetPortSpeedn(portSpeedn)
	return o
}

// SetPortSpeedn adds the portSpeedN to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetPortSpeedn(portSpeedn *string) {
	o.PortSpeedn = portSpeedn
}

// WithQ adds the q to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithQ(q *string) *CircuitsCircuitTerminationsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSite adds the site to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithSite(site *string) *CircuitsCircuitTerminationsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithSiten(siten *string) *CircuitsCircuitTerminationsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithSiteID(siteID *int64) *CircuitsCircuitTerminationsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetSiteID(siteID *int64) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithSiteIDn(siteIDn *string) *CircuitsCircuitTerminationsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTermSide adds the termSide to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithTermSide(termSide *string) *CircuitsCircuitTerminationsListParams {
	o.SetTermSide(termSide)
	return o
}

// SetTermSide adds the termSide to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetTermSide(termSide *string) {
	o.TermSide = termSide
}

// WithTermSiden adds the termSiden to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithTermSiden(termSiden *string) *CircuitsCircuitTerminationsListParams {
	o.SetTermSiden(termSiden)
	return o
}

// SetTermSiden adds the termSideN to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetTermSiden(termSiden *string) {
	o.TermSiden = termSiden
}

// WithUpstreamSpeed adds the upstreamSpeed to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithUpstreamSpeed(upstreamSpeed *float64) *CircuitsCircuitTerminationsListParams {
	o.SetUpstreamSpeed(upstreamSpeed)
	return o
}

// SetUpstreamSpeed adds the upstreamSpeed to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetUpstreamSpeed(upstreamSpeed *float64) {
	o.UpstreamSpeed = upstreamSpeed
}

// WithUpstreamSpeedGt adds the upstreamSpeedGt to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithUpstreamSpeedGt(upstreamSpeedGt *string) *CircuitsCircuitTerminationsListParams {
	o.SetUpstreamSpeedGt(upstreamSpeedGt)
	return o
}

// SetUpstreamSpeedGt adds the upstreamSpeedGt to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetUpstreamSpeedGt(upstreamSpeedGt *string) {
	o.UpstreamSpeedGt = upstreamSpeedGt
}

// WithUpstreamSpeedGte adds the upstreamSpeedGte to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithUpstreamSpeedGte(upstreamSpeedGte *string) *CircuitsCircuitTerminationsListParams {
	o.SetUpstreamSpeedGte(upstreamSpeedGte)
	return o
}

// SetUpstreamSpeedGte adds the upstreamSpeedGte to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetUpstreamSpeedGte(upstreamSpeedGte *string) {
	o.UpstreamSpeedGte = upstreamSpeedGte
}

// WithUpstreamSpeedLt adds the upstreamSpeedLt to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithUpstreamSpeedLt(upstreamSpeedLt *string) *CircuitsCircuitTerminationsListParams {
	o.SetUpstreamSpeedLt(upstreamSpeedLt)
	return o
}

// SetUpstreamSpeedLt adds the upstreamSpeedLt to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetUpstreamSpeedLt(upstreamSpeedLt *string) {
	o.UpstreamSpeedLt = upstreamSpeedLt
}

// WithUpstreamSpeedLte adds the upstreamSpeedLte to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithUpstreamSpeedLte(upstreamSpeedLte *string) *CircuitsCircuitTerminationsListParams {
	o.SetUpstreamSpeedLte(upstreamSpeedLte)
	return o
}

// SetUpstreamSpeedLte adds the upstreamSpeedLte to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetUpstreamSpeedLte(upstreamSpeedLte *string) {
	o.UpstreamSpeedLte = upstreamSpeedLte
}

// WithUpstreamSpeedn adds the upstreamSpeedn to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithUpstreamSpeedn(upstreamSpeedn *string) *CircuitsCircuitTerminationsListParams {
	o.SetUpstreamSpeedn(upstreamSpeedn)
	return o
}

// SetUpstreamSpeedn adds the upstreamSpeedN to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetUpstreamSpeedn(upstreamSpeedn *string) {
	o.UpstreamSpeedn = upstreamSpeedn
}

// WithXconnectID adds the xconnectID to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectID(xconnectID *int64) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectID(xconnectID)
	return o
}

// SetXconnectID adds the xconnectId to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectID(xconnectID *int64) {
	o.XconnectID = xconnectID
}

// WithXconnectIDIc adds the xconnectIDIc to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectIDIc(xconnectIDIc *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectIDIc(xconnectIDIc)
	return o
}

// SetXconnectIDIc adds the xconnectIdIc to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectIDIc(xconnectIDIc *string) {
	o.XconnectIDIc = xconnectIDIc
}

// WithXconnectIDIe adds the xconnectIDIe to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectIDIe(xconnectIDIe *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectIDIe(xconnectIDIe)
	return o
}

// SetXconnectIDIe adds the xconnectIdIe to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectIDIe(xconnectIDIe *string) {
	o.XconnectIDIe = xconnectIDIe
}

// WithXconnectIDIew adds the xconnectIDIew to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectIDIew(xconnectIDIew *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectIDIew(xconnectIDIew)
	return o
}

// SetXconnectIDIew adds the xconnectIdIew to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectIDIew(xconnectIDIew *string) {
	o.XconnectIDIew = xconnectIDIew
}

// WithXconnectIDIsw adds the xconnectIDIsw to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectIDIsw(xconnectIDIsw *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectIDIsw(xconnectIDIsw)
	return o
}

// SetXconnectIDIsw adds the xconnectIdIsw to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectIDIsw(xconnectIDIsw *string) {
	o.XconnectIDIsw = xconnectIDIsw
}

// WithXconnectIDn adds the xconnectIDn to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectIDn(xconnectIDn *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectIDn(xconnectIDn)
	return o
}

// SetXconnectIDn adds the xconnectIdN to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectIDn(xconnectIDn *string) {
	o.XconnectIDn = xconnectIDn
}

// WithXconnectIDNic adds the xconnectIDNic to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectIDNic(xconnectIDNic *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectIDNic(xconnectIDNic)
	return o
}

// SetXconnectIDNic adds the xconnectIdNic to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectIDNic(xconnectIDNic *string) {
	o.XconnectIDNic = xconnectIDNic
}

// WithXconnectIDNie adds the xconnectIDNie to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectIDNie(xconnectIDNie *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectIDNie(xconnectIDNie)
	return o
}

// SetXconnectIDNie adds the xconnectIdNie to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectIDNie(xconnectIDNie *string) {
	o.XconnectIDNie = xconnectIDNie
}

// WithXconnectIDNiew adds the xconnectIDNiew to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectIDNiew(xconnectIDNiew *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectIDNiew(xconnectIDNiew)
	return o
}

// SetXconnectIDNiew adds the xconnectIdNiew to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectIDNiew(xconnectIDNiew *string) {
	o.XconnectIDNiew = xconnectIDNiew
}

// WithXconnectIDNisw adds the xconnectIDNisw to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectIDNisw(xconnectIDNisw *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectIDNisw(xconnectIDNisw)
	return o
}

// SetXconnectIDNisw adds the xconnectIdNisw to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectIDNisw(xconnectIDNisw *string) {
	o.XconnectIDNisw = xconnectIDNisw
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTerminationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CircuitID != nil {

		// query param circuit_id
		var qrCircuitID int64
		if o.CircuitID != nil {
			qrCircuitID = *o.CircuitID
		}
		qCircuitID := swag.FormatInt64(qrCircuitID)
		if qCircuitID != "" {
			if err := r.SetQueryParam("circuit_id", qCircuitID); err != nil {
				return err
			}
		}

	}

	if o.CircuitIDn != nil {

		// query param circuit_id__n
		var qrCircuitIDn string
		if o.CircuitIDn != nil {
			qrCircuitIDn = *o.CircuitIDn
		}
		qCircuitIDn := qrCircuitIDn
		if qCircuitIDn != "" {
			if err := r.SetQueryParam("circuit_id__n", qCircuitIDn); err != nil {
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

	if o.PortSpeed != nil {

		// query param port_speed
		var qrPortSpeed float64
		if o.PortSpeed != nil {
			qrPortSpeed = *o.PortSpeed
		}
		qPortSpeed := swag.FormatFloat64(qrPortSpeed)
		if qPortSpeed != "" {
			if err := r.SetQueryParam("port_speed", qPortSpeed); err != nil {
				return err
			}
		}

	}

	if o.PortSpeedGt != nil {

		// query param port_speed__gt
		var qrPortSpeedGt string
		if o.PortSpeedGt != nil {
			qrPortSpeedGt = *o.PortSpeedGt
		}
		qPortSpeedGt := qrPortSpeedGt
		if qPortSpeedGt != "" {
			if err := r.SetQueryParam("port_speed__gt", qPortSpeedGt); err != nil {
				return err
			}
		}

	}

	if o.PortSpeedGte != nil {

		// query param port_speed__gte
		var qrPortSpeedGte string
		if o.PortSpeedGte != nil {
			qrPortSpeedGte = *o.PortSpeedGte
		}
		qPortSpeedGte := qrPortSpeedGte
		if qPortSpeedGte != "" {
			if err := r.SetQueryParam("port_speed__gte", qPortSpeedGte); err != nil {
				return err
			}
		}

	}

	if o.PortSpeedLt != nil {

		// query param port_speed__lt
		var qrPortSpeedLt string
		if o.PortSpeedLt != nil {
			qrPortSpeedLt = *o.PortSpeedLt
		}
		qPortSpeedLt := qrPortSpeedLt
		if qPortSpeedLt != "" {
			if err := r.SetQueryParam("port_speed__lt", qPortSpeedLt); err != nil {
				return err
			}
		}

	}

	if o.PortSpeedLte != nil {

		// query param port_speed__lte
		var qrPortSpeedLte string
		if o.PortSpeedLte != nil {
			qrPortSpeedLte = *o.PortSpeedLte
		}
		qPortSpeedLte := qrPortSpeedLte
		if qPortSpeedLte != "" {
			if err := r.SetQueryParam("port_speed__lte", qPortSpeedLte); err != nil {
				return err
			}
		}

	}

	if o.PortSpeedn != nil {

		// query param port_speed__n
		var qrPortSpeedn string
		if o.PortSpeedn != nil {
			qrPortSpeedn = *o.PortSpeedn
		}
		qPortSpeedn := qrPortSpeedn
		if qPortSpeedn != "" {
			if err := r.SetQueryParam("port_speed__n", qPortSpeedn); err != nil {
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

	if o.Siten != nil {

		// query param site__n
		var qrSiten string
		if o.Siten != nil {
			qrSiten = *o.Siten
		}
		qSiten := qrSiten
		if qSiten != "" {
			if err := r.SetQueryParam("site__n", qSiten); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID int64
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := swag.FormatInt64(qrSiteID)
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if o.SiteIDn != nil {

		// query param site_id__n
		var qrSiteIDn string
		if o.SiteIDn != nil {
			qrSiteIDn = *o.SiteIDn
		}
		qSiteIDn := qrSiteIDn
		if qSiteIDn != "" {
			if err := r.SetQueryParam("site_id__n", qSiteIDn); err != nil {
				return err
			}
		}

	}

	if o.TermSide != nil {

		// query param term_side
		var qrTermSide string
		if o.TermSide != nil {
			qrTermSide = *o.TermSide
		}
		qTermSide := qrTermSide
		if qTermSide != "" {
			if err := r.SetQueryParam("term_side", qTermSide); err != nil {
				return err
			}
		}

	}

	if o.TermSiden != nil {

		// query param term_side__n
		var qrTermSiden string
		if o.TermSiden != nil {
			qrTermSiden = *o.TermSiden
		}
		qTermSiden := qrTermSiden
		if qTermSiden != "" {
			if err := r.SetQueryParam("term_side__n", qTermSiden); err != nil {
				return err
			}
		}

	}

	if o.UpstreamSpeed != nil {

		// query param upstream_speed
		var qrUpstreamSpeed float64
		if o.UpstreamSpeed != nil {
			qrUpstreamSpeed = *o.UpstreamSpeed
		}
		qUpstreamSpeed := swag.FormatFloat64(qrUpstreamSpeed)
		if qUpstreamSpeed != "" {
			if err := r.SetQueryParam("upstream_speed", qUpstreamSpeed); err != nil {
				return err
			}
		}

	}

	if o.UpstreamSpeedGt != nil {

		// query param upstream_speed__gt
		var qrUpstreamSpeedGt string
		if o.UpstreamSpeedGt != nil {
			qrUpstreamSpeedGt = *o.UpstreamSpeedGt
		}
		qUpstreamSpeedGt := qrUpstreamSpeedGt
		if qUpstreamSpeedGt != "" {
			if err := r.SetQueryParam("upstream_speed__gt", qUpstreamSpeedGt); err != nil {
				return err
			}
		}

	}

	if o.UpstreamSpeedGte != nil {

		// query param upstream_speed__gte
		var qrUpstreamSpeedGte string
		if o.UpstreamSpeedGte != nil {
			qrUpstreamSpeedGte = *o.UpstreamSpeedGte
		}
		qUpstreamSpeedGte := qrUpstreamSpeedGte
		if qUpstreamSpeedGte != "" {
			if err := r.SetQueryParam("upstream_speed__gte", qUpstreamSpeedGte); err != nil {
				return err
			}
		}

	}

	if o.UpstreamSpeedLt != nil {

		// query param upstream_speed__lt
		var qrUpstreamSpeedLt string
		if o.UpstreamSpeedLt != nil {
			qrUpstreamSpeedLt = *o.UpstreamSpeedLt
		}
		qUpstreamSpeedLt := qrUpstreamSpeedLt
		if qUpstreamSpeedLt != "" {
			if err := r.SetQueryParam("upstream_speed__lt", qUpstreamSpeedLt); err != nil {
				return err
			}
		}

	}

	if o.UpstreamSpeedLte != nil {

		// query param upstream_speed__lte
		var qrUpstreamSpeedLte string
		if o.UpstreamSpeedLte != nil {
			qrUpstreamSpeedLte = *o.UpstreamSpeedLte
		}
		qUpstreamSpeedLte := qrUpstreamSpeedLte
		if qUpstreamSpeedLte != "" {
			if err := r.SetQueryParam("upstream_speed__lte", qUpstreamSpeedLte); err != nil {
				return err
			}
		}

	}

	if o.UpstreamSpeedn != nil {

		// query param upstream_speed__n
		var qrUpstreamSpeedn string
		if o.UpstreamSpeedn != nil {
			qrUpstreamSpeedn = *o.UpstreamSpeedn
		}
		qUpstreamSpeedn := qrUpstreamSpeedn
		if qUpstreamSpeedn != "" {
			if err := r.SetQueryParam("upstream_speed__n", qUpstreamSpeedn); err != nil {
				return err
			}
		}

	}

	if o.XconnectID != nil {

		// query param xconnect_id
		var qrXconnectID int64
		if o.XconnectID != nil {
			qrXconnectID = *o.XconnectID
		}
		qXconnectID := swag.FormatInt64(qrXconnectID)
		if qXconnectID != "" {
			if err := r.SetQueryParam("xconnect_id", qXconnectID); err != nil {
				return err
			}
		}

	}

	if o.XconnectIDIc != nil {

		// query param xconnect_id__ic
		var qrXconnectIDIc string
		if o.XconnectIDIc != nil {
			qrXconnectIDIc = *o.XconnectIDIc
		}
		qXconnectIDIc := qrXconnectIDIc
		if qXconnectIDIc != "" {
			if err := r.SetQueryParam("xconnect_id__ic", qXconnectIDIc); err != nil {
				return err
			}
		}

	}

	if o.XconnectIDIe != nil {

		// query param xconnect_id__ie
		var qrXconnectIDIe string
		if o.XconnectIDIe != nil {
			qrXconnectIDIe = *o.XconnectIDIe
		}
		qXconnectIDIe := qrXconnectIDIe
		if qXconnectIDIe != "" {
			if err := r.SetQueryParam("xconnect_id__ie", qXconnectIDIe); err != nil {
				return err
			}
		}

	}

	if o.XconnectIDIew != nil {

		// query param xconnect_id__iew
		var qrXconnectIDIew string
		if o.XconnectIDIew != nil {
			qrXconnectIDIew = *o.XconnectIDIew
		}
		qXconnectIDIew := qrXconnectIDIew
		if qXconnectIDIew != "" {
			if err := r.SetQueryParam("xconnect_id__iew", qXconnectIDIew); err != nil {
				return err
			}
		}

	}

	if o.XconnectIDIsw != nil {

		// query param xconnect_id__isw
		var qrXconnectIDIsw string
		if o.XconnectIDIsw != nil {
			qrXconnectIDIsw = *o.XconnectIDIsw
		}
		qXconnectIDIsw := qrXconnectIDIsw
		if qXconnectIDIsw != "" {
			if err := r.SetQueryParam("xconnect_id__isw", qXconnectIDIsw); err != nil {
				return err
			}
		}

	}

	if o.XconnectIDn != nil {

		// query param xconnect_id__n
		var qrXconnectIDn string
		if o.XconnectIDn != nil {
			qrXconnectIDn = *o.XconnectIDn
		}
		qXconnectIDn := qrXconnectIDn
		if qXconnectIDn != "" {
			if err := r.SetQueryParam("xconnect_id__n", qXconnectIDn); err != nil {
				return err
			}
		}

	}

	if o.XconnectIDNic != nil {

		// query param xconnect_id__nic
		var qrXconnectIDNic string
		if o.XconnectIDNic != nil {
			qrXconnectIDNic = *o.XconnectIDNic
		}
		qXconnectIDNic := qrXconnectIDNic
		if qXconnectIDNic != "" {
			if err := r.SetQueryParam("xconnect_id__nic", qXconnectIDNic); err != nil {
				return err
			}
		}

	}

	if o.XconnectIDNie != nil {

		// query param xconnect_id__nie
		var qrXconnectIDNie string
		if o.XconnectIDNie != nil {
			qrXconnectIDNie = *o.XconnectIDNie
		}
		qXconnectIDNie := qrXconnectIDNie
		if qXconnectIDNie != "" {
			if err := r.SetQueryParam("xconnect_id__nie", qXconnectIDNie); err != nil {
				return err
			}
		}

	}

	if o.XconnectIDNiew != nil {

		// query param xconnect_id__niew
		var qrXconnectIDNiew string
		if o.XconnectIDNiew != nil {
			qrXconnectIDNiew = *o.XconnectIDNiew
		}
		qXconnectIDNiew := qrXconnectIDNiew
		if qXconnectIDNiew != "" {
			if err := r.SetQueryParam("xconnect_id__niew", qXconnectIDNiew); err != nil {
				return err
			}
		}

	}

	if o.XconnectIDNisw != nil {

		// query param xconnect_id__nisw
		var qrXconnectIDNisw string
		if o.XconnectIDNisw != nil {
			qrXconnectIDNisw = *o.XconnectIDNisw
		}
		qXconnectIDNisw := qrXconnectIDNisw
		if qXconnectIDNisw != "" {
			if err := r.SetQueryParam("xconnect_id__nisw", qXconnectIDNisw); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
