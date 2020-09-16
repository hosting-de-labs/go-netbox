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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hosting-de-labs/go-netbox/netbox/models"
)

// VirtualizationClusterTypesUpdateReader is a Reader for the VirtualizationClusterTypesUpdate structure.
type VirtualizationClusterTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterTypesUpdateOK creates a VirtualizationClusterTypesUpdateOK with default headers values
func NewVirtualizationClusterTypesUpdateOK() *VirtualizationClusterTypesUpdateOK {
	return &VirtualizationClusterTypesUpdateOK{}
}

/*VirtualizationClusterTypesUpdateOK handles this case with default header values.

VirtualizationClusterTypesUpdateOK virtualization cluster types update o k
*/
type VirtualizationClusterTypesUpdateOK struct {
	Payload *models.ClusterType
}

func (o *VirtualizationClusterTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterTypesUpdateOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterTypesUpdateDefault creates a VirtualizationClusterTypesUpdateDefault with default headers values
func NewVirtualizationClusterTypesUpdateDefault(code int) *VirtualizationClusterTypesUpdateDefault {
	return &VirtualizationClusterTypesUpdateDefault{
		_statusCode: code,
	}
}

/*VirtualizationClusterTypesUpdateDefault handles this case with default header values.

VirtualizationClusterTypesUpdateDefault virtualization cluster types update default
*/
type VirtualizationClusterTypesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster types update default response
func (o *VirtualizationClusterTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterTypesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/{id}/][%d] virtualization_cluster-types_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterTypesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
