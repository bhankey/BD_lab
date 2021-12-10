// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportTurnOverResponse report turn over response
//
// swagger:model ReportTurnOverResponse
type ReportTurnOverResponse struct {

	// account id
	// Minimum: 1
	AccountID int64 `json:"account_id"`

	// end sum
	EndSum float64 `json:"end_sum"`

	// month report
	// Max Items: 12
	// Min Items: 12
	MonthReport []*ReportTurnOverByMonth `json:"month_report"`

	// start sum
	StartSum float64 `json:"start_sum"`
}

// Validate validates this report turn over response
func (m *ReportTurnOverResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonthReport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportTurnOverResponse) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.MinimumInt("account_id", "body", m.AccountID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ReportTurnOverResponse) validateMonthReport(formats strfmt.Registry) error {
	if swag.IsZero(m.MonthReport) { // not required
		return nil
	}

	iMonthReportSize := int64(len(m.MonthReport))

	if err := validate.MinItems("month_report", "body", iMonthReportSize, 12); err != nil {
		return err
	}

	if err := validate.MaxItems("month_report", "body", iMonthReportSize, 12); err != nil {
		return err
	}

	for i := 0; i < len(m.MonthReport); i++ {
		if swag.IsZero(m.MonthReport[i]) { // not required
			continue
		}

		if m.MonthReport[i] != nil {
			if err := m.MonthReport[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("month_report" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("month_report" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this report turn over response based on the context it is used
func (m *ReportTurnOverResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMonthReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportTurnOverResponse) contextValidateMonthReport(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MonthReport); i++ {

		if m.MonthReport[i] != nil {
			if err := m.MonthReport[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("month_report" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("month_report" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportTurnOverResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportTurnOverResponse) UnmarshalBinary(b []byte) error {
	var res ReportTurnOverResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
