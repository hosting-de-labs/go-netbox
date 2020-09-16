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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIpamVlanGroupsListParams creates a new IpamVlanGroupsListParams object
// with the default values initialized.
func NewIpamVlanGroupsListParams() *IpamVlanGroupsListParams {
	var ()
	return &IpamVlanGroupsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsListParamsWithTimeout creates a new IpamVlanGroupsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamVlanGroupsListParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsListParams {
	var ()
	return &IpamVlanGroupsListParams{

		timeout: timeout,
	}
}

// NewIpamVlanGroupsListParamsWithContext creates a new IpamVlanGroupsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamVlanGroupsListParamsWithContext(ctx context.Context) *IpamVlanGroupsListParams {
	var ()
	return &IpamVlanGroupsListParams{

		Context: ctx,
	}
}

// NewIpamVlanGroupsListParamsWithHTTPClient creates a new IpamVlanGroupsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamVlanGroupsListParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsListParams {
	var ()
	return &IpamVlanGroupsListParams{
		HTTPClient: client,
	}
}

/*IpamVlanGroupsListParams contains all the parameters to send to the API endpoint
for the ipam vlan groups list operation typically these are written to a http.Request
*/
type IpamVlanGroupsListParams struct {

	/*Description*/
	Description *string
	/*DescriptionIc*/
	DescriptionIc *string
	/*DescriptionIe*/
	DescriptionIe *string
	/*DescriptionIew*/
	DescriptionIew *string
	/*DescriptionIsw*/
	DescriptionIsw *string
	/*Descriptionn*/
	Descriptionn *string
	/*DescriptionNic*/
	DescriptionNic *string
	/*DescriptionNie*/
	DescriptionNie *string
	/*DescriptionNiew*/
	DescriptionNiew *string
	/*DescriptionNisw*/
	DescriptionNisw *string
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
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Region*/
	Region *string
	/*Regionn*/
	Regionn *string
	/*RegionID*/
	RegionID *int64
	/*RegionIDn*/
	RegionIDn *string
	/*Site*/
	Site *string
	/*Siten*/
	Siten *string
	/*SiteID*/
	SiteID *int64
	/*SiteIDn*/
	SiteIDn *string
	/*Slug*/
	Slug *string
	/*SlugIc*/
	SlugIc *string
	/*SlugIe*/
	SlugIe *string
	/*SlugIew*/
	SlugIew *string
	/*SlugIsw*/
	SlugIsw *string
	/*Slugn*/
	Slugn *string
	/*SlugNic*/
	SlugNic *string
	/*SlugNie*/
	SlugNie *string
	/*SlugNiew*/
	SlugNiew *string
	/*SlugNisw*/
	SlugNisw *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithContext(ctx context.Context) *IpamVlanGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescription(description *string) *IpamVlanGroupsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionIc(descriptionIc *string) *IpamVlanGroupsListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionIe(descriptionIe *string) *IpamVlanGroupsListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionIew(descriptionIew *string) *IpamVlanGroupsListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionIsw(descriptionIsw *string) *IpamVlanGroupsListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionn(descriptionn *string) *IpamVlanGroupsListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionNic(descriptionNic *string) *IpamVlanGroupsListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionNie(descriptionNie *string) *IpamVlanGroupsListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionNiew(descriptionNiew *string) *IpamVlanGroupsListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionNisw(descriptionNisw *string) *IpamVlanGroupsListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithID adds the id to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithID(id *int64) *IpamVlanGroupsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetID(id *int64) {
	o.ID = id
}

// WithIDGt adds the iDGt to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDGt(iDGt *string) *IpamVlanGroupsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDGte(iDGte *string) *IpamVlanGroupsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDLt(iDLt *string) *IpamVlanGroupsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDLte(iDLte *string) *IpamVlanGroupsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDn(iDn *string) *IpamVlanGroupsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithLimit(limit *int64) *IpamVlanGroupsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithName(name *string) *IpamVlanGroupsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameIc(nameIc *string) *IpamVlanGroupsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameIe(nameIe *string) *IpamVlanGroupsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameIew(nameIew *string) *IpamVlanGroupsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameIsw(nameIsw *string) *IpamVlanGroupsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNamen(namen *string) *IpamVlanGroupsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameNic(nameNic *string) *IpamVlanGroupsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameNie(nameNie *string) *IpamVlanGroupsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameNiew(nameNiew *string) *IpamVlanGroupsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameNisw(nameNisw *string) *IpamVlanGroupsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithOffset(offset *int64) *IpamVlanGroupsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithQ(q *string) *IpamVlanGroupsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithRegion(region *string) *IpamVlanGroupsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithRegionn(regionn *string) *IpamVlanGroupsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithRegionID(regionID *int64) *IpamVlanGroupsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetRegionID(regionID *int64) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithRegionIDn(regionIDn *string) *IpamVlanGroupsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSite(site *string) *IpamVlanGroupsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSiten(siten *string) *IpamVlanGroupsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSiteID(siteID *int64) *IpamVlanGroupsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSiteID(siteID *int64) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSiteIDn(siteIDn *string) *IpamVlanGroupsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithSlug adds the slug to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlug(slug *string) *IpamVlanGroupsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugIc(slugIc *string) *IpamVlanGroupsListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugIe(slugIe *string) *IpamVlanGroupsListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugIew(slugIew *string) *IpamVlanGroupsListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugIsw(slugIsw *string) *IpamVlanGroupsListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugn(slugn *string) *IpamVlanGroupsListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugNic(slugNic *string) *IpamVlanGroupsListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugNie(slugNie *string) *IpamVlanGroupsListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugNiew(slugNiew *string) *IpamVlanGroupsListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugNisw(slugNisw *string) *IpamVlanGroupsListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string
		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {
			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}

	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string
		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {
			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}

	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string
		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {
			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}

	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string
		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {
			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}

	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string
		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {
			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}

	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string
		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {
			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}

	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string
		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {
			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}

	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string
		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {
			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}

	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string
		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {
			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
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

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if o.Regionn != nil {

		// query param region__n
		var qrRegionn string
		if o.Regionn != nil {
			qrRegionn = *o.Regionn
		}
		qRegionn := qrRegionn
		if qRegionn != "" {
			if err := r.SetQueryParam("region__n", qRegionn); err != nil {
				return err
			}
		}

	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID int64
		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := swag.FormatInt64(qrRegionID)
		if qRegionID != "" {
			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
				return err
			}
		}

	}

	if o.RegionIDn != nil {

		// query param region_id__n
		var qrRegionIDn string
		if o.RegionIDn != nil {
			qrRegionIDn = *o.RegionIDn
		}
		qRegionIDn := qrRegionIDn
		if qRegionIDn != "" {
			if err := r.SetQueryParam("region_id__n", qRegionIDn); err != nil {
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

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string
		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {
			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}

	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string
		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {
			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}

	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string
		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {
			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}

	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string
		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {
			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}

	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string
		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {
			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}

	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string
		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {
			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}

	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string
		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {
			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}

	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string
		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {
			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}

	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string
		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {
			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
