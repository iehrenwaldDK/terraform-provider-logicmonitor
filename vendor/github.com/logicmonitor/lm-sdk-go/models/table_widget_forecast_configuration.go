// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TableWidgetForecastConfiguration table widget forecast configuration
// swagger:model TableWidgetForecastConfiguration
type TableWidgetForecastConfiguration struct {

	// Forecast method for the widget :Linear | ARIMA
	// Read Only: true
	Algorithm string `json:"algorithm,omitempty"`

	// The percent confidence that should be required for a forecasted alert.
	// Read Only: true
	Confidence int32 `json:"confidence,omitempty"`

	// 	The minimum alert severity the forecasting should include, one of warn | error | critical
	// Read Only: true
	Severity string `json:"severity,omitempty"`

	// The training data time range (the data on which forecasting is calculated). Options are Last 7 days, Last 14 days, Last 30 days, Last calendar month, Last 365 days or a custom time range
	// Read Only: true
	TimeRange string `json:"timeRange,omitempty"`
}

// Validate validates this table widget forecast configuration
func (m *TableWidgetForecastConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TableWidgetForecastConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableWidgetForecastConfiguration) UnmarshalBinary(b []byte) error {
	var res TableWidgetForecastConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}