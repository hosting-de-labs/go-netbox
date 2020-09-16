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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewVirtualizationClustersListParams creates a new VirtualizationClustersListParams object
// with the default values initialized.
func NewVirtualizationClustersListParams() *VirtualizationClustersListParams {
	var ()
	return &VirtualizationClustersListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersListParamsWithTimeout creates a new VirtualizationClustersListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClustersListParamsWithTimeout(timeout time.Duration) *VirtualizationClustersListParams {
	var ()
	return &VirtualizationClustersListParams{

		timeout: timeout,
	}
}

// NewVirtualizationClustersListParamsWithContext creates a new VirtualizationClustersListParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClustersListParamsWithContext(ctx context.Context) *VirtualizationClustersListParams {
	var ()
	return &VirtualizationClustersListParams{

		Context: ctx,
	}
}

// NewVirtualizationClustersListParamsWithHTTPClient creates a new VirtualizationClustersListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClustersListParamsWithHTTPClient(client *http.Client) *VirtualizationClustersListParams {
	var ()
	return &VirtualizationClustersListParams{
		HTTPClient: client,
	}
}

/*VirtualizationClustersListParams contains all the parameters to send to the API endpoint
for the virtualization clusters list operation typically these are written to a http.Request
*/
type VirtualizationClustersListParams struct {

	/*Created*/
	Created *string
	/*CreatedGte*/
	CreatedGte *string
	/*CreatedLte*/
	CreatedLte *string
	/*Group*/
	Group *string
	/*Groupn*/
	Groupn *string
	/*GroupID*/
	GroupID *int64
	/*GroupIDn*/
	GroupIDn *string
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
	/*LastUpdated*/
	LastUpdated *string
	/*LastUpdatedGte*/
	LastUpdatedGte *string
	/*LastUpdatedLte*/
	LastUpdatedLte *string
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
	/*Tag*/
	Tag *string
	/*Tagn*/
	Tagn *string
	/*Tenant*/
	Tenant *string
	/*Tenantn*/
	Tenantn *string
	/*TenantGroup*/
	TenantGroup *string
	/*TenantGroupn*/
	TenantGroupn *string
	/*TenantGroupID*/
	TenantGroupID *int64
	/*TenantGroupIDn*/
	TenantGroupIDn *string
	/*TenantID*/
	TenantID *int64
	/*TenantIDn*/
	TenantIDn *string
	/*Type*/
	Type *string
	/*Typen*/
	Typen *string
	/*TypeID*/
	TypeID *int64
	/*TypeIDn*/
	TypeIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTimeout(timeout time.Duration) *VirtualizationClustersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithContext(ctx context.Context) *VirtualizationClustersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithHTTPClient(client *http.Client) *VirtualizationClustersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithCreated(created *string) *VirtualizationClustersListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithCreatedGte(createdGte *string) *VirtualizationClustersListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithCreatedLte(createdLte *string) *VirtualizationClustersListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithGroup adds the group to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithGroup(group *string) *VirtualizationClustersListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupn adds the groupn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithGroupn(groupn *string) *VirtualizationClustersListParams {
	o.SetGroupn(groupn)
	return o
}

// SetGroupn adds the groupN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetGroupn(groupn *string) {
	o.Groupn = groupn
}

// WithGroupID adds the groupID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithGroupID(groupID *int64) *VirtualizationClustersListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetGroupID(groupID *int64) {
	o.GroupID = groupID
}

// WithGroupIDn adds the groupIDn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithGroupIDn(groupIDn *string) *VirtualizationClustersListParams {
	o.SetGroupIDn(groupIDn)
	return o
}

// SetGroupIDn adds the groupIdN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetGroupIDn(groupIDn *string) {
	o.GroupIDn = groupIDn
}

// WithID adds the id to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithID(id *int64) *VirtualizationClustersListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetID(id *int64) {
	o.ID = id
}

// WithIDGt adds the iDGt to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithIDGt(iDGt *string) *VirtualizationClustersListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithIDGte(iDGte *string) *VirtualizationClustersListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithIDLt(iDLt *string) *VirtualizationClustersListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithIDLte(iDLte *string) *VirtualizationClustersListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithIDn(iDn *string) *VirtualizationClustersListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithLastUpdated(lastUpdated *string) *VirtualizationClustersListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithLastUpdatedGte(lastUpdatedGte *string) *VirtualizationClustersListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithLastUpdatedLte(lastUpdatedLte *string) *VirtualizationClustersListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithLimit(limit *int64) *VirtualizationClustersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithName(name *string) *VirtualizationClustersListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithNameIc(nameIc *string) *VirtualizationClustersListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithNameIe(nameIe *string) *VirtualizationClustersListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithNameIew(nameIew *string) *VirtualizationClustersListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithNameIsw(nameIsw *string) *VirtualizationClustersListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithNamen(namen *string) *VirtualizationClustersListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithNameNic(nameNic *string) *VirtualizationClustersListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithNameNie(nameNie *string) *VirtualizationClustersListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithNameNiew(nameNiew *string) *VirtualizationClustersListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithNameNisw(nameNisw *string) *VirtualizationClustersListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithOffset(offset *int64) *VirtualizationClustersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithQ(q *string) *VirtualizationClustersListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithRegion(region *string) *VirtualizationClustersListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithRegionn(regionn *string) *VirtualizationClustersListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithRegionID(regionID *int64) *VirtualizationClustersListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetRegionID(regionID *int64) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithRegionIDn(regionIDn *string) *VirtualizationClustersListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithSite(site *string) *VirtualizationClustersListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithSiten(siten *string) *VirtualizationClustersListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithSiteID(siteID *int64) *VirtualizationClustersListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetSiteID(siteID *int64) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithSiteIDn(siteIDn *string) *VirtualizationClustersListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTag(tag *string) *VirtualizationClustersListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTagn(tagn *string) *VirtualizationClustersListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTenant(tenant *string) *VirtualizationClustersListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTenantn(tenantn *string) *VirtualizationClustersListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTenantGroup(tenantGroup *string) *VirtualizationClustersListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTenantGroupn(tenantGroupn *string) *VirtualizationClustersListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTenantGroupID(tenantGroupID *int64) *VirtualizationClustersListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTenantGroupID(tenantGroupID *int64) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTenantGroupIDn(tenantGroupIDn *string) *VirtualizationClustersListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTenantID(tenantID *int64) *VirtualizationClustersListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTenantID(tenantID *int64) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTenantIDn(tenantIDn *string) *VirtualizationClustersListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithType adds the typeVar to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithType(typeVar *string) *VirtualizationClustersListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTypen(typen *string) *VirtualizationClustersListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WithTypeID adds the typeID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTypeID(typeID *int64) *VirtualizationClustersListParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTypeID(typeID *int64) {
	o.TypeID = typeID
}

// WithTypeIDn adds the typeIDn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTypeIDn(typeIDn *string) *VirtualizationClustersListParams {
	o.SetTypeIDn(typeIDn)
	return o
}

// SetTypeIDn adds the typeIdN to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTypeIDn(typeIDn *string) {
	o.TypeIDn = typeIDn
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string
		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {
			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}

	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string
		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {
			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}

	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string
		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {
			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}

	}

	if o.Group != nil {

		// query param group
		var qrGroup string
		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {
			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}

	}

	if o.Groupn != nil {

		// query param group__n
		var qrGroupn string
		if o.Groupn != nil {
			qrGroupn = *o.Groupn
		}
		qGroupn := qrGroupn
		if qGroupn != "" {
			if err := r.SetQueryParam("group__n", qGroupn); err != nil {
				return err
			}
		}

	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID int64
		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := swag.FormatInt64(qrGroupID)
		if qGroupID != "" {
			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
				return err
			}
		}

	}

	if o.GroupIDn != nil {

		// query param group_id__n
		var qrGroupIDn string
		if o.GroupIDn != nil {
			qrGroupIDn = *o.GroupIDn
		}
		qGroupIDn := qrGroupIDn
		if qGroupIDn != "" {
			if err := r.SetQueryParam("group_id__n", qGroupIDn); err != nil {
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

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string
		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {
			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}

	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string
		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {
			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}

	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string
		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {
			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
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

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string
		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {
			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
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

	if o.Tenantn != nil {

		// query param tenant__n
		var qrTenantn string
		if o.Tenantn != nil {
			qrTenantn = *o.Tenantn
		}
		qTenantn := qrTenantn
		if qTenantn != "" {
			if err := r.SetQueryParam("tenant__n", qTenantn); err != nil {
				return err
			}
		}

	}

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string
		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {
			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}

	}

	if o.TenantGroupn != nil {

		// query param tenant_group__n
		var qrTenantGroupn string
		if o.TenantGroupn != nil {
			qrTenantGroupn = *o.TenantGroupn
		}
		qTenantGroupn := qrTenantGroupn
		if qTenantGroupn != "" {
			if err := r.SetQueryParam("tenant_group__n", qTenantGroupn); err != nil {
				return err
			}
		}

	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID int64
		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := swag.FormatInt64(qrTenantGroupID)
		if qTenantGroupID != "" {
			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
				return err
			}
		}

	}

	if o.TenantGroupIDn != nil {

		// query param tenant_group_id__n
		var qrTenantGroupIDn string
		if o.TenantGroupIDn != nil {
			qrTenantGroupIDn = *o.TenantGroupIDn
		}
		qTenantGroupIDn := qrTenantGroupIDn
		if qTenantGroupIDn != "" {
			if err := r.SetQueryParam("tenant_group_id__n", qTenantGroupIDn); err != nil {
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

	if o.TenantIDn != nil {

		// query param tenant_id__n
		var qrTenantIDn string
		if o.TenantIDn != nil {
			qrTenantIDn = *o.TenantIDn
		}
		qTenantIDn := qrTenantIDn
		if qTenantIDn != "" {
			if err := r.SetQueryParam("tenant_id__n", qTenantIDn); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if o.Typen != nil {

		// query param type__n
		var qrTypen string
		if o.Typen != nil {
			qrTypen = *o.Typen
		}
		qTypen := qrTypen
		if qTypen != "" {
			if err := r.SetQueryParam("type__n", qTypen); err != nil {
				return err
			}
		}

	}

	if o.TypeID != nil {

		// query param type_id
		var qrTypeID int64
		if o.TypeID != nil {
			qrTypeID = *o.TypeID
		}
		qTypeID := swag.FormatInt64(qrTypeID)
		if qTypeID != "" {
			if err := r.SetQueryParam("type_id", qTypeID); err != nil {
				return err
			}
		}

	}

	if o.TypeIDn != nil {

		// query param type_id__n
		var qrTypeIDn string
		if o.TypeIDn != nil {
			qrTypeIDn = *o.TypeIDn
		}
		qTypeIDn := qrTypeIDn
		if qTypeIDn != "" {
			if err := r.SetQueryParam("type_id__n", qTypeIDn); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
