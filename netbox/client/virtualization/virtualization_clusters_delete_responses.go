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

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClustersDeleteReader is a Reader for the VirtualizationClustersDelete structure.
type VirtualizationClustersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClustersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClustersDeleteNoContent creates a VirtualizationClustersDeleteNoContent with default headers values
func NewVirtualizationClustersDeleteNoContent() *VirtualizationClustersDeleteNoContent {
	return &VirtualizationClustersDeleteNoContent{}
}

/*VirtualizationClustersDeleteNoContent handles this case with default header values.

VirtualizationClustersDeleteNoContent virtualization clusters delete no content
*/
type VirtualizationClustersDeleteNoContent struct {
}

func (o *VirtualizationClustersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/clusters/{id}/][%d] virtualizationClustersDeleteNoContent ", 204)
}

func (o *VirtualizationClustersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
