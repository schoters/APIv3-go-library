// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetWebhook get webhook
// swagger:model getWebhook
type GetWebhook struct {

	// Creation UTC date-time of the webhook (YYYY-MM-DDTHH:mm:ss.SSSZ)
	// Required: true
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// Description of the webhook
	// Required: true
	Description *string `json:"description"`

	// events
	// Required: true
	Events []string `json:"events"`

	// ID of the webhook
	// Required: true
	ID *int64 `json:"id"`

	// Last modification UTC date-time of the webhook (YYYY-MM-DDTHH:mm:ss.SSSZ)
	// Required: true
	ModifiedAt *strfmt.DateTime `json:"modifiedAt"`

	// Type of webhook (marketing or transac)
	// Required: true
	Type *string `json:"type"`

	// URL of the webhook
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this get webhook
func (m *GetWebhook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModifiedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWebhook) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetWebhook) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *GetWebhook) validateEvents(formats strfmt.Registry) error {

	if err := validate.Required("events", "body", m.Events); err != nil {
		return err
	}

	return nil
}

func (m *GetWebhook) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GetWebhook) validateModifiedAt(formats strfmt.Registry) error {

	if err := validate.Required("modifiedAt", "body", m.ModifiedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("modifiedAt", "body", "date-time", m.ModifiedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var getWebhookTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["marketing","transac"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getWebhookTypeTypePropEnum = append(getWebhookTypeTypePropEnum, v)
	}
}

const (
	// GetWebhookTypeMarketing captures enum value "marketing"
	GetWebhookTypeMarketing string = "marketing"
	// GetWebhookTypeTransac captures enum value "transac"
	GetWebhookTypeTransac string = "transac"
)

// prop value enum
func (m *GetWebhook) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getWebhookTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetWebhook) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *GetWebhook) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetWebhook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWebhook) UnmarshalBinary(b []byte) error {
	var res GetWebhook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
