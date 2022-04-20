// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"tf-provider-example/models"
)

// GetDeviceByIDReader is a Reader for the GetDeviceByID structure.
type GetDeviceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceByIDOK creates a GetDeviceByIDOK with default headers values
func NewGetDeviceByIDOK() *GetDeviceByIDOK {
	return &GetDeviceByIDOK{}
}

/* GetDeviceByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceByIDOK struct {
	Payload *models.Device
}

func (o *GetDeviceByIDOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}][%d] getDeviceByIdOK  %+v", 200, o.Payload)
}
func (o *GetDeviceByIDOK) GetPayload() *models.Device {
	return o.Payload
}

func (o *GetDeviceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Device)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceByIDDefault creates a GetDeviceByIDDefault with default headers values
func NewGetDeviceByIDDefault(code int) *GetDeviceByIDDefault {
	return &GetDeviceByIDDefault{
		_statusCode: code,
	}
}

/* GetDeviceByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device by Id default response
func (o *GetDeviceByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceByIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}][%d] getDeviceById default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
