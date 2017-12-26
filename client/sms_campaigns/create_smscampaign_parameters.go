// Code generated by go-swagger; DO NOT EDIT.

package sms_campaigns

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

	models "github.com/sendinblue/APIv3-go-library/models"
)

// NewCreateSMSCampaignParams creates a new CreateSMSCampaignParams object
// with the default values initialized.
func NewCreateSMSCampaignParams() *CreateSMSCampaignParams {
	var ()
	return &CreateSMSCampaignParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSMSCampaignParamsWithTimeout creates a new CreateSMSCampaignParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSMSCampaignParamsWithTimeout(timeout time.Duration) *CreateSMSCampaignParams {
	var ()
	return &CreateSMSCampaignParams{

		timeout: timeout,
	}
}

// NewCreateSMSCampaignParamsWithContext creates a new CreateSMSCampaignParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSMSCampaignParamsWithContext(ctx context.Context) *CreateSMSCampaignParams {
	var ()
	return &CreateSMSCampaignParams{

		Context: ctx,
	}
}

// NewCreateSMSCampaignParamsWithHTTPClient creates a new CreateSMSCampaignParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSMSCampaignParamsWithHTTPClient(client *http.Client) *CreateSMSCampaignParams {
	var ()
	return &CreateSMSCampaignParams{
		HTTPClient: client,
	}
}

/*CreateSMSCampaignParams contains all the parameters to send to the API endpoint
for the create s m s campaign operation typically these are written to a http.Request
*/
type CreateSMSCampaignParams struct {

	/*CreateSmsCampaign
	  Values to create an SMS Campaign

	*/
	CreateSmsCampaign *models.CreateSmsCampaign

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create s m s campaign params
func (o *CreateSMSCampaignParams) WithTimeout(timeout time.Duration) *CreateSMSCampaignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create s m s campaign params
func (o *CreateSMSCampaignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create s m s campaign params
func (o *CreateSMSCampaignParams) WithContext(ctx context.Context) *CreateSMSCampaignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create s m s campaign params
func (o *CreateSMSCampaignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create s m s campaign params
func (o *CreateSMSCampaignParams) WithHTTPClient(client *http.Client) *CreateSMSCampaignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create s m s campaign params
func (o *CreateSMSCampaignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateSmsCampaign adds the createSmsCampaign to the create s m s campaign params
func (o *CreateSMSCampaignParams) WithCreateSmsCampaign(createSmsCampaign *models.CreateSmsCampaign) *CreateSMSCampaignParams {
	o.SetCreateSmsCampaign(createSmsCampaign)
	return o
}

// SetCreateSmsCampaign adds the createSmsCampaign to the create s m s campaign params
func (o *CreateSMSCampaignParams) SetCreateSmsCampaign(createSmsCampaign *models.CreateSmsCampaign) {
	o.CreateSmsCampaign = createSmsCampaign
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSMSCampaignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateSmsCampaign != nil {
		if err := r.SetBodyParam(o.CreateSmsCampaign); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
