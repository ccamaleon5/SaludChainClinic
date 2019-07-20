// Code generated by go-swagger; DO NOT EDIT.

package clinic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/ccamaleon5/SaludChainClient/models"
)

// BookMedicalAppointmentOKCode is the HTTP code returned for type BookMedicalAppointmentOK
const BookMedicalAppointmentOKCode int = 200

/*BookMedicalAppointmentOK successful operation

swagger:response bookMedicalAppointmentOK
*/
type BookMedicalAppointmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.Credential `json:"body,omitempty"`
}

// NewBookMedicalAppointmentOK creates BookMedicalAppointmentOK with default headers values
func NewBookMedicalAppointmentOK() *BookMedicalAppointmentOK {

	return &BookMedicalAppointmentOK{}
}

// WithPayload adds the payload to the book medical appointment o k response
func (o *BookMedicalAppointmentOK) WithPayload(payload *models.Credential) *BookMedicalAppointmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the book medical appointment o k response
func (o *BookMedicalAppointmentOK) SetPayload(payload *models.Credential) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BookMedicalAppointmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BookMedicalAppointmentBadRequestCode is the HTTP code returned for type BookMedicalAppointmentBadRequest
const BookMedicalAppointmentBadRequestCode int = 400

/*BookMedicalAppointmentBadRequest Invalid data supplied

swagger:response bookMedicalAppointmentBadRequest
*/
type BookMedicalAppointmentBadRequest struct {
}

// NewBookMedicalAppointmentBadRequest creates BookMedicalAppointmentBadRequest with default headers values
func NewBookMedicalAppointmentBadRequest() *BookMedicalAppointmentBadRequest {

	return &BookMedicalAppointmentBadRequest{}
}

// WriteResponse to the client
func (o *BookMedicalAppointmentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// BookMedicalAppointmentNotFoundCode is the HTTP code returned for type BookMedicalAppointmentNotFound
const BookMedicalAppointmentNotFoundCode int = 404

/*BookMedicalAppointmentNotFound Resource doesn't exist

swagger:response bookMedicalAppointmentNotFound
*/
type BookMedicalAppointmentNotFound struct {
}

// NewBookMedicalAppointmentNotFound creates BookMedicalAppointmentNotFound with default headers values
func NewBookMedicalAppointmentNotFound() *BookMedicalAppointmentNotFound {

	return &BookMedicalAppointmentNotFound{}
}

// WriteResponse to the client
func (o *BookMedicalAppointmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// BookMedicalAppointmentInternalServerErrorCode is the HTTP code returned for type BookMedicalAppointmentInternalServerError
const BookMedicalAppointmentInternalServerErrorCode int = 500

/*BookMedicalAppointmentInternalServerError Error Internal Server

swagger:response bookMedicalAppointmentInternalServerError
*/
type BookMedicalAppointmentInternalServerError struct {
}

// NewBookMedicalAppointmentInternalServerError creates BookMedicalAppointmentInternalServerError with default headers values
func NewBookMedicalAppointmentInternalServerError() *BookMedicalAppointmentInternalServerError {

	return &BookMedicalAppointmentInternalServerError{}
}

// WriteResponse to the client
func (o *BookMedicalAppointmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
