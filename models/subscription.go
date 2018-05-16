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

// Subscription subscription
// swagger:model subscription
type Subscription struct {

	// active
	Active bool `json:"active,omitempty"`

	// canceled at
	CanceledAt *strfmt.DateTime `json:"canceledAt,omitempty"`

	// canceled by
	// Max Length: 64
	CanceledBy string `json:"canceledBy,omitempty"`

	// created at
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// current period end
	CurrentPeriodEnd strfmt.DateTime `json:"currentPeriodEnd,omitempty"`

	// current period start
	CurrentPeriodStart strfmt.DateTime `json:"currentPeriodStart,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// last modified at
	LastModifiedAt *strfmt.DateTime `json:"lastModifiedAt,omitempty"`

	// last modified by
	// Max Length: 64
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// payment method
	// Max Length: 128
	PaymentMethod string `json:"paymentMethod,omitempty"`

	// payment processor subscription Id
	// Max Length: 128
	PaymentProcessorSubscriptionID string `json:"paymentProcessorSubscriptionId,omitempty"`

	// plan Id
	PlanID int32 `json:"planId,omitempty"`

	// trial end
	TrialEnd *strfmt.DateTime `json:"trialEnd,omitempty"`

	// warning email sent at
	WarningEmailSentAt *strfmt.DateTime `json:"warningEmailSentAt,omitempty"`
}

// Validate validates this subscription
func (m *Subscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanceledBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastModifiedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentProcessorSubscriptionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscription) validateCanceledBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CanceledBy) { // not required
		return nil
	}

	if err := validate.MaxLength("canceledBy", "body", string(m.CanceledBy), 64); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateLastModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedBy) { // not required
		return nil
	}

	if err := validate.MaxLength("lastModifiedBy", "body", string(m.LastModifiedBy), 64); err != nil {
		return err
	}

	return nil
}

var subscriptionTypePaymentMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CreditCard","PurchaseOrder"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTypePaymentMethodPropEnum = append(subscriptionTypePaymentMethodPropEnum, v)
	}
}

const (
	// SubscriptionPaymentMethodCreditCard captures enum value "CreditCard"
	SubscriptionPaymentMethodCreditCard string = "CreditCard"
	// SubscriptionPaymentMethodPurchaseOrder captures enum value "PurchaseOrder"
	SubscriptionPaymentMethodPurchaseOrder string = "PurchaseOrder"
)

// prop value enum
func (m *Subscription) validatePaymentMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionTypePaymentMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validatePaymentMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethod) { // not required
		return nil
	}

	if err := validate.MaxLength("paymentMethod", "body", string(m.PaymentMethod), 128); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentMethodEnum("paymentMethod", "body", m.PaymentMethod); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validatePaymentProcessorSubscriptionID(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentProcessorSubscriptionID) { // not required
		return nil
	}

	if err := validate.MaxLength("paymentProcessorSubscriptionId", "body", string(m.PaymentProcessorSubscriptionID), 128); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscription) UnmarshalBinary(b []byte) error {
	var res Subscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
