// Code generated by go-swagger; DO NOT EDIT.

package bookcases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MarlikAlmighty/library/pkg/models"
)

// ListAllBooksFromBookcasesReader is a Reader for the ListAllBooksFromBookcases structure.
type ListAllBooksFromBookcasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllBooksFromBookcasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllBooksFromBookcasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAllBooksFromBookcasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllBooksFromBookcasesOK creates a ListAllBooksFromBookcasesOK with default headers values
func NewListAllBooksFromBookcasesOK() *ListAllBooksFromBookcasesOK {
	return &ListAllBooksFromBookcasesOK{}
}

/* ListAllBooksFromBookcasesOK describes a response with status code 200, with default header values.

OK
*/
type ListAllBooksFromBookcasesOK struct {
	Payload models.Books
}

func (o *ListAllBooksFromBookcasesOK) Error() string {
	return fmt.Sprintf("[GET /bookcases/{id}][%d] listAllBooksFromBookcasesOK  %+v", 200, o.Payload)
}
func (o *ListAllBooksFromBookcasesOK) GetPayload() models.Books {
	return o.Payload
}

func (o *ListAllBooksFromBookcasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllBooksFromBookcasesBadRequest creates a ListAllBooksFromBookcasesBadRequest with default headers values
func NewListAllBooksFromBookcasesBadRequest() *ListAllBooksFromBookcasesBadRequest {
	return &ListAllBooksFromBookcasesBadRequest{}
}

/* ListAllBooksFromBookcasesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListAllBooksFromBookcasesBadRequest struct {
	Payload *models.Fail
}

func (o *ListAllBooksFromBookcasesBadRequest) Error() string {
	return fmt.Sprintf("[GET /bookcases/{id}][%d] listAllBooksFromBookcasesBadRequest  %+v", 400, o.Payload)
}
func (o *ListAllBooksFromBookcasesBadRequest) GetPayload() *models.Fail {
	return o.Payload
}

func (o *ListAllBooksFromBookcasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
