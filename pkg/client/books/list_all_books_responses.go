// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MarlikAlmighty/library/pkg/models"
)

// ListAllBooksReader is a Reader for the ListAllBooks structure.
type ListAllBooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllBooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllBooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAllBooksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllBooksOK creates a ListAllBooksOK with default headers values
func NewListAllBooksOK() *ListAllBooksOK {
	return &ListAllBooksOK{}
}

/* ListAllBooksOK describes a response with status code 200, with default header values.

OK
*/
type ListAllBooksOK struct {
	Payload models.Books
}

func (o *ListAllBooksOK) Error() string {
	return fmt.Sprintf("[GET /books][%d] listAllBooksOK  %+v", 200, o.Payload)
}
func (o *ListAllBooksOK) GetPayload() models.Books {
	return o.Payload
}

func (o *ListAllBooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllBooksBadRequest creates a ListAllBooksBadRequest with default headers values
func NewListAllBooksBadRequest() *ListAllBooksBadRequest {
	return &ListAllBooksBadRequest{}
}

/* ListAllBooksBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListAllBooksBadRequest struct {
	Payload *models.Fail
}

func (o *ListAllBooksBadRequest) Error() string {
	return fmt.Sprintf("[GET /books][%d] listAllBooksBadRequest  %+v", 400, o.Payload)
}
func (o *ListAllBooksBadRequest) GetPayload() *models.Fail {
	return o.Payload
}

func (o *ListAllBooksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
