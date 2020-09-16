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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hosting-de-labs/go-netbox/netbox/models"
)

// ExtrasGraphsCreateReader is a Reader for the ExtrasGraphsCreate structure.
type ExtrasGraphsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasGraphsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasGraphsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasGraphsCreateCreated creates a ExtrasGraphsCreateCreated with default headers values
func NewExtrasGraphsCreateCreated() *ExtrasGraphsCreateCreated {
	return &ExtrasGraphsCreateCreated{}
}

/*ExtrasGraphsCreateCreated handles this case with default header values.

ExtrasGraphsCreateCreated extras graphs create created
*/
type ExtrasGraphsCreateCreated struct {
	Payload *models.Graph
}

func (o *ExtrasGraphsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/graphs/][%d] extrasGraphsCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasGraphsCreateCreated) GetPayload() *models.Graph {
	return o.Payload
}

func (o *ExtrasGraphsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Graph)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasGraphsCreateDefault creates a ExtrasGraphsCreateDefault with default headers values
func NewExtrasGraphsCreateDefault(code int) *ExtrasGraphsCreateDefault {
	return &ExtrasGraphsCreateDefault{
		_statusCode: code,
	}
}

/*ExtrasGraphsCreateDefault handles this case with default header values.

ExtrasGraphsCreateDefault extras graphs create default
*/
type ExtrasGraphsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras graphs create default response
func (o *ExtrasGraphsCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasGraphsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /extras/graphs/][%d] extras_graphs_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasGraphsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasGraphsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
