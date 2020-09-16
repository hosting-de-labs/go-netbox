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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hosting-de-labs/go-netbox/netbox/models"
)

// DcimVirtualChassisCreateReader is a Reader for the DcimVirtualChassisCreate structure.
type DcimVirtualChassisCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimVirtualChassisCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualChassisCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualChassisCreateCreated creates a DcimVirtualChassisCreateCreated with default headers values
func NewDcimVirtualChassisCreateCreated() *DcimVirtualChassisCreateCreated {
	return &DcimVirtualChassisCreateCreated{}
}

/*DcimVirtualChassisCreateCreated handles this case with default header values.

DcimVirtualChassisCreateCreated dcim virtual chassis create created
*/
type DcimVirtualChassisCreateCreated struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/virtual-chassis/][%d] dcimVirtualChassisCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimVirtualChassisCreateCreated) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualChassisCreateDefault creates a DcimVirtualChassisCreateDefault with default headers values
func NewDcimVirtualChassisCreateDefault(code int) *DcimVirtualChassisCreateDefault {
	return &DcimVirtualChassisCreateDefault{
		_statusCode: code,
	}
}

/*DcimVirtualChassisCreateDefault handles this case with default header values.

DcimVirtualChassisCreateDefault dcim virtual chassis create default
*/
type DcimVirtualChassisCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim virtual chassis create default response
func (o *DcimVirtualChassisCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimVirtualChassisCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/virtual-chassis/][%d] dcim_virtual-chassis_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
