package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hashicorp/go-hclog"

	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/data/database"
	"go-microservice-tutorial/organization-api/data/database/migration"
	"go-microservice-tutorial/organization-api/handlers/api"
	"go-microservice-tutorial/organization-api/handlers/licensehandler"
	"go-microservice-tutorial/organization-api/handlers/tenantshandler"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var bindAddress = ":9090"

func main() {

	// initializing resources
	l := log.New(os.Stdout, "org-api ", log.LstdFlags)
	v := data.NewValidation()
	db := database.NewSqliteDB("./sqlite.db")
	// new logger
	logger := hclog.Default()

	// grpc client for license service
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// migrate DB & gen data
	migration.DBMigration(db)
	migration.GenerateData(db)

	// create the handlers
	ph := tenantshandler.NewTenants(l, v, db)
	licenseHandler := licensehandler.NewLicenseHandler(logger, v, db, conn)

	// create an api
	apiCommon := api.NewAPI()

	// API Base Path
	apiBasePath := "/api/beta2"

	// create a new serve mux and register the handlers
	mainRouter := mux.NewRouter()

	// handlers for API
	apiRouter := mainRouter.PathPrefix(apiBasePath).Subrouter()
	apiRouter.Use(apiCommon.CommonAPIMiddleware)

	// Tenant API
	getR := apiRouter.Methods(http.MethodGet).Subrouter()
	// getR.HandleFunc("/tenants", api.APIHandler{ ph.ListAll})
	// getR.HandleFunc("/tenants/{id}", api.APIHandler{ ph.ListSingle})
	getR.Handle("/tenants", api.APIHandler{Handler: ph.ListAll})
	getR.Handle("/tenants/{id}", api.APIHandler{Handler: ph.ListSingle})

	putR := apiRouter.Methods(http.MethodPut).Subrouter()
	// putR.HandleFunc("/tenants", ph.Update)
	putR.Handle("/tenants", api.APIHandler{Handler: ph.Update})
	putR.Use(ph.MiddlewareValidateTenantUpdate)

	postR := apiRouter.Methods(http.MethodPost).Subrouter()
	// postR.HandleFunc("/tenants", ph.Create)
	postR.Handle("/tenants", api.APIHandler{Handler: ph.Create})
	postR.Use(ph.MiddlewareValidateTenantCreate)

	deleteR := apiRouter.Methods(http.MethodDelete).Subrouter()
	// deleteR.HandleFunc("/tenants/{id:[0-9]+}", api.APIHandler(ph.Delete))
	deleteR.Handle("/tenants/{id}", api.APIHandler{Handler: ph.Delete})

	// License API
	getR.Handle("/license/get_license_by_id/{id}", api.APIHandler{Handler: licenseHandler.GetLicenseByID})
	getR.Handle("/license/get_licenses_by_tenant_id/{id}", api.APIHandler{Handler: licenseHandler.GetLicensesByTenantID})
	getR.Handle("/license/generate_license_for_tenant_id/{id}", api.APIHandler{Handler: licenseHandler.GenerateLicenceForTenant})

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	// Doc Handers
	getDoc := mainRouter.Methods(http.MethodGet).Subrouter()
	getDoc.Handle("/docs", sh)
	getDoc.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS

	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := http.Server{
		Addr:         bindAddress,             // configure the bind address
		Handler:      corsHandler(mainRouter), // set the default handler
		ErrorLog:     l,                       // set the logger for the server
		ReadTimeout:  5 * time.Second,         // max time to read request from the client
		WriteTimeout: 10 * time.Second,        // max time to write response to the client
		IdleTimeout:  120 * time.Second,       // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
