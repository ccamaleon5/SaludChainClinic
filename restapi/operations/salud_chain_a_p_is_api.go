// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ccamaleon5/SaludChainClient/restapi/operations/clinic"
	"github.com/ccamaleon5/SaludChainClient/restapi/operations/health"
	"github.com/ccamaleon5/SaludChainClient/restapi/operations/insurance"
	"github.com/ccamaleon5/SaludChainClient/restapi/operations/pharmacy"
)

// NewSaludChainAPIsAPI creates a new SaludChainAPIs instance
func NewSaludChainAPIsAPI(spec *loads.Document) *SaludChainAPIsAPI {
	return &SaludChainAPIsAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		HealthRemoveRecordSharedHandler: health.RemoveRecordSharedHandlerFunc(func(params health.RemoveRecordSharedParams) middleware.Responder {
			return middleware.NotImplemented("operation HealthRemoveRecordShared has not yet been implemented")
		}),
		ClinicBookMedicalAppointmentHandler: clinic.BookMedicalAppointmentHandlerFunc(func(params clinic.BookMedicalAppointmentParams) middleware.Responder {
			return middleware.NotImplemented("operation ClinicBookMedicalAppointment has not yet been implemented")
		}),
		ClinicBookNewMedicalAppointmentHandler: clinic.BookNewMedicalAppointmentHandlerFunc(func(params clinic.BookNewMedicalAppointmentParams) middleware.Responder {
			return middleware.NotImplemented("operation ClinicBookNewMedicalAppointment has not yet been implemented")
		}),
		HealthCreateAccountHandler: health.CreateAccountHandlerFunc(func(params health.CreateAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation HealthCreateAccount has not yet been implemented")
		}),
		InsuranceCreateDoctorCredentialHandler: insurance.CreateDoctorCredentialHandlerFunc(func(params insurance.CreateDoctorCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation InsuranceCreateDoctorCredential has not yet been implemented")
		}),
		InsuranceCreateHealthCredentialHandler: insurance.CreateHealthCredentialHandlerFunc(func(params insurance.CreateHealthCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation InsuranceCreateHealthCredential has not yet been implemented")
		}),
		InsuranceGetHealthCredentialByIDHandler: insurance.GetHealthCredentialByIDHandlerFunc(func(params insurance.GetHealthCredentialByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation InsuranceGetHealthCredentialByID has not yet been implemented")
		}),
		InsuranceRevokeCredentialHandler: insurance.RevokeCredentialHandlerFunc(func(params insurance.RevokeCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation InsuranceRevokeCredential has not yet been implemented")
		}),
		ClinicSetMedicalAttentionHandler: clinic.SetMedicalAttentionHandlerFunc(func(params clinic.SetMedicalAttentionParams) middleware.Responder {
			return middleware.NotImplemented("operation ClinicSetMedicalAttention has not yet been implemented")
		}),
		HealthShareMedicalRecordHandler: health.ShareMedicalRecordHandlerFunc(func(params health.ShareMedicalRecordParams) middleware.Responder {
			return middleware.NotImplemented("operation HealthShareMedicalRecord has not yet been implemented")
		}),
		InsuranceVerifyHealthCredentialHandler: insurance.VerifyHealthCredentialHandlerFunc(func(params insurance.VerifyHealthCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation InsuranceVerifyHealthCredential has not yet been implemented")
		}),
		PharmacyVerifyRecipeHandler: pharmacy.VerifyRecipeHandlerFunc(func(params pharmacy.VerifyRecipeParams) middleware.Responder {
			return middleware.NotImplemented("operation PharmacyVerifyRecipe has not yet been implemented")
		}),
	}
}

/*SaludChainAPIsAPI This is a SaludChain Client that send transactions to SaludChain Blockchain

SaludChain Client exposes APIs to be consume by patients, clinics, insurances, pharmacies and health institutions  */
type SaludChainAPIsAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// HealthRemoveRecordSharedHandler sets the operation handler for the remove record shared operation
	HealthRemoveRecordSharedHandler health.RemoveRecordSharedHandler
	// ClinicBookMedicalAppointmentHandler sets the operation handler for the book medical appointment operation
	ClinicBookMedicalAppointmentHandler clinic.BookMedicalAppointmentHandler
	// ClinicBookNewMedicalAppointmentHandler sets the operation handler for the book new medical appointment operation
	ClinicBookNewMedicalAppointmentHandler clinic.BookNewMedicalAppointmentHandler
	// HealthCreateAccountHandler sets the operation handler for the create account operation
	HealthCreateAccountHandler health.CreateAccountHandler
	// InsuranceCreateDoctorCredentialHandler sets the operation handler for the create doctor credential operation
	InsuranceCreateDoctorCredentialHandler insurance.CreateDoctorCredentialHandler
	// InsuranceCreateHealthCredentialHandler sets the operation handler for the create health credential operation
	InsuranceCreateHealthCredentialHandler insurance.CreateHealthCredentialHandler
	// InsuranceGetHealthCredentialByIDHandler sets the operation handler for the get health credential by Id operation
	InsuranceGetHealthCredentialByIDHandler insurance.GetHealthCredentialByIDHandler
	// InsuranceRevokeCredentialHandler sets the operation handler for the revoke credential operation
	InsuranceRevokeCredentialHandler insurance.RevokeCredentialHandler
	// ClinicSetMedicalAttentionHandler sets the operation handler for the set medical attention operation
	ClinicSetMedicalAttentionHandler clinic.SetMedicalAttentionHandler
	// HealthShareMedicalRecordHandler sets the operation handler for the share medical record operation
	HealthShareMedicalRecordHandler health.ShareMedicalRecordHandler
	// InsuranceVerifyHealthCredentialHandler sets the operation handler for the verify health credential operation
	InsuranceVerifyHealthCredentialHandler insurance.VerifyHealthCredentialHandler
	// PharmacyVerifyRecipeHandler sets the operation handler for the verify recipe operation
	PharmacyVerifyRecipeHandler pharmacy.VerifyRecipeHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *SaludChainAPIsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SaludChainAPIsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SaludChainAPIsAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SaludChainAPIsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SaludChainAPIsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SaludChainAPIsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SaludChainAPIsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SaludChainAPIsAPI
func (o *SaludChainAPIsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.HealthRemoveRecordSharedHandler == nil {
		unregistered = append(unregistered, "health.RemoveRecordSharedHandler")
	}

	if o.ClinicBookMedicalAppointmentHandler == nil {
		unregistered = append(unregistered, "clinic.BookMedicalAppointmentHandler")
	}

	if o.ClinicBookNewMedicalAppointmentHandler == nil {
		unregistered = append(unregistered, "clinic.BookNewMedicalAppointmentHandler")
	}

	if o.HealthCreateAccountHandler == nil {
		unregistered = append(unregistered, "health.CreateAccountHandler")
	}

	if o.InsuranceCreateDoctorCredentialHandler == nil {
		unregistered = append(unregistered, "insurance.CreateDoctorCredentialHandler")
	}

	if o.InsuranceCreateHealthCredentialHandler == nil {
		unregistered = append(unregistered, "insurance.CreateHealthCredentialHandler")
	}

	if o.InsuranceGetHealthCredentialByIDHandler == nil {
		unregistered = append(unregistered, "insurance.GetHealthCredentialByIDHandler")
	}

	if o.InsuranceRevokeCredentialHandler == nil {
		unregistered = append(unregistered, "insurance.RevokeCredentialHandler")
	}

	if o.ClinicSetMedicalAttentionHandler == nil {
		unregistered = append(unregistered, "clinic.SetMedicalAttentionHandler")
	}

	if o.HealthShareMedicalRecordHandler == nil {
		unregistered = append(unregistered, "health.ShareMedicalRecordHandler")
	}

	if o.InsuranceVerifyHealthCredentialHandler == nil {
		unregistered = append(unregistered, "insurance.VerifyHealthCredentialHandler")
	}

	if o.PharmacyVerifyRecipeHandler == nil {
		unregistered = append(unregistered, "pharmacy.VerifyRecipeHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SaludChainAPIsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SaludChainAPIsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *SaludChainAPIsAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *SaludChainAPIsAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *SaludChainAPIsAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SaludChainAPIsAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the salud chain a p is API
func (o *SaludChainAPIsAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SaludChainAPIsAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/health/patient/medicalrecord/share"] = health.NewRemoveRecordShared(o.context, o.HealthRemoveRecordSharedHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clinic/medicalappointment"] = clinic.NewBookMedicalAppointment(o.context, o.ClinicBookMedicalAppointmentHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clinic/medicalappointment/create"] = clinic.NewBookNewMedicalAppointment(o.context, o.ClinicBookNewMedicalAppointmentHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health/account/{publicKey}"] = health.NewCreateAccount(o.context, o.HealthCreateAccountHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/insurance/credential/doctor"] = insurance.NewCreateDoctorCredential(o.context, o.InsuranceCreateDoctorCredentialHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/insurance/credential"] = insurance.NewCreateHealthCredential(o.context, o.InsuranceCreateHealthCredentialHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/insurance/credential/{publicKey}"] = insurance.NewGetHealthCredentialByID(o.context, o.InsuranceGetHealthCredentialByIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/insurance/credential/{publicKey}"] = insurance.NewRevokeCredential(o.context, o.InsuranceRevokeCredentialHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/clinic/medicalappointment"] = clinic.NewSetMedicalAttention(o.context, o.ClinicSetMedicalAttentionHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/health/patient/medicalrecord/share"] = health.NewShareMedicalRecord(o.context, o.HealthShareMedicalRecordHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/insurance/credential/{hashCredential}/verify"] = insurance.NewVerifyHealthCredential(o.context, o.InsuranceVerifyHealthCredentialHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pharmacy/recipe/{hashRecipe}/verify"] = pharmacy.NewVerifyRecipe(o.context, o.PharmacyVerifyRecipeHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SaludChainAPIsAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SaludChainAPIsAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SaludChainAPIsAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SaludChainAPIsAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
