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

// NewRequestSMSRecipientExportParams creates a new RequestSMSRecipientExportParams object
// with the default values initialized.
func NewRequestSMSRecipientExportParams() *RequestSMSRecipientExportParams {
	var ()
	return &RequestSMSRecipientExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestSMSRecipientExportParamsWithTimeout creates a new RequestSMSRecipientExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestSMSRecipientExportParamsWithTimeout(timeout time.Duration) *RequestSMSRecipientExportParams {
	var ()
	return &RequestSMSRecipientExportParams{

		timeout: timeout,
	}
}

// NewRequestSMSRecipientExportParamsWithContext creates a new RequestSMSRecipientExportParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestSMSRecipientExportParamsWithContext(ctx context.Context) *RequestSMSRecipientExportParams {
	var ()
	return &RequestSMSRecipientExportParams{

		Context: ctx,
	}
}

// NewRequestSMSRecipientExportParamsWithHTTPClient creates a new RequestSMSRecipientExportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestSMSRecipientExportParamsWithHTTPClient(client *http.Client) *RequestSMSRecipientExportParams {
	var ()
	return &RequestSMSRecipientExportParams{
		HTTPClient: client,
	}
}

/*RequestSMSRecipientExportParams contains all the parameters to send to the API endpoint
for the request s m s recipient export operation typically these are written to a http.Request
*/
type RequestSMSRecipientExportParams struct {

	/*CampaignID
	  id of the campaign

	*/
	CampaignID int64
	/*RecipientExport
	  Values to send for a recipient export request

	*/
	RecipientExport *models.RequestSMSRecipientExport

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) WithTimeout(timeout time.Duration) *RequestSMSRecipientExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) WithContext(ctx context.Context) *RequestSMSRecipientExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) WithHTTPClient(client *http.Client) *RequestSMSRecipientExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) WithCampaignID(campaignID int64) *RequestSMSRecipientExportParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithRecipientExport adds the recipientExport to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) WithRecipientExport(recipientExport *models.RequestSMSRecipientExport) *RequestSMSRecipientExportParams {
	o.SetRecipientExport(recipientExport)
	return o
}

// SetRecipientExport adds the recipientExport to the request s m s recipient export params
func (o *RequestSMSRecipientExportParams) SetRecipientExport(recipientExport *models.RequestSMSRecipientExport) {
	o.RecipientExport = recipientExport
}

// WriteToRequest writes these params to a swagger request
func (o *RequestSMSRecipientExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.RecipientExport != nil {
		if err := r.SetBodyParam(o.RecipientExport); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
