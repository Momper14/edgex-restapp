// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/momper14/edgex-restapp/models"
)

// PostV1AdminUsersCreatedCode is the HTTP code returned for type PostV1AdminUsersCreated
const PostV1AdminUsersCreatedCode int = 201

/*PostV1AdminUsersCreated Created

swagger:response postV1AdminUsersCreated
*/
type PostV1AdminUsersCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPostV1AdminUsersCreated creates PostV1AdminUsersCreated with default headers values
func NewPostV1AdminUsersCreated() *PostV1AdminUsersCreated {

	return &PostV1AdminUsersCreated{}
}

// WithPayload adds the payload to the post v1 admin users created response
func (o *PostV1AdminUsersCreated) WithPayload(payload *models.User) *PostV1AdminUsersCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 admin users created response
func (o *PostV1AdminUsersCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1AdminUsersCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostV1AdminUsersUnauthorizedCode is the HTTP code returned for type PostV1AdminUsersUnauthorized
const PostV1AdminUsersUnauthorizedCode int = 401

/*PostV1AdminUsersUnauthorized Unauthorized

swagger:response postV1AdminUsersUnauthorized
*/
type PostV1AdminUsersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostV1AdminUsersUnauthorized creates PostV1AdminUsersUnauthorized with default headers values
func NewPostV1AdminUsersUnauthorized() *PostV1AdminUsersUnauthorized {

	return &PostV1AdminUsersUnauthorized{}
}

// WithPayload adds the payload to the post v1 admin users unauthorized response
func (o *PostV1AdminUsersUnauthorized) WithPayload(payload string) *PostV1AdminUsersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 admin users unauthorized response
func (o *PostV1AdminUsersUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1AdminUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostV1AdminUsersForbiddenCode is the HTTP code returned for type PostV1AdminUsersForbidden
const PostV1AdminUsersForbiddenCode int = 403

/*PostV1AdminUsersForbidden forbidden

swagger:response postV1AdminUsersForbidden
*/
type PostV1AdminUsersForbidden struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostV1AdminUsersForbidden creates PostV1AdminUsersForbidden with default headers values
func NewPostV1AdminUsersForbidden() *PostV1AdminUsersForbidden {

	return &PostV1AdminUsersForbidden{}
}

// WithPayload adds the payload to the post v1 admin users forbidden response
func (o *PostV1AdminUsersForbidden) WithPayload(payload string) *PostV1AdminUsersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 admin users forbidden response
func (o *PostV1AdminUsersForbidden) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1AdminUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostV1AdminUsersConflictCode is the HTTP code returned for type PostV1AdminUsersConflict
const PostV1AdminUsersConflictCode int = 409

/*PostV1AdminUsersConflict conflict

swagger:response postV1AdminUsersConflict
*/
type PostV1AdminUsersConflict struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostV1AdminUsersConflict creates PostV1AdminUsersConflict with default headers values
func NewPostV1AdminUsersConflict() *PostV1AdminUsersConflict {

	return &PostV1AdminUsersConflict{}
}

// WithPayload adds the payload to the post v1 admin users conflict response
func (o *PostV1AdminUsersConflict) WithPayload(payload string) *PostV1AdminUsersConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 admin users conflict response
func (o *PostV1AdminUsersConflict) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1AdminUsersConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostV1AdminUsersInternalServerErrorCode is the HTTP code returned for type PostV1AdminUsersInternalServerError
const PostV1AdminUsersInternalServerErrorCode int = 500

/*PostV1AdminUsersInternalServerError For unknown or unanticipated issues

swagger:response postV1AdminUsersInternalServerError
*/
type PostV1AdminUsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostV1AdminUsersInternalServerError creates PostV1AdminUsersInternalServerError with default headers values
func NewPostV1AdminUsersInternalServerError() *PostV1AdminUsersInternalServerError {

	return &PostV1AdminUsersInternalServerError{}
}

// WithPayload adds the payload to the post v1 admin users internal server error response
func (o *PostV1AdminUsersInternalServerError) WithPayload(payload string) *PostV1AdminUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 admin users internal server error response
func (o *PostV1AdminUsersInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1AdminUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
