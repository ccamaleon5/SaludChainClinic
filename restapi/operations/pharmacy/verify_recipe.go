// Code generated by go-swagger; DO NOT EDIT.

package pharmacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// VerifyRecipeHandlerFunc turns a function with the right signature into a verify recipe handler
type VerifyRecipeHandlerFunc func(VerifyRecipeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyRecipeHandlerFunc) Handle(params VerifyRecipeParams) middleware.Responder {
	return fn(params)
}

// VerifyRecipeHandler interface for that can handle valid verify recipe params
type VerifyRecipeHandler interface {
	Handle(VerifyRecipeParams) middleware.Responder
}

// NewVerifyRecipe creates a new http.Handler for the verify recipe operation
func NewVerifyRecipe(ctx *middleware.Context, handler VerifyRecipeHandler) *VerifyRecipe {
	return &VerifyRecipe{Context: ctx, Handler: handler}
}

/*VerifyRecipe swagger:route GET /pharmacy/recipe/{hashRecipe}/verify pharmacy verifyRecipe

Verify Patient's recipe

This service is responsible for verify a recipe given by Doctor in Pharmacy

*/
type VerifyRecipe struct {
	Context *middleware.Context
	Handler VerifyRecipeHandler
}

func (o *VerifyRecipe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerifyRecipeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
