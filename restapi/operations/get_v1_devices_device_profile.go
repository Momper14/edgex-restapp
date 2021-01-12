// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/momper14/edgex-restapp/models"
)

// GetV1DevicesDeviceProfileHandlerFunc turns a function with the right signature into a get v1 devices device profile handler
type GetV1DevicesDeviceProfileHandlerFunc func(GetV1DevicesDeviceProfileParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1DevicesDeviceProfileHandlerFunc) Handle(params GetV1DevicesDeviceProfileParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetV1DevicesDeviceProfileHandler interface for that can handle valid get v1 devices device profile params
type GetV1DevicesDeviceProfileHandler interface {
	Handle(GetV1DevicesDeviceProfileParams, *models.User) middleware.Responder
}

// NewGetV1DevicesDeviceProfile creates a new http.Handler for the get v1 devices device profile operation
func NewGetV1DevicesDeviceProfile(ctx *middleware.Context, handler GetV1DevicesDeviceProfileHandler) *GetV1DevicesDeviceProfile {
	return &GetV1DevicesDeviceProfile{Context: ctx, Handler: handler}
}

/*GetV1DevicesDeviceProfile swagger:route GET /v1/devices/{device}/profile getV1DevicesDeviceProfile

Get deviceprofile of device

*/
type GetV1DevicesDeviceProfile struct {
	Context *middleware.Context
	Handler GetV1DevicesDeviceProfileHandler
}

func (o *GetV1DevicesDeviceProfile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1DevicesDeviceProfileParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
