// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/momper14/edgex-restapp/models"
)

// GetV1AdminUsersUserOKCode is the HTTP code returned for type GetV1AdminUsersUserOK
const GetV1AdminUsersUserOKCode int = 200

/*GetV1AdminUsersUserOK Successful

swagger:response getV1AdminUsersUserOK
*/
type GetV1AdminUsersUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetV1AdminUsersUserOK creates GetV1AdminUsersUserOK with default headers values
func NewGetV1AdminUsersUserOK() *GetV1AdminUsersUserOK {

	return &GetV1AdminUsersUserOK{}
}

// WithPayload adds the payload to the get v1 admin users user o k response
func (o *GetV1AdminUsersUserOK) WithPayload(payload *models.User) *GetV1AdminUsersUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 admin users user o k response
func (o *GetV1AdminUsersUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1AdminUsersUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1AdminUsersUserUnauthorizedCode is the HTTP code returned for type GetV1AdminUsersUserUnauthorized
const GetV1AdminUsersUserUnauthorizedCode int = 401

/*GetV1AdminUsersUserUnauthorized Unauthorized

swagger:response getV1AdminUsersUserUnauthorized
*/
type GetV1AdminUsersUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetV1AdminUsersUserUnauthorized creates GetV1AdminUsersUserUnauthorized with default headers values
func NewGetV1AdminUsersUserUnauthorized() *GetV1AdminUsersUserUnauthorized {

	return &GetV1AdminUsersUserUnauthorized{}
}

// WithPayload adds the payload to the get v1 admin users user unauthorized response
func (o *GetV1AdminUsersUserUnauthorized) WithPayload(payload string) *GetV1AdminUsersUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 admin users user unauthorized response
func (o *GetV1AdminUsersUserUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1AdminUsersUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetV1AdminUsersUserForbiddenCode is the HTTP code returned for type GetV1AdminUsersUserForbidden
const GetV1AdminUsersUserForbiddenCode int = 403

/*GetV1AdminUsersUserForbidden forbidden

swagger:response getV1AdminUsersUserForbidden
*/
type GetV1AdminUsersUserForbidden struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetV1AdminUsersUserForbidden creates GetV1AdminUsersUserForbidden with default headers values
func NewGetV1AdminUsersUserForbidden() *GetV1AdminUsersUserForbidden {

	return &GetV1AdminUsersUserForbidden{}
}

// WithPayload adds the payload to the get v1 admin users user forbidden response
func (o *GetV1AdminUsersUserForbidden) WithPayload(payload string) *GetV1AdminUsersUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 admin users user forbidden response
func (o *GetV1AdminUsersUserForbidden) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1AdminUsersUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetV1AdminUsersUserNotFoundCode is the HTTP code returned for type GetV1AdminUsersUserNotFound
const GetV1AdminUsersUserNotFoundCode int = 404

/*GetV1AdminUsersUserNotFound The specified resource was not found

swagger:response getV1AdminUsersUserNotFound
*/
type GetV1AdminUsersUserNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetV1AdminUsersUserNotFound creates GetV1AdminUsersUserNotFound with default headers values
func NewGetV1AdminUsersUserNotFound() *GetV1AdminUsersUserNotFound {

	return &GetV1AdminUsersUserNotFound{}
}

// WithPayload adds the payload to the get v1 admin users user not found response
func (o *GetV1AdminUsersUserNotFound) WithPayload(payload string) *GetV1AdminUsersUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 admin users user not found response
func (o *GetV1AdminUsersUserNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1AdminUsersUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetV1AdminUsersUserInternalServerErrorCode is the HTTP code returned for type GetV1AdminUsersUserInternalServerError
const GetV1AdminUsersUserInternalServerErrorCode int = 500

/*GetV1AdminUsersUserInternalServerError For unknown or unanticipated issues

swagger:response getV1AdminUsersUserInternalServerError
*/
type GetV1AdminUsersUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetV1AdminUsersUserInternalServerError creates GetV1AdminUsersUserInternalServerError with default headers values
func NewGetV1AdminUsersUserInternalServerError() *GetV1AdminUsersUserInternalServerError {

	return &GetV1AdminUsersUserInternalServerError{}
}

// WithPayload adds the payload to the get v1 admin users user internal server error response
func (o *GetV1AdminUsersUserInternalServerError) WithPayload(payload string) *GetV1AdminUsersUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 admin users user internal server error response
func (o *GetV1AdminUsersUserInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1AdminUsersUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
