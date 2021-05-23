// Code generated by go-docs; DO NOT EDIT.

package shelves

// This file was generated by the docs tool.
// Editing this file might prove futile when you re-run the docs generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/MarlikAlmighty/library/models"
)

// ShelvesOKCode is the HTTP code returned for type ShelvesOK
const ShelvesOKCode int = 200

/*ShelvesOK Successful response

docs:response shelvesOK
*/
type ShelvesOK struct {

	/*
	  In: Body
	*/
	Payload models.Library `json:"body,omitempty"`
}

// NewShelvesOK creates ShelvesOK with default headers values
func NewShelvesOK() *ShelvesOK {

	return &ShelvesOK{}
}

// WithPayload adds the payload to the shelves o k response
func (o *ShelvesOK) WithPayload(payload models.Library) *ShelvesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the shelves o k response
func (o *ShelvesOK) SetPayload(payload models.Library) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShelvesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Library{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ShelvesBadRequestCode is the HTTP code returned for type ShelvesBadRequest
const ShelvesBadRequestCode int = 400

/*ShelvesBadRequest Bad request

docs:response shelvesBadRequest
*/
type ShelvesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Fail `json:"body,omitempty"`
}

// NewShelvesBadRequest creates ShelvesBadRequest with default headers values
func NewShelvesBadRequest() *ShelvesBadRequest {

	return &ShelvesBadRequest{}
}

// WithPayload adds the payload to the shelves bad request response
func (o *ShelvesBadRequest) WithPayload(payload *models.Fail) *ShelvesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the shelves bad request response
func (o *ShelvesBadRequest) SetPayload(payload *models.Fail) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShelvesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
