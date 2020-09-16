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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hosting-de-labs/go-netbox/netbox/models"
)

// CircuitsCircuitTypesPartialUpdateReader is a Reader for the CircuitsCircuitTypesPartialUpdate structure.
type CircuitsCircuitTypesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTypesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTypesPartialUpdateOK creates a CircuitsCircuitTypesPartialUpdateOK with default headers values
func NewCircuitsCircuitTypesPartialUpdateOK() *CircuitsCircuitTypesPartialUpdateOK {
	return &CircuitsCircuitTypesPartialUpdateOK{}
}

/*CircuitsCircuitTypesPartialUpdateOK handles this case with default header values.

CircuitsCircuitTypesPartialUpdateOK circuits circuit types partial update o k
*/
type CircuitsCircuitTypesPartialUpdateOK struct {
	Payload *models.CircuitType
}

func (o *CircuitsCircuitTypesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesPartialUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTypesPartialUpdateDefault creates a CircuitsCircuitTypesPartialUpdateDefault with default headers values
func NewCircuitsCircuitTypesPartialUpdateDefault(code int) *CircuitsCircuitTypesPartialUpdateDefault {
	return &CircuitsCircuitTypesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*CircuitsCircuitTypesPartialUpdateDefault handles this case with default header values.

CircuitsCircuitTypesPartialUpdateDefault circuits circuit types partial update default
*/
type CircuitsCircuitTypesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit types partial update default response
func (o *CircuitsCircuitTypesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTypesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-types/{id}/][%d] circuits_circuit-types_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTypesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
