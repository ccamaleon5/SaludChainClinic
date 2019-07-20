// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/ccamaleon5/SaludChainClient/models"
)

// RemoveRecordSharedOKCode is the HTTP code returned for type RemoveRecordSharedOK
const RemoveRecordSharedOKCode int = 200

/*RemoveRecordSharedOK successful operation

swagger:response removeRecordSharedOK
*/
type RemoveRecordSharedOK struct {

	/*
	  In: Body
	*/
	Payload *models.Credential `json:"body,omitempty"`
}

// NewRemoveRecordSharedOK creates RemoveRecordSharedOK with default headers values
func NewRemoveRecordSharedOK() *RemoveRecordSharedOK {

	return &RemoveRecordSharedOK{}
}

// WithPayload adds the payload to the remove record shared o k response
func (o *RemoveRecordSharedOK) WithPayload(payload *models.Credential) *RemoveRecordSharedOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove record shared o k response
func (o *RemoveRecordSharedOK) SetPayload(payload *models.Credential) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveRecordSharedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveRecordSharedBadRequestCode is the HTTP code returned for type RemoveRecordSharedBadRequest
const RemoveRecordSharedBadRequestCode int = 400

/*RemoveRecordSharedBadRequest Invalid data supplied

swagger:response removeRecordSharedBadRequest
*/
type RemoveRecordSharedBadRequest struct {
}

// NewRemoveRecordSharedBadRequest creates RemoveRecordSharedBadRequest with default headers values
func NewRemoveRecordSharedBadRequest() *RemoveRecordSharedBadRequest {

	return &RemoveRecordSharedBadRequest{}
}

// WriteResponse to the client
func (o *RemoveRecordSharedBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// RemoveRecordSharedNotFoundCode is the HTTP code returned for type RemoveRecordSharedNotFound
const RemoveRecordSharedNotFoundCode int = 404

/*RemoveRecordSharedNotFound Resource doesn't exist

swagger:response removeRecordSharedNotFound
*/
type RemoveRecordSharedNotFound struct {
}

// NewRemoveRecordSharedNotFound creates RemoveRecordSharedNotFound with default headers values
func NewRemoveRecordSharedNotFound() *RemoveRecordSharedNotFound {

	return &RemoveRecordSharedNotFound{}
}

// WriteResponse to the client
func (o *RemoveRecordSharedNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// RemoveRecordSharedInternalServerErrorCode is the HTTP code returned for type RemoveRecordSharedInternalServerError
const RemoveRecordSharedInternalServerErrorCode int = 500

/*RemoveRecordSharedInternalServerError Error Internal Server

swagger:response removeRecordSharedInternalServerError
*/
type RemoveRecordSharedInternalServerError struct {
}

// NewRemoveRecordSharedInternalServerError creates RemoveRecordSharedInternalServerError with default headers values
func NewRemoveRecordSharedInternalServerError() *RemoveRecordSharedInternalServerError {

	return &RemoveRecordSharedInternalServerError{}
}

// WriteResponse to the client
func (o *RemoveRecordSharedInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
