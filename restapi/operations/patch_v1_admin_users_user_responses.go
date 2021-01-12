// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PatchV1AdminUsersUserOKCode is the HTTP code returned for type PatchV1AdminUsersUserOK
const PatchV1AdminUsersUserOKCode int = 200

/*PatchV1AdminUsersUserOK Updated

swagger:response patchV1AdminUsersUserOK
*/
type PatchV1AdminUsersUserOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchV1AdminUsersUserOK creates PatchV1AdminUsersUserOK with default headers values
func NewPatchV1AdminUsersUserOK() *PatchV1AdminUsersUserOK {

	return &PatchV1AdminUsersUserOK{}
}

// WithPayload adds the payload to the patch v1 admin users user o k response
func (o *PatchV1AdminUsersUserOK) WithPayload(payload string) *PatchV1AdminUsersUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch v1 admin users user o k response
func (o *PatchV1AdminUsersUserOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchV1AdminUsersUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchV1AdminUsersUserBadRequestCode is the HTTP code returned for type PatchV1AdminUsersUserBadRequest
const PatchV1AdminUsersUserBadRequestCode int = 400

/*PatchV1AdminUsersUserBadRequest Update request is invalid

swagger:response patchV1AdminUsersUserBadRequest
*/
type PatchV1AdminUsersUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchV1AdminUsersUserBadRequest creates PatchV1AdminUsersUserBadRequest with default headers values
func NewPatchV1AdminUsersUserBadRequest() *PatchV1AdminUsersUserBadRequest {

	return &PatchV1AdminUsersUserBadRequest{}
}

// WithPayload adds the payload to the patch v1 admin users user bad request response
func (o *PatchV1AdminUsersUserBadRequest) WithPayload(payload string) *PatchV1AdminUsersUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch v1 admin users user bad request response
func (o *PatchV1AdminUsersUserBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchV1AdminUsersUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchV1AdminUsersUserUnauthorizedCode is the HTTP code returned for type PatchV1AdminUsersUserUnauthorized
const PatchV1AdminUsersUserUnauthorizedCode int = 401

/*PatchV1AdminUsersUserUnauthorized Unauthorized

swagger:response patchV1AdminUsersUserUnauthorized
*/
type PatchV1AdminUsersUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchV1AdminUsersUserUnauthorized creates PatchV1AdminUsersUserUnauthorized with default headers values
func NewPatchV1AdminUsersUserUnauthorized() *PatchV1AdminUsersUserUnauthorized {

	return &PatchV1AdminUsersUserUnauthorized{}
}

// WithPayload adds the payload to the patch v1 admin users user unauthorized response
func (o *PatchV1AdminUsersUserUnauthorized) WithPayload(payload string) *PatchV1AdminUsersUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch v1 admin users user unauthorized response
func (o *PatchV1AdminUsersUserUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchV1AdminUsersUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchV1AdminUsersUserForbiddenCode is the HTTP code returned for type PatchV1AdminUsersUserForbidden
const PatchV1AdminUsersUserForbiddenCode int = 403

/*PatchV1AdminUsersUserForbidden forbidden

swagger:response patchV1AdminUsersUserForbidden
*/
type PatchV1AdminUsersUserForbidden struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchV1AdminUsersUserForbidden creates PatchV1AdminUsersUserForbidden with default headers values
func NewPatchV1AdminUsersUserForbidden() *PatchV1AdminUsersUserForbidden {

	return &PatchV1AdminUsersUserForbidden{}
}

// WithPayload adds the payload to the patch v1 admin users user forbidden response
func (o *PatchV1AdminUsersUserForbidden) WithPayload(payload string) *PatchV1AdminUsersUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch v1 admin users user forbidden response
func (o *PatchV1AdminUsersUserForbidden) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchV1AdminUsersUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchV1AdminUsersUserNotFoundCode is the HTTP code returned for type PatchV1AdminUsersUserNotFound
const PatchV1AdminUsersUserNotFoundCode int = 404

/*PatchV1AdminUsersUserNotFound The specified resource was not found

swagger:response patchV1AdminUsersUserNotFound
*/
type PatchV1AdminUsersUserNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchV1AdminUsersUserNotFound creates PatchV1AdminUsersUserNotFound with default headers values
func NewPatchV1AdminUsersUserNotFound() *PatchV1AdminUsersUserNotFound {

	return &PatchV1AdminUsersUserNotFound{}
}

// WithPayload adds the payload to the patch v1 admin users user not found response
func (o *PatchV1AdminUsersUserNotFound) WithPayload(payload string) *PatchV1AdminUsersUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch v1 admin users user not found response
func (o *PatchV1AdminUsersUserNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchV1AdminUsersUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchV1AdminUsersUserInternalServerErrorCode is the HTTP code returned for type PatchV1AdminUsersUserInternalServerError
const PatchV1AdminUsersUserInternalServerErrorCode int = 500

/*PatchV1AdminUsersUserInternalServerError For unknown or unanticipated issues

swagger:response patchV1AdminUsersUserInternalServerError
*/
type PatchV1AdminUsersUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchV1AdminUsersUserInternalServerError creates PatchV1AdminUsersUserInternalServerError with default headers values
func NewPatchV1AdminUsersUserInternalServerError() *PatchV1AdminUsersUserInternalServerError {

	return &PatchV1AdminUsersUserInternalServerError{}
}

// WithPayload adds the payload to the patch v1 admin users user internal server error response
func (o *PatchV1AdminUsersUserInternalServerError) WithPayload(payload string) *PatchV1AdminUsersUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch v1 admin users user internal server error response
func (o *PatchV1AdminUsersUserInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchV1AdminUsersUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
