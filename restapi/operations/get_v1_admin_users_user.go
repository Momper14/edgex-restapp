// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/momper14/edgex-restapp/models"
)

// GetV1AdminUsersUserHandlerFunc turns a function with the right signature into a get v1 admin users user handler
type GetV1AdminUsersUserHandlerFunc func(GetV1AdminUsersUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1AdminUsersUserHandlerFunc) Handle(params GetV1AdminUsersUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetV1AdminUsersUserHandler interface for that can handle valid get v1 admin users user params
type GetV1AdminUsersUserHandler interface {
	Handle(GetV1AdminUsersUserParams, *models.User) middleware.Responder
}

// NewGetV1AdminUsersUser creates a new http.Handler for the get v1 admin users user operation
func NewGetV1AdminUsersUser(ctx *middleware.Context, handler GetV1AdminUsersUserHandler) *GetV1AdminUsersUser {
	return &GetV1AdminUsersUser{Context: ctx, Handler: handler}
}

/*GetV1AdminUsersUser swagger:route GET /v1/admin/users/{user} getV1AdminUsersUser

Find user by name

*/
type GetV1AdminUsersUser struct {
	Context *middleware.Context
	Handler GetV1AdminUsersUserHandler
}

func (o *GetV1AdminUsersUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1AdminUsersUserParams()
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
