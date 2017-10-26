// Code generated by go-swagger; DO NOT EDIT.

package smtp

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

// NewSendTransacEmailParams creates a new SendTransacEmailParams object
// with the default values initialized.
func NewSendTransacEmailParams() *SendTransacEmailParams {
	var ()
	return &SendTransacEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendTransacEmailParamsWithTimeout creates a new SendTransacEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendTransacEmailParamsWithTimeout(timeout time.Duration) *SendTransacEmailParams {
	var ()
	return &SendTransacEmailParams{

		timeout: timeout,
	}
}

// NewSendTransacEmailParamsWithContext creates a new SendTransacEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendTransacEmailParamsWithContext(ctx context.Context) *SendTransacEmailParams {
	var ()
	return &SendTransacEmailParams{

		Context: ctx,
	}
}

// NewSendTransacEmailParamsWithHTTPClient creates a new SendTransacEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendTransacEmailParamsWithHTTPClient(client *http.Client) *SendTransacEmailParams {
	var ()
	return &SendTransacEmailParams{
		HTTPClient: client,
	}
}

/*SendTransacEmailParams contains all the parameters to send to the API endpoint
for the send transac email operation typically these are written to a http.Request
*/
type SendTransacEmailParams struct {

	/*SendSMTPEmail
	  Values to send a transactional email

	*/
	SendSMTPEmail *models.SendSMTPEmail

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send transac email params
func (o *SendTransacEmailParams) WithTimeout(timeout time.Duration) *SendTransacEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send transac email params
func (o *SendTransacEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send transac email params
func (o *SendTransacEmailParams) WithContext(ctx context.Context) *SendTransacEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send transac email params
func (o *SendTransacEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send transac email params
func (o *SendTransacEmailParams) WithHTTPClient(client *http.Client) *SendTransacEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send transac email params
func (o *SendTransacEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSendSMTPEmail adds the sendSMTPEmail to the send transac email params
func (o *SendTransacEmailParams) WithSendSMTPEmail(sendSMTPEmail *models.SendSMTPEmail) *SendTransacEmailParams {
	o.SetSendSMTPEmail(sendSMTPEmail)
	return o
}

// SetSendSMTPEmail adds the sendSmtpEmail to the send transac email params
func (o *SendTransacEmailParams) SetSendSMTPEmail(sendSMTPEmail *models.SendSMTPEmail) {
	o.SendSMTPEmail = sendSMTPEmail
}

// WriteToRequest writes these params to a swagger request
func (o *SendTransacEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SendSMTPEmail != nil {
		if err := r.SetBodyParam(o.SendSMTPEmail); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
