// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetV1DevicesDeviceValuedescriptorsParams creates a new GetV1DevicesDeviceValuedescriptorsParams object
// no default values defined in spec.
func NewGetV1DevicesDeviceValuedescriptorsParams() GetV1DevicesDeviceValuedescriptorsParams {

	return GetV1DevicesDeviceValuedescriptorsParams{}
}

// GetV1DevicesDeviceValuedescriptorsParams contains all the bound params for the get v1 devices device valuedescriptors operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetV1DevicesDeviceValuedescriptors
type GetV1DevicesDeviceValuedescriptorsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The name of the device
	  Required: true
	  In: path
	*/
	Device string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetV1DevicesDeviceValuedescriptorsParams() beforehand.
func (o *GetV1DevicesDeviceValuedescriptorsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDevice, rhkDevice, _ := route.Params.GetOK("device")
	if err := o.bindDevice(rDevice, rhkDevice, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDevice binds and validates parameter Device from path.
func (o *GetV1DevicesDeviceValuedescriptorsParams) bindDevice(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Device = raw

	return nil
}
