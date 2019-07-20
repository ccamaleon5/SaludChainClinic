// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RemoveRecordSharedHandlerFunc turns a function with the right signature into a remove record shared handler
type RemoveRecordSharedHandlerFunc func(RemoveRecordSharedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveRecordSharedHandlerFunc) Handle(params RemoveRecordSharedParams) middleware.Responder {
	return fn(params)
}

// RemoveRecordSharedHandler interface for that can handle valid remove record shared params
type RemoveRecordSharedHandler interface {
	Handle(RemoveRecordSharedParams) middleware.Responder
}

// NewRemoveRecordShared creates a new http.Handler for the remove record shared operation
func NewRemoveRecordShared(ctx *middleware.Context, handler RemoveRecordSharedHandler) *RemoveRecordShared {
	return &RemoveRecordShared{Context: ctx, Handler: handler}
}

/*RemoveRecordShared swagger:route DELETE /health/patient/medicalrecord/share health removeRecordShared

Revoke sharing medical information

This service is responsible for remove a medical record shared

*/
type RemoveRecordShared struct {
	Context *middleware.Context
	Handler RemoveRecordSharedHandler
}

func (o *RemoveRecordShared) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveRecordSharedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
