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

// DcimFrontPortTemplatesCreateReader is a Reader for the DcimFrontPortTemplatesCreate structure.
type DcimFrontPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimFrontPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortTemplatesCreateCreated creates a DcimFrontPortTemplatesCreateCreated with default headers values
func NewDcimFrontPortTemplatesCreateCreated() *DcimFrontPortTemplatesCreateCreated {
	return &DcimFrontPortTemplatesCreateCreated{}
}

/*DcimFrontPortTemplatesCreateCreated handles this case with default header values.

DcimFrontPortTemplatesCreateCreated dcim front port templates create created
*/
type DcimFrontPortTemplatesCreateCreated struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/front-port-templates/][%d] dcimFrontPortTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimFrontPortTemplatesCreateCreated) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesCreateDefault creates a DcimFrontPortTemplatesCreateDefault with default headers values
func NewDcimFrontPortTemplatesCreateDefault(code int) *DcimFrontPortTemplatesCreateDefault {
	return &DcimFrontPortTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*DcimFrontPortTemplatesCreateDefault handles this case with default header values.

DcimFrontPortTemplatesCreateDefault dcim front port templates create default
*/
type DcimFrontPortTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front port templates create default response
func (o *DcimFrontPortTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/front-port-templates/][%d] dcim_front-port-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
