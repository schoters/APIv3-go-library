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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// NewGetSmsCampaignParams creates a new GetSmsCampaignParams object
// with the default values initialized.
func NewGetSmsCampaignParams() *GetSmsCampaignParams {
	var ()
	return &GetSmsCampaignParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSmsCampaignParamsWithTimeout creates a new GetSmsCampaignParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSmsCampaignParamsWithTimeout(timeout time.Duration) *GetSmsCampaignParams {
	var ()
	return &GetSmsCampaignParams{

		timeout: timeout,
	}
}

// NewGetSmsCampaignParamsWithContext creates a new GetSmsCampaignParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSmsCampaignParamsWithContext(ctx context.Context) *GetSmsCampaignParams {
	var ()
	return &GetSmsCampaignParams{

		Context: ctx,
	}
}

// NewGetSmsCampaignParamsWithHTTPClient creates a new GetSmsCampaignParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSmsCampaignParamsWithHTTPClient(client *http.Client) *GetSmsCampaignParams {
	var ()
	return &GetSmsCampaignParams{
		HTTPClient: client,
	}
}

/*GetSmsCampaignParams contains all the parameters to send to the API endpoint
for the get sms campaign operation typically these are written to a http.Request
*/
type GetSmsCampaignParams struct {

	/*CampaignID
	  id of the SMS campaign

	*/
	CampaignID int64
	/*GetSmsCampaign
	  Values to update an SMS Campaign

	*/
	GetSmsCampaign *models.GetSmsCampaign

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sms campaign params
func (o *GetSmsCampaignParams) WithTimeout(timeout time.Duration) *GetSmsCampaignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sms campaign params
func (o *GetSmsCampaignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sms campaign params
func (o *GetSmsCampaignParams) WithContext(ctx context.Context) *GetSmsCampaignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sms campaign params
func (o *GetSmsCampaignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sms campaign params
func (o *GetSmsCampaignParams) WithHTTPClient(client *http.Client) *GetSmsCampaignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sms campaign params
func (o *GetSmsCampaignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the get sms campaign params
func (o *GetSmsCampaignParams) WithCampaignID(campaignID int64) *GetSmsCampaignParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the get sms campaign params
func (o *GetSmsCampaignParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithGetSmsCampaign adds the getSmsCampaign to the get sms campaign params
func (o *GetSmsCampaignParams) WithGetSmsCampaign(getSmsCampaign *models.GetSmsCampaign) *GetSmsCampaignParams {
	o.SetGetSmsCampaign(getSmsCampaign)
	return o
}

// SetGetSmsCampaign adds the getSmsCampaign to the get sms campaign params
func (o *GetSmsCampaignParams) SetGetSmsCampaign(getSmsCampaign *models.GetSmsCampaign) {
	o.GetSmsCampaign = getSmsCampaign
}

// WriteToRequest writes these params to a swagger request
func (o *GetSmsCampaignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.GetSmsCampaign != nil {
		if err := r.SetBodyParam(o.GetSmsCampaign); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
