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

// NewUpdateCampaignStatusParams creates a new UpdateCampaignStatusParams object
// with the default values initialized.
func NewUpdateCampaignStatusParams() *UpdateCampaignStatusParams {
	var ()
	return &UpdateCampaignStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCampaignStatusParamsWithTimeout creates a new UpdateCampaignStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCampaignStatusParamsWithTimeout(timeout time.Duration) *UpdateCampaignStatusParams {
	var ()
	return &UpdateCampaignStatusParams{

		timeout: timeout,
	}
}

// NewUpdateCampaignStatusParamsWithContext creates a new UpdateCampaignStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCampaignStatusParamsWithContext(ctx context.Context) *UpdateCampaignStatusParams {
	var ()
	return &UpdateCampaignStatusParams{

		Context: ctx,
	}
}

// NewUpdateCampaignStatusParamsWithHTTPClient creates a new UpdateCampaignStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCampaignStatusParamsWithHTTPClient(client *http.Client) *UpdateCampaignStatusParams {
	var ()
	return &UpdateCampaignStatusParams{
		HTTPClient: client,
	}
}

/*UpdateCampaignStatusParams contains all the parameters to send to the API endpoint
for the update campaign status operation typically these are written to a http.Request
*/
type UpdateCampaignStatusParams struct {

	/*CampaignID
	  Id of the campaign

	*/
	CampaignID int64
	/*Status
	  Status of the campaign

	*/
	Status *models.UpdateCampaignStatus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update campaign status params
func (o *UpdateCampaignStatusParams) WithTimeout(timeout time.Duration) *UpdateCampaignStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update campaign status params
func (o *UpdateCampaignStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update campaign status params
func (o *UpdateCampaignStatusParams) WithContext(ctx context.Context) *UpdateCampaignStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update campaign status params
func (o *UpdateCampaignStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update campaign status params
func (o *UpdateCampaignStatusParams) WithHTTPClient(client *http.Client) *UpdateCampaignStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update campaign status params
func (o *UpdateCampaignStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the update campaign status params
func (o *UpdateCampaignStatusParams) WithCampaignID(campaignID int64) *UpdateCampaignStatusParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the update campaign status params
func (o *UpdateCampaignStatusParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithStatus adds the status to the update campaign status params
func (o *UpdateCampaignStatusParams) WithStatus(status *models.UpdateCampaignStatus) *UpdateCampaignStatusParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the update campaign status params
func (o *UpdateCampaignStatusParams) SetStatus(status *models.UpdateCampaignStatus) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCampaignStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.Status != nil {
		if err := r.SetBodyParam(o.Status); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
