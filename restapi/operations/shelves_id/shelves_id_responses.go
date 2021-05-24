// Code generated by go-swagger; DO NOT EDIT.

package shelves_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MarlikAlmighty/library/models"
)

// ShelvesIDOKCode is the HTTP code returned for type ShelvesIDOK
const ShelvesIDOKCode int = 200

/*ShelvesIDOK Successful response

swagger:response shelvesIdOK
*/
type ShelvesIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Shelf `json:"body,omitempty"`
}

// NewShelvesIDOK creates ShelvesIDOK with default headers values
func NewShelvesIDOK() *ShelvesIDOK {

	return &ShelvesIDOK{}
}

// WithPayload adds the payload to the shelves Id o k response
func (o *ShelvesIDOK) WithPayload(payload *models.Shelf) *ShelvesIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the shelves Id o k response
func (o *ShelvesIDOK) SetPayload(payload *models.Shelf) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShelvesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShelvesIDBadRequestCode is the HTTP code returned for type ShelvesIDBadRequest
const ShelvesIDBadRequestCode int = 400

/*ShelvesIDBadRequest Bad request

swagger:response shelvesIdBadRequest
*/
type ShelvesIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Fail `json:"body,omitempty"`
}

// NewShelvesIDBadRequest creates ShelvesIDBadRequest with default headers values
func NewShelvesIDBadRequest() *ShelvesIDBadRequest {

	return &ShelvesIDBadRequest{}
}

// WithPayload adds the payload to the shelves Id bad request response
func (o *ShelvesIDBadRequest) WithPayload(payload *models.Fail) *ShelvesIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the shelves Id bad request response
func (o *ShelvesIDBadRequest) SetPayload(payload *models.Fail) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShelvesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
