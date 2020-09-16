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

// DcimCablesPartialUpdateReader is a Reader for the DcimCablesPartialUpdate structure.
type DcimCablesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCablesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesPartialUpdateOK creates a DcimCablesPartialUpdateOK with default headers values
func NewDcimCablesPartialUpdateOK() *DcimCablesPartialUpdateOK {
	return &DcimCablesPartialUpdateOK{}
}

/*DcimCablesPartialUpdateOK handles this case with default header values.

DcimCablesPartialUpdateOK dcim cables partial update o k
*/
type DcimCablesPartialUpdateOK struct {
	Payload *models.Cable
}

func (o *DcimCablesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/cables/{id}/][%d] dcimCablesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimCablesPartialUpdateOK) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesPartialUpdateDefault creates a DcimCablesPartialUpdateDefault with default headers values
func NewDcimCablesPartialUpdateDefault(code int) *DcimCablesPartialUpdateDefault {
	return &DcimCablesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimCablesPartialUpdateDefault handles this case with default header values.

DcimCablesPartialUpdateDefault dcim cables partial update default
*/
type DcimCablesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cables partial update default response
func (o *DcimCablesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimCablesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/cables/{id}/][%d] dcim_cables_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
