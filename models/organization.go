// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Organization organization
// swagger:model organization
type Organization struct {

	// accepted at
	AcceptedAt strfmt.DateTime `json:"acceptedAt,omitempty"`

	// accepted by
	// Max Length: 64
	AcceptedBy string `json:"acceptedBy,omitempty"`

	// active
	// Required: true
	Active *bool `json:"active"`

	// address line1
	// Required: true
	// Max Length: 210
	// Min Length: 0
	AddressLine1 *string `json:"addressLine1"`

	// address line2
	// Max Length: 210
	// Min Length: 0
	AddressLine2 *string `json:"addressLine2,omitempty"`

	// address line3
	// Max Length: 210
	// Min Length: 0
	AddressLine3 *string `json:"addressLine3,omitempty"`

	// allow new UI
	AllowNewUI bool `json:"allowNewUI,omitempty"`

	// city
	// Required: true
	// Max Length: 50
	// Min Length: 0
	City *string `json:"city"`

	// country
	// Required: true
	// Max Length: 50
	// Min Length: 0
	Country *string `json:"country"`

	// created at
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// created by
	// Max Length: 64
	CreatedBy string `json:"createdBy,omitempty"`

	// free trial hours
	// Required: true
	FreeTrialHours *int32 `json:"freeTrialHours"`

	// id
	ID int32 `json:"id,omitempty"`

	// last modified at
	LastModifiedAt strfmt.DateTime `json:"lastModifiedAt,omitempty"`

	// last modified by
	// Max Length: 64
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// name
	// Required: true
	// Max Length: 100
	// Min Length: 0
	Name *string `json:"name"`

	// payment processor customer Id
	// Max Length: 128
	PaymentProcessorCustomerID string `json:"paymentProcessorCustomerId,omitempty"`

	// postal code
	// Required: true
	// Max Length: 50
	// Min Length: 0
	PostalCode *string `json:"postalCode"`

	// running simulation limit
	// Required: true
	RunningSimulationLimit *int32 `json:"runningSimulationLimit"`

	// saas agreement
	SaasAgreement string `json:"saasAgreement,omitempty"`

	// saas agreement accepted
	// Required: true
	SaasAgreementAccepted *bool `json:"saasAgreementAccepted"`

	// state
	// Required: true
	// Max Length: 50
	// Min Length: 0
	State *string `json:"state"`

	// subscriptions
	Subscriptions []*Subscription `json:"subscriptions"`
}

// Validate validates this organization
func (m *Organization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateActive(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressLine1(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressLine2(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressLine3(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFreeTrialHours(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastModifiedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentProcessorCustomerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRunningSimulationLimit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSaasAgreementAccepted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Organization) validateAcceptedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptedBy) { // not required
		return nil
	}

	if err := validate.MaxLength("acceptedBy", "body", string(m.AcceptedBy), 64); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("addressLine1", "body", m.AddressLine1); err != nil {
		return err
	}

	if err := validate.MinLength("addressLine1", "body", string(*m.AddressLine1), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("addressLine1", "body", string(*m.AddressLine1), 210); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateAddressLine2(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine2) { // not required
		return nil
	}

	if err := validate.MinLength("addressLine2", "body", string(*m.AddressLine2), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("addressLine2", "body", string(*m.AddressLine2), 210); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateAddressLine3(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine3) { // not required
		return nil
	}

	if err := validate.MinLength("addressLine3", "body", string(*m.AddressLine3), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("addressLine3", "body", string(*m.AddressLine3), 210); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	if err := validate.MinLength("city", "body", string(*m.City), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", string(*m.City), 50); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.MinLength("country", "body", string(*m.Country), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("country", "body", string(*m.Country), 50); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if err := validate.MaxLength("createdBy", "body", string(m.CreatedBy), 64); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateFreeTrialHours(formats strfmt.Registry) error {

	if err := validate.Required("freeTrialHours", "body", m.FreeTrialHours); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateLastModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedBy) { // not required
		return nil
	}

	if err := validate.MaxLength("lastModifiedBy", "body", string(m.LastModifiedBy), 64); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validatePaymentProcessorCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentProcessorCustomerID) { // not required
		return nil
	}

	if err := validate.MaxLength("paymentProcessorCustomerId", "body", string(m.PaymentProcessorCustomerID), 128); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validatePostalCode(formats strfmt.Registry) error {

	if err := validate.Required("postalCode", "body", m.PostalCode); err != nil {
		return err
	}

	if err := validate.MinLength("postalCode", "body", string(*m.PostalCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("postalCode", "body", string(*m.PostalCode), 50); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateRunningSimulationLimit(formats strfmt.Registry) error {

	if err := validate.Required("runningSimulationLimit", "body", m.RunningSimulationLimit); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateSaasAgreementAccepted(formats strfmt.Registry) error {

	if err := validate.Required("saasAgreementAccepted", "body", m.SaasAgreementAccepted); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if err := validate.MinLength("state", "body", string(*m.State), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("state", "body", string(*m.State), 50); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateSubscriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Subscriptions) { // not required
		return nil
	}

	for i := 0; i < len(m.Subscriptions); i++ {

		if swag.IsZero(m.Subscriptions[i]) { // not required
			continue
		}

		if m.Subscriptions[i] != nil {

			if err := m.Subscriptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Organization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Organization) UnmarshalBinary(b []byte) error {
	var res Organization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
