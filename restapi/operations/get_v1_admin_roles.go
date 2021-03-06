// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/momper14/edgex-restapp/models"
)

// GetV1AdminRolesHandlerFunc turns a function with the right signature into a get v1 admin roles handler
type GetV1AdminRolesHandlerFunc func(GetV1AdminRolesParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1AdminRolesHandlerFunc) Handle(params GetV1AdminRolesParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetV1AdminRolesHandler interface for that can handle valid get v1 admin roles params
type GetV1AdminRolesHandler interface {
	Handle(GetV1AdminRolesParams, *models.User) middleware.Responder
}

// NewGetV1AdminRoles creates a new http.Handler for the get v1 admin roles operation
func NewGetV1AdminRoles(ctx *middleware.Context, handler GetV1AdminRolesHandler) *GetV1AdminRoles {
	return &GetV1AdminRoles{Context: ctx, Handler: handler}
}

/*GetV1AdminRoles swagger:route GET /v1/admin/roles getV1AdminRoles

Find all roles

*/
type GetV1AdminRoles struct {
	Context *middleware.Context
	Handler GetV1AdminRolesHandler
}

func (o *GetV1AdminRoles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1AdminRolesParams()
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
