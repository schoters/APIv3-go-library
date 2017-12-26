// Code generated by go-swagger; DO NOT EDIT.

package reseller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// AddCreditsReader is a Reader for the AddCredits structure.
type AddCreditsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCreditsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddCreditsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddCreditsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddCreditsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddCreditsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddCreditsOK creates a AddCreditsOK with default headers values
func NewAddCreditsOK() *AddCreditsOK {
	return &AddCreditsOK{}
}

/*AddCreditsOK handles this case with default header values.

Credits added
*/
type AddCreditsOK struct {
	Payload *models.RemainingCreditModel
}

func (o *AddCreditsOK) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childId}/credits/add][%d] addCreditsOK  %+v", 200, o.Payload)
}

func (o *AddCreditsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemainingCreditModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCreditsBadRequest creates a AddCreditsBadRequest with default headers values
func NewAddCreditsBadRequest() *AddCreditsBadRequest {
	return &AddCreditsBadRequest{}
}

/*AddCreditsBadRequest handles this case with default header values.

bad request
*/
type AddCreditsBadRequest struct {
	Payload *models.ErrorModel
}

func (o *AddCreditsBadRequest) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childId}/credits/add][%d] addCreditsBadRequest  %+v", 400, o.Payload)
}

func (o *AddCreditsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCreditsForbidden creates a AddCreditsForbidden with default headers values
func NewAddCreditsForbidden() *AddCreditsForbidden {
	return &AddCreditsForbidden{}
}

/*AddCreditsForbidden handles this case with default header values.

Current account is not a reseller
*/
type AddCreditsForbidden struct {
	Payload *models.ErrorModel
}

func (o *AddCreditsForbidden) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childId}/credits/add][%d] addCreditsForbidden  %+v", 403, o.Payload)
}

func (o *AddCreditsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCreditsNotFound creates a AddCreditsNotFound with default headers values
func NewAddCreditsNotFound() *AddCreditsNotFound {
	return &AddCreditsNotFound{}
}

/*AddCreditsNotFound handles this case with default header values.

Child ID not found
*/
type AddCreditsNotFound struct {
	Payload *models.ErrorModel
}

func (o *AddCreditsNotFound) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childId}/credits/add][%d] addCreditsNotFound  %+v", 404, o.Payload)
}

func (o *AddCreditsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
