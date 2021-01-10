module go-microservice-tutorial/organization-api

go 1.15

require (
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/runtime v0.19.24
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.10
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/google/uuid v1.1.2
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-hclog v0.15.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/stretchr/testify v1.6.1
	go-microservice-tutorial/license-service v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.34.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.9
)

replace go-microservice-tutorial/license-service => ../license-service
