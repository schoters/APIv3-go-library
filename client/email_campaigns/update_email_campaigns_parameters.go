// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// NewUpdateEmailCampaignsParams creates a new UpdateEmailCampaignsParams object
// with the default values initialized.
func NewUpdateEmailCampaignsParams() *UpdateEmailCampaignsParams {
	var ()
	return &UpdateEmailCampaignsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEmailCampaignsParamsWithTimeout creates a new UpdateEmailCampaignsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEmailCampaignsParamsWithTimeout(timeout time.Duration) *UpdateEmailCampaignsParams {
	var ()
	return &UpdateEmailCampaignsParams{

		timeout: timeout,
	}
}

// NewUpdateEmailCampaignsParamsWithContext creates a new UpdateEmailCampaignsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEmailCampaignsParamsWithContext(ctx context.Context) *UpdateEmailCampaignsParams {
	var ()
	return &UpdateEmailCampaignsParams{

		Context: ctx,
	}
}

// NewUpdateEmailCampaignsParamsWithHTTPClient creates a new UpdateEmailCampaignsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEmailCampaignsParamsWithHTTPClient(client *http.Client) *UpdateEmailCampaignsParams {
	var ()
	return &UpdateEmailCampaignsParams{
		HTTPClient: client,
	}
}

/*UpdateEmailCampaignsParams contains all the parameters to send to the API endpoint
for the update email campaigns operation typically these are written to a http.Request
*/
type UpdateEmailCampaignsParams struct {

	/*CampaignID
	  Id of the campaign

	*/
	CampaignID int64
	/*EmailCampaign
	  Values to update a campaign

	*/
	EmailCampaign *models.UpdateEmailCampaign

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update email campaigns params
func (o *UpdateEmailCampaignsParams) WithTimeout(timeout time.Duration) *UpdateEmailCampaignsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update email campaigns params
func (o *UpdateEmailCampaignsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update email campaigns params
func (o *UpdateEmailCampaignsParams) WithContext(ctx context.Context) *UpdateEmailCampaignsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update email campaigns params
func (o *UpdateEmailCampaignsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update email campaigns params
func (o *UpdateEmailCampaignsParams) WithHTTPClient(client *http.Client) *UpdateEmailCampaignsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update email campaigns params
func (o *UpdateEmailCampaignsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the update email campaigns params
func (o *UpdateEmailCampaignsParams) WithCampaignID(campaignID int64) *UpdateEmailCampaignsParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the update email campaigns params
func (o *UpdateEmailCampaignsParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithEmailCampaign adds the emailCampaign to the update email campaigns params
func (o *UpdateEmailCampaignsParams) WithEmailCampaign(emailCampaign *models.UpdateEmailCampaign) *UpdateEmailCampaignsParams {
	o.SetEmailCampaign(emailCampaign)
	return o
}

// SetEmailCampaign adds the emailCampaign to the update email campaigns params
func (o *UpdateEmailCampaignsParams) SetEmailCampaign(emailCampaign *models.UpdateEmailCampaign) {
	o.EmailCampaign = emailCampaign
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEmailCampaignsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.EmailCampaign != nil {
		if err := r.SetBodyParam(o.EmailCampaign); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
