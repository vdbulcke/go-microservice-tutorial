package data

import (
	"context"
	"errors"
	"go-microservice-tutorial/license-service/protos/license"
	licenseprotos "go-microservice-tutorial/license-service/protos/license"
	"go-microservice-tutorial/organization-api/data/database"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// License defines the structure for an API tenant
// swagger:model
type License struct {
	// the id for the license
	//
	// required: false
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`

	// the License for the license
	//
	// required: true
	License string `json:"license" validate:"required"`

	// the TenantID this license belongs to
	//
	// required: true
	TenantID uuid.UUID `json:"tenant_id"  gorm:"column:tenant_id" validate:"required"`

	// the Tenant this license belongs to
	// (Foreign key)
	//
	// required: false
	Tenant Tenant `json:"-"`

	// the CreatedAt timestamp for the tenant
	//
	// required: false
	CreatedAt time.Time `json:"-" `

	// the UpdatedAt timestamp for the tenant
	//
	// required: false
	UpdatedAt time.Time `json:"-" `
}

// Licenses a slice of license
type Licenses []*License

// BeforeCreate will set a UUID rather than numeric ID.
func (l *License) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	l.ID = uuid

	return
}

// LicenseDBClient Client used by the handler
// for handling operation with license
type LicenseDBClient struct {
	logger     hclog.Logger
	db         *database.DB
	grpcClient licenseprotos.LicenseServiceClient
}

// NewLicenseDBClient create a new LicenseDBClient
func NewLicenseDBClient(l hclog.Logger, db *database.DB, cc grpc.ClientConnInterface) *LicenseDBClient {
	return &LicenseDBClient{
		logger:     l,
		db:         db,
		grpcClient: license.NewLicenseServiceClient(cc),
	}
}

// CreateLicenseForTenant creates a license
func (licensedbclient *LicenseDBClient) CreateLicenseForTenant(tenantID uuid.UUID) (*License, error) {

	// get tenant ID from tenant data
	tenant, tenantErr := GetTenantByID(tenantID, licensedbclient.db)
	if tenantErr != nil {
		return nil, tenantErr

	}

	licensedbclient.logger.Info(" found tenant", tenant)

	tenantIDStr := tenant.ID.String()
	orgName := tenant.Name

	// Get New License
	newlicense, err := licensedbclient.createNewLicenceForTenant(tenantIDStr, orgName)
	if err != nil {
		licensedbclient.logger.Error("getting new license ", err)
	}

	// Create a new License in DB
	licenseObj := &License{
		TenantID: tenantID,
		// Tenant:   *tenant,
		License: newlicense,
	}

	// Create new entry
	result := licensedbclient.db.Client.Omit("Tenant").Create(&licenseObj)
	if result.Error != nil {
		licensedbclient.logger.Error("error inserting license in DB", result.Error)
		return nil, result.Error
	}

	return licenseObj, nil

}

// createNewLicenceForTenant call the License Service (via grpc) and
// return a new license
func (licensedbclient *LicenseDBClient) createNewLicenceForTenant(tenantID string, orgName string) (string, error) {

	// Create a new License Request
	newLicenseRequest := &licenseprotos.LicenseRequest{
		TenantID:         tenantID,
		OrganizationName: orgName,
	}

	// use grpc client to get a new license
	newlicense, err := licensedbclient.grpcClient.NewLicense(context.Background(), newLicenseRequest)
	if err != nil {
		return "", err
	}

	// return new license
	return newlicense.License, err
}

// GetLicenceByUUID get license from DB
// return a new license
func (licensedbclient *LicenseDBClient) GetLicenceByUUID(id uuid.UUID) (*License, error) {

	var license License

	result := licensedbclient.db.Client.First(&license, id)
	if result.Error != nil {
		return nil, errors.New("License Not Found " + id.String())
	}

	// check if result is empty
	if result.RowsAffected == 0 {
		return nil, errors.New("License Not Found " + id.String())
	}

	return &license, nil

}

// GetLicensesByTenantID get all licenses from DB for a given tenant
// return a new license
func (licensedbclient *LicenseDBClient) GetLicensesByTenantID(id uuid.UUID) (*Licenses, error) {

	var licenses Licenses

	result := licensedbclient.db.Client.Where("tenant_id = ?", id).Find(&licenses)
	if result.Error != nil {
		return nil, errors.New("Licenses Not Found for tenant id " + id.String())
	}

	// check if result is empty
	if result.RowsAffected == 0 {
		return nil, errors.New("License Not Found " + id.String())
	}

	return &licenses, nil

}
