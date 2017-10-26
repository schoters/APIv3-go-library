// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// NewCreateListParams creates a new CreateListParams object
// with the default values initialized.
func NewCreateListParams() *CreateListParams {
	var ()
	return &CreateListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateListParamsWithTimeout creates a new CreateListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateListParamsWithTimeout(timeout time.Duration) *CreateListParams {
	var ()
	return &CreateListParams{

		timeout: timeout,
	}
}

// NewCreateListParamsWithContext creates a new CreateListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateListParamsWithContext(ctx context.Context) *CreateListParams {
	var ()
	return &CreateListParams{

		Context: ctx,
	}
}

// NewCreateListParamsWithHTTPClient creates a new CreateListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateListParamsWithHTTPClient(client *http.Client) *CreateListParams {
	var ()
	return &CreateListParams{
		HTTPClient: client,
	}
}

/*CreateListParams contains all the parameters to send to the API endpoint
for the create list operation typically these are written to a http.Request
*/
type CreateListParams struct {

	/*CreateList
	  Values to create a list

	*/
	CreateList *models.CreateList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create list params
func (o *CreateListParams) WithTimeout(timeout time.Duration) *CreateListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create list params
func (o *CreateListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create list params
func (o *CreateListParams) WithContext(ctx context.Context) *CreateListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create list params
func (o *CreateListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create list params
func (o *CreateListParams) WithHTTPClient(client *http.Client) *CreateListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create list params
func (o *CreateListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateList adds the createList to the create list params
func (o *CreateListParams) WithCreateList(createList *models.CreateList) *CreateListParams {
	o.SetCreateList(createList)
	return o
}

// SetCreateList adds the createList to the create list params
func (o *CreateListParams) SetCreateList(createList *models.CreateList) {
	o.CreateList = createList
}

// WriteToRequest writes these params to a swagger request
func (o *CreateListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateList != nil {
		if err := r.SetBodyParam(o.CreateList); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
