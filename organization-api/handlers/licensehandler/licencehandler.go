package licensehandler

import (
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"

	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/data/database"
)

// LicenseHandler handler for license
type LicenseHandler struct {
	logger hclog.Logger
	v      *data.Validation
	db     *data.LicenseDBClient
}

// NewLicenseHandler Creates a new License Handler
func NewLicenseHandler(l hclog.Logger, v *data.Validation, db *database.DB, cc grpc.ClientConnInterface) *LicenseHandler {
	return &LicenseHandler{
		logger: l,
		v:      v,
		db:     data.NewLicenseDBClient(l, db, cc),
	}
}
