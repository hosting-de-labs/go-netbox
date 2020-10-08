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

package extras

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

// NewExtrasImageAttachmentsListParams creates a new ExtrasImageAttachmentsListParams object
// with the default values initialized.
func NewExtrasImageAttachmentsListParams() *ExtrasImageAttachmentsListParams {
	var ()
	return &ExtrasImageAttachmentsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsListParamsWithTimeout creates a new ExtrasImageAttachmentsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasImageAttachmentsListParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsListParams {
	var ()
	return &ExtrasImageAttachmentsListParams{

		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsListParamsWithContext creates a new ExtrasImageAttachmentsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasImageAttachmentsListParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsListParams {
	var ()
	return &ExtrasImageAttachmentsListParams{

		Context: ctx,
	}
}

// NewExtrasImageAttachmentsListParamsWithHTTPClient creates a new ExtrasImageAttachmentsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasImageAttachmentsListParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsListParams {
	var ()
	return &ExtrasImageAttachmentsListParams{
		HTTPClient: client,
	}
}

/*ExtrasImageAttachmentsListParams contains all the parameters to send to the API endpoint
for the extras image attachments list operation typically these are written to a http.Request
*/
type ExtrasImageAttachmentsListParams struct {

	/*ContentType*/
	ContentType *string
	/*ContentTypen*/
	ContentTypen *string
	/*ID*/
	ID *int64
	/*IDGt*/
	IDGt *string
	/*IDGte*/
	IDGte *string
	/*IDLt*/
	IDLt *string
	/*IDLte*/
	IDLte *string
	/*IDn*/
	IDn *string
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
	/*ObjectID*/
	ObjectID *int64
	/*ObjectIDGt*/
	ObjectIDGt *string
	/*ObjectIDGte*/
	ObjectIDGte *string
	/*ObjectIDLt*/
	ObjectIDLt *string
	/*ObjectIDLte*/
	ObjectIDLte *string
	/*ObjectIDn*/
	ObjectIDn *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithContentType(contentType *string) *ExtrasImageAttachmentsListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithContentTypen adds the contentTypen to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithContentTypen(contentTypen *string) *ExtrasImageAttachmentsListParams {
	o.SetContentTypen(contentTypen)
	return o
}

// SetContentTypen adds the contentTypeN to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetContentTypen(contentTypen *string) {
	o.ContentTypen = contentTypen
}

// WithID adds the id to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithID(id *int64) *ExtrasImageAttachmentsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetID(id *int64) {
	o.ID = id
}

// WithIDGt adds the iDGt to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDGt(iDGt *string) *ExtrasImageAttachmentsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDGte(iDGte *string) *ExtrasImageAttachmentsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDLt(iDLt *string) *ExtrasImageAttachmentsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDLte(iDLte *string) *ExtrasImageAttachmentsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDn(iDn *string) *ExtrasImageAttachmentsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithLimit(limit *int64) *ExtrasImageAttachmentsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithName(name *string) *ExtrasImageAttachmentsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameIc(nameIc *string) *ExtrasImageAttachmentsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameIe(nameIe *string) *ExtrasImageAttachmentsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameIew(nameIew *string) *ExtrasImageAttachmentsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameIsw(nameIsw *string) *ExtrasImageAttachmentsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNamen(namen *string) *ExtrasImageAttachmentsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameNic(nameNic *string) *ExtrasImageAttachmentsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameNie(nameNie *string) *ExtrasImageAttachmentsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameNiew(nameNiew *string) *ExtrasImageAttachmentsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameNisw(nameNisw *string) *ExtrasImageAttachmentsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithObjectID adds the objectID to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectID(objectID *int64) *ExtrasImageAttachmentsListParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectID(objectID *int64) {
	o.ObjectID = objectID
}

// WithObjectIDGt adds the objectIDGt to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDGt(objectIDGt *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDGt(objectIDGt)
	return o
}

// SetObjectIDGt adds the objectIdGt to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDGt(objectIDGt *string) {
	o.ObjectIDGt = objectIDGt
}

// WithObjectIDGte adds the objectIDGte to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDGte(objectIDGte *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDGte(objectIDGte)
	return o
}

// SetObjectIDGte adds the objectIdGte to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDGte(objectIDGte *string) {
	o.ObjectIDGte = objectIDGte
}

// WithObjectIDLt adds the objectIDLt to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDLt(objectIDLt *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDLt(objectIDLt)
	return o
}

// SetObjectIDLt adds the objectIdLt to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDLt(objectIDLt *string) {
	o.ObjectIDLt = objectIDLt
}

// WithObjectIDLte adds the objectIDLte to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDLte(objectIDLte *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDLte(objectIDLte)
	return o
}

// SetObjectIDLte adds the objectIdLte to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDLte(objectIDLte *string) {
	o.ObjectIDLte = objectIDLte
}

// WithObjectIDn adds the objectIDn to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDn(objectIDn *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDn(objectIDn)
	return o
}

// SetObjectIDn adds the objectIdN to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDn(objectIDn *string) {
	o.ObjectIDn = objectIDn
}

// WithOffset adds the offset to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithOffset(offset *int64) *ExtrasImageAttachmentsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// query param content_type
		var qrContentType string
		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {
			if err := r.SetQueryParam("content_type", qContentType); err != nil {
				return err
			}
		}

	}

	if o.ContentTypen != nil {

		// query param content_type__n
		var qrContentTypen string
		if o.ContentTypen != nil {
			qrContentTypen = *o.ContentTypen
		}
		qContentTypen := qrContentTypen
		if qContentTypen != "" {
			if err := r.SetQueryParam("content_type__n", qContentTypen); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID int64
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string
		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {
			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}

	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string
		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {
			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}

	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string
		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {
			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}

	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string
		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {
			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}

	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string
		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {
			if err := r.SetQueryParam("id__n", qIDn); err != nil {
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

	if o.ObjectID != nil {

		// query param object_id
		var qrObjectID int64
		if o.ObjectID != nil {
			qrObjectID = *o.ObjectID
		}
		qObjectID := swag.FormatInt64(qrObjectID)
		if qObjectID != "" {
			if err := r.SetQueryParam("object_id", qObjectID); err != nil {
				return err
			}
		}

	}

	if o.ObjectIDGt != nil {

		// query param object_id__gt
		var qrObjectIDGt string
		if o.ObjectIDGt != nil {
			qrObjectIDGt = *o.ObjectIDGt
		}
		qObjectIDGt := qrObjectIDGt
		if qObjectIDGt != "" {
			if err := r.SetQueryParam("object_id__gt", qObjectIDGt); err != nil {
				return err
			}
		}

	}

	if o.ObjectIDGte != nil {

		// query param object_id__gte
		var qrObjectIDGte string
		if o.ObjectIDGte != nil {
			qrObjectIDGte = *o.ObjectIDGte
		}
		qObjectIDGte := qrObjectIDGte
		if qObjectIDGte != "" {
			if err := r.SetQueryParam("object_id__gte", qObjectIDGte); err != nil {
				return err
			}
		}

	}

	if o.ObjectIDLt != nil {

		// query param object_id__lt
		var qrObjectIDLt string
		if o.ObjectIDLt != nil {
			qrObjectIDLt = *o.ObjectIDLt
		}
		qObjectIDLt := qrObjectIDLt
		if qObjectIDLt != "" {
			if err := r.SetQueryParam("object_id__lt", qObjectIDLt); err != nil {
				return err
			}
		}

	}

	if o.ObjectIDLte != nil {

		// query param object_id__lte
		var qrObjectIDLte string
		if o.ObjectIDLte != nil {
			qrObjectIDLte = *o.ObjectIDLte
		}
		qObjectIDLte := qrObjectIDLte
		if qObjectIDLte != "" {
			if err := r.SetQueryParam("object_id__lte", qObjectIDLte); err != nil {
				return err
			}
		}

	}

	if o.ObjectIDn != nil {

		// query param object_id__n
		var qrObjectIDn string
		if o.ObjectIDn != nil {
			qrObjectIDn = *o.ObjectIDn
		}
		qObjectIDn := qrObjectIDn
		if qObjectIDn != "" {
			if err := r.SetQueryParam("object_id__n", qObjectIDn); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
