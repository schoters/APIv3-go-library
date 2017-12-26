// Code generated by go-swagger; DO NOT EDIT.

package process

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// GetProcessReader is a Reader for the GetProcess structure.
type GetProcessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProcessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetProcessBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProcessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProcessOK creates a GetProcessOK with default headers values
func NewGetProcessOK() *GetProcessOK {
	return &GetProcessOK{}
}

/*GetProcessOK handles this case with default header values.

process informations
*/
type GetProcessOK struct {
	Payload *models.GetProcess
}

func (o *GetProcessOK) Error() string {
	return fmt.Sprintf("[GET /processes/{processId}][%d] getProcessOK  %+v", 200, o.Payload)
}

func (o *GetProcessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetProcess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessBadRequest creates a GetProcessBadRequest with default headers values
func NewGetProcessBadRequest() *GetProcessBadRequest {
	return &GetProcessBadRequest{}
}

/*GetProcessBadRequest handles this case with default header values.

bad request
*/
type GetProcessBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetProcessBadRequest) Error() string {
	return fmt.Sprintf("[GET /processes/{processId}][%d] getProcessBadRequest  %+v", 400, o.Payload)
}

func (o *GetProcessBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessNotFound creates a GetProcessNotFound with default headers values
func NewGetProcessNotFound() *GetProcessNotFound {
	return &GetProcessNotFound{}
}

/*GetProcessNotFound handles this case with default header values.

Process ID not found
*/
type GetProcessNotFound struct {
	Payload *models.ErrorModel
}

func (o *GetProcessNotFound) Error() string {
	return fmt.Sprintf("[GET /processes/{processId}][%d] getProcessNotFound  %+v", 404, o.Payload)
}

func (o *GetProcessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
