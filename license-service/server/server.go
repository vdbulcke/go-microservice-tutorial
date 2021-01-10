package server

import (
	context "context"
	licenseprotos "go-microservice-tutorial/license-service/protos/license"

	"github.com/hashicorp/go-hclog"
)

// LicenseServer license server
type LicenseServer struct {
	logger hclog.Logger
	licenseprotos.UnimplementedLicenseServiceServer
}

// NewLicenseServer retruns a new LicenseServer
func NewLicenseServer(l hclog.Logger) *LicenseServer {
	return &LicenseServer{logger: l}
}

// NewLicense generates a new License based on the LicenseRequest
// implements the NewLicense function from the LicenseServiceServer interface
func (ls *LicenseServer) NewLicense(ctx context.Context, lr *licenseprotos.LicenseRequest) (*licenseprotos.License, error) {
	ls.logger.Info("Handle request for GenerateNewLicense", "tenantID", lr.GetTenantID(), "org", lr.GetOrganizationName())

	// create fake License
	// newLicense := &licenseprotos.License{TenantID: lr.GetTenantID(), License: "new-license-generated"}

	// return newLicense, nil

	return &licenseprotos.License{TenantID: lr.GetTenantID(), License: "new-license-generated"}, nil

}
