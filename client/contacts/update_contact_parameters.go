// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// NewUpdateContactParams creates a new UpdateContactParams object
// with the default values initialized.
func NewUpdateContactParams() *UpdateContactParams {
	var ()
	return &UpdateContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateContactParamsWithTimeout creates a new UpdateContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateContactParamsWithTimeout(timeout time.Duration) *UpdateContactParams {
	var ()
	return &UpdateContactParams{

		timeout: timeout,
	}
}

// NewUpdateContactParamsWithContext creates a new UpdateContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateContactParamsWithContext(ctx context.Context) *UpdateContactParams {
	var ()
	return &UpdateContactParams{

		Context: ctx,
	}
}

// NewUpdateContactParamsWithHTTPClient creates a new UpdateContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateContactParamsWithHTTPClient(client *http.Client) *UpdateContactParams {
	var ()
	return &UpdateContactParams{
		HTTPClient: client,
	}
}

/*UpdateContactParams contains all the parameters to send to the API endpoint
for the update contact operation typically these are written to a http.Request
*/
type UpdateContactParams struct {

	/*Email
	  Email (urlencoded) of the contact

	*/
	Email string
	/*UpdateContact
	  Values to update a contact

	*/
	UpdateContact *models.UpdateContact

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update contact params
func (o *UpdateContactParams) WithTimeout(timeout time.Duration) *UpdateContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update contact params
func (o *UpdateContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update contact params
func (o *UpdateContactParams) WithContext(ctx context.Context) *UpdateContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update contact params
func (o *UpdateContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update contact params
func (o *UpdateContactParams) WithHTTPClient(client *http.Client) *UpdateContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update contact params
func (o *UpdateContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the update contact params
func (o *UpdateContactParams) WithEmail(email string) *UpdateContactParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the update contact params
func (o *UpdateContactParams) SetEmail(email string) {
	o.Email = email
}

// WithUpdateContact adds the updateContact to the update contact params
func (o *UpdateContactParams) WithUpdateContact(updateContact *models.UpdateContact) *UpdateContactParams {
	o.SetUpdateContact(updateContact)
	return o
}

// SetUpdateContact adds the updateContact to the update contact params
func (o *UpdateContactParams) SetUpdateContact(updateContact *models.UpdateContact) {
	o.UpdateContact = updateContact
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param email
	if err := r.SetPathParam("email", o.Email); err != nil {
		return err
	}

	if o.UpdateContact != nil {
		if err := r.SetBodyParam(o.UpdateContact); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
