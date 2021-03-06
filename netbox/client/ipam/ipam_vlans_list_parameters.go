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

// NewIpamVlansListParams creates a new IpamVlansListParams object
// with the default values initialized.
func NewIpamVlansListParams() *IpamVlansListParams {
	var ()
	return &IpamVlansListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansListParamsWithTimeout creates a new IpamVlansListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamVlansListParamsWithTimeout(timeout time.Duration) *IpamVlansListParams {
	var ()
	return &IpamVlansListParams{

		timeout: timeout,
	}
}

// NewIpamVlansListParamsWithContext creates a new IpamVlansListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamVlansListParamsWithContext(ctx context.Context) *IpamVlansListParams {
	var ()
	return &IpamVlansListParams{

		Context: ctx,
	}
}

// NewIpamVlansListParamsWithHTTPClient creates a new IpamVlansListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamVlansListParamsWithHTTPClient(client *http.Client) *IpamVlansListParams {
	var ()
	return &IpamVlansListParams{
		HTTPClient: client,
	}
}

/*IpamVlansListParams contains all the parameters to send to the API endpoint
for the ipam vlans list operation typically these are written to a http.Request
*/
type IpamVlansListParams struct {

	/*Created*/
	Created *string
	/*CreatedGte*/
	CreatedGte *string
	/*CreatedLte*/
	CreatedLte *string
	/*Group*/
	Group *string
	/*GroupID*/
	GroupID *int64
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
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
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Region*/
	Region *string
	/*RegionID*/
	RegionID *int64
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *int64
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *int64
	/*Status*/
	Status *string
	/*Tag*/
	Tag *string
	/*Tenant*/
	Tenant *string
	/*TenantGroup*/
	TenantGroup *string
	/*TenantGroupID*/
	TenantGroupID *int64
	/*TenantID*/
	TenantID *int64
	/*Vid*/
	Vid *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vlans list params
func (o *IpamVlansListParams) WithTimeout(timeout time.Duration) *IpamVlansListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans list params
func (o *IpamVlansListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans list params
func (o *IpamVlansListParams) WithContext(ctx context.Context) *IpamVlansListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans list params
func (o *IpamVlansListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans list params
func (o *IpamVlansListParams) WithHTTPClient(client *http.Client) *IpamVlansListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans list params
func (o *IpamVlansListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam vlans list params
func (o *IpamVlansListParams) WithCreated(created *string) *IpamVlansListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam vlans list params
func (o *IpamVlansListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam vlans list params
func (o *IpamVlansListParams) WithCreatedGte(createdGte *string) *IpamVlansListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam vlans list params
func (o *IpamVlansListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam vlans list params
func (o *IpamVlansListParams) WithCreatedLte(createdLte *string) *IpamVlansListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam vlans list params
func (o *IpamVlansListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithGroup adds the group to the ipam vlans list params
func (o *IpamVlansListParams) WithGroup(group *string) *IpamVlansListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the ipam vlans list params
func (o *IpamVlansListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupID adds the groupID to the ipam vlans list params
func (o *IpamVlansListParams) WithGroupID(groupID *int64) *IpamVlansListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the ipam vlans list params
func (o *IpamVlansListParams) SetGroupID(groupID *int64) {
	o.GroupID = groupID
}

// WithIDIn adds the iDIn to the ipam vlans list params
func (o *IpamVlansListParams) WithIDIn(iDIn *string) *IpamVlansListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam vlans list params
func (o *IpamVlansListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLastUpdated adds the lastUpdated to the ipam vlans list params
func (o *IpamVlansListParams) WithLastUpdated(lastUpdated *string) *IpamVlansListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam vlans list params
func (o *IpamVlansListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam vlans list params
func (o *IpamVlansListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamVlansListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam vlans list params
func (o *IpamVlansListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam vlans list params
func (o *IpamVlansListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamVlansListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam vlans list params
func (o *IpamVlansListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam vlans list params
func (o *IpamVlansListParams) WithLimit(limit *int64) *IpamVlansListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam vlans list params
func (o *IpamVlansListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam vlans list params
func (o *IpamVlansListParams) WithName(name *string) *IpamVlansListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam vlans list params
func (o *IpamVlansListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the ipam vlans list params
func (o *IpamVlansListParams) WithOffset(offset *int64) *IpamVlansListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam vlans list params
func (o *IpamVlansListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam vlans list params
func (o *IpamVlansListParams) WithQ(q *string) *IpamVlansListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam vlans list params
func (o *IpamVlansListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the ipam vlans list params
func (o *IpamVlansListParams) WithRegion(region *string) *IpamVlansListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the ipam vlans list params
func (o *IpamVlansListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the ipam vlans list params
func (o *IpamVlansListParams) WithRegionID(regionID *int64) *IpamVlansListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the ipam vlans list params
func (o *IpamVlansListParams) SetRegionID(regionID *int64) {
	o.RegionID = regionID
}

// WithRole adds the role to the ipam vlans list params
func (o *IpamVlansListParams) WithRole(role *string) *IpamVlansListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the ipam vlans list params
func (o *IpamVlansListParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the ipam vlans list params
func (o *IpamVlansListParams) WithRoleID(roleID *int64) *IpamVlansListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the ipam vlans list params
func (o *IpamVlansListParams) SetRoleID(roleID *int64) {
	o.RoleID = roleID
}

// WithSite adds the site to the ipam vlans list params
func (o *IpamVlansListParams) WithSite(site *string) *IpamVlansListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the ipam vlans list params
func (o *IpamVlansListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the ipam vlans list params
func (o *IpamVlansListParams) WithSiteID(siteID *int64) *IpamVlansListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the ipam vlans list params
func (o *IpamVlansListParams) SetSiteID(siteID *int64) {
	o.SiteID = siteID
}

// WithStatus adds the status to the ipam vlans list params
func (o *IpamVlansListParams) WithStatus(status *string) *IpamVlansListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the ipam vlans list params
func (o *IpamVlansListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the ipam vlans list params
func (o *IpamVlansListParams) WithTag(tag *string) *IpamVlansListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam vlans list params
func (o *IpamVlansListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the ipam vlans list params
func (o *IpamVlansListParams) WithTenant(tenant *string) *IpamVlansListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam vlans list params
func (o *IpamVlansListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantGroup adds the tenantGroup to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantGroup(tenantGroup *string) *IpamVlansListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupID adds the tenantGroupID to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantGroupID(tenantGroupID *int64) *IpamVlansListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantGroupID(tenantGroupID *int64) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantID adds the tenantID to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantID(tenantID *int64) *IpamVlansListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantID(tenantID *int64) {
	o.TenantID = tenantID
}

// WithVid adds the vid to the ipam vlans list params
func (o *IpamVlansListParams) WithVid(vid *int64) *IpamVlansListParams {
	o.SetVid(vid)
	return o
}

// SetVid adds the vid to the ipam vlans list params
func (o *IpamVlansListParams) SetVid(vid *int64) {
	o.Vid = vid
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID int64
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := swag.FormatInt64(qrRoleID)
		if qRoleID != "" {
			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
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

	if o.Vid != nil {

		// query param vid
		var qrVid int64
		if o.Vid != nil {
			qrVid = *o.Vid
		}
		qVid := swag.FormatInt64(qrVid)
		if qVid != "" {
			if err := r.SetQueryParam("vid", qVid); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
