// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/ccamaleon5/SaludChainClient/restapi/operations"
	"github.com/ccamaleon5/SaludChainClient/restapi/operations/clinic"
	"github.com/ccamaleon5/SaludChainClient/restapi/operations/health"
	"github.com/ccamaleon5/SaludChainClient/restapi/operations/insurance"
	"github.com/ccamaleon5/SaludChainClient/restapi/operations/pharmacy"

	"github.com/ccamaleon5/SaludChainClient/business"

	"github.com/rs/cors"
)

//go:generate swagger generate server --target ../../SaludChainClient --name SaludChainAPIs --spec ../swagger/swaggerui/swagger.json

func configureFlags(api *operations.SaludChainAPIsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SaludChainAPIsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.HealthRemoveRecordSharedHandler == nil {
		api.HealthRemoveRecordSharedHandler = health.RemoveRecordSharedHandlerFunc(func(params health.RemoveRecordSharedParams) middleware.Responder {
			return middleware.NotImplemented("operation health.RemoveRecordShared has not yet been implemented")
		})
	}
	if api.ClinicBookMedicalAppointmentHandler == nil {
		api.ClinicBookMedicalAppointmentHandler = clinic.BookMedicalAppointmentHandlerFunc(func(params clinic.BookMedicalAppointmentParams) middleware.Responder {
			return middleware.NotImplemented("operation clinic.BookMedicalAppointment has not yet been implemented")
		})
	}

	api.ClinicSetMedicalAttentionHandler = clinic.SetMedicalAttentionHandlerFunc(func(params clinic.SetMedicalAttentionParams) middleware.Responder {
		_, err := business.AttendPatient(params.Body.PublicKey, params.Body.MedicalAppointmentID, params.Body.RecordMedical)
		if err != nil {
			return clinic.NewSetMedicalAttentionBadRequest()
		}

		return clinic.NewSetMedicalAttentionOK()
	})

	api.ClinicBookNewMedicalAppointmentHandler = clinic.BookNewMedicalAppointmentHandlerFunc(func(params clinic.BookNewMedicalAppointmentParams) middleware.Responder {
		_, err := business.NewMedicalAppointment(params.Body.PatientID, params.Body.DoctorID, params.Body.PubKey, params.Body.Comments)
		if err != nil {
			return clinic.NewBookNewMedicalAppointmentBadRequest()
		}

		return clinic.NewBookNewMedicalAppointmentOK()
	})

		api.InsuranceCreateHealthCredentialHandler = insurance.CreateHealthCredentialHandlerFunc(func(params insurance.CreateHealthCredentialParams) middleware.Responder {
			response, err := business.CreatePatient(params.PublicKey, params.IDAccount, params.Body)
			if err != nil {
				return insurance.NewCreateHealthCredentialBadRequest()
			}
			return insurance.NewCreateHealthCredentialOK().WithPayload(response)
		})

		api.InsuranceCreateDoctorCredentialHandler = insurance.CreateDoctorCredentialHandlerFunc(func(params insurance.CreateDoctorCredentialParams) middleware.Responder {
			response, err := business.CreateDoctor(params.PublicKey, params.IDAccount, params.Body)
			if err != nil {
				return insurance.NewCreateDoctorCredentialBadRequest()
			}
			return insurance.NewCreateDoctorCredentialOK().WithPayload(response)
		})
	
	if api.InsuranceGetHealthCredentialByIDHandler == nil {
		api.InsuranceGetHealthCredentialByIDHandler = insurance.GetHealthCredentialByIDHandlerFunc(func(params insurance.GetHealthCredentialByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation insurance.GetHealthCredentialByID has not yet been implemented")
		})
	}
	if api.InsuranceRevokeCredentialHandler == nil {
		api.InsuranceRevokeCredentialHandler = insurance.RevokeCredentialHandlerFunc(func(params insurance.RevokeCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation insurance.RevokeCredential has not yet been implemented")
		})
	}
	if api.ClinicSetMedicalAttentionHandler == nil {
		api.ClinicSetMedicalAttentionHandler = clinic.SetMedicalAttentionHandlerFunc(func(params clinic.SetMedicalAttentionParams) middleware.Responder {
			return middleware.NotImplemented("operation clinic.SetMedicalAttention has not yet been implemented")
		})
	}
	if api.HealthShareMedicalRecordHandler == nil {
		api.HealthShareMedicalRecordHandler = health.ShareMedicalRecordHandlerFunc(func(params health.ShareMedicalRecordParams) middleware.Responder {
			return middleware.NotImplemented("operation health.ShareMedicalRecord has not yet been implemented")
		})
	}

	api.HealthCreateAccountHandler = health.CreateAccountHandlerFunc(func(params health.CreateAccountParams) middleware.Responder {
		response, err := business.CreateAccount(params.PublicKey)
		if err != nil {
			return health.NewCreateAccountBadRequest()
		}
		responseBody := &health.CreateAccountOKBody{ID: response}

		return health.NewCreateAccountOK().WithPayload(responseBody)
	})

	

	if api.InsuranceVerifyHealthCredentialHandler == nil {
		api.InsuranceVerifyHealthCredentialHandler = insurance.VerifyHealthCredentialHandlerFunc(func(params insurance.VerifyHealthCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation insurance.VerifyHealthCredential has not yet been implemented")
		})
	}
	if api.PharmacyVerifyRecipeHandler == nil {
		api.PharmacyVerifyRecipeHandler = pharmacy.VerifyRecipeHandlerFunc(func(params pharmacy.VerifyRecipeParams) middleware.Responder {
			return middleware.NotImplemented("operation pharmacy.VerifyRecipe has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {

	corsHandler := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{},
		MaxAge:         1000,
	})
	return corsHandler.Handler(uiMiddleware(handler))
}

func uiMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Shortcut helpers for swagger-ui
		if r.URL.Path == "/swagger-ui" || r.URL.Path == "/api/help" {
			http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
			return
		}
		// Serving ./swagger-ui/
		if strings.Index(r.URL.Path, "/swagger-ui/") == 0 {
			http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swaggerui/"))).ServeHTTP(w, r)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
