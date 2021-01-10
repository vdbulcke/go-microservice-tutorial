package data

import (
	"errors"
	"fmt"
	"go-microservice-tutorial/organization-api/data/database"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// Tenant defines the structure for an API tenant
// swagger:model
type Tenant struct {
	// the id for the tenant
	//
	// required: false
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`

	// the name for this tenant
	//
	// required: true
	// max length: 255
	Name string `json:"name" gorm:"uniqueIndex" validate:"required"`

	// the description for this tenant
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the CreatedAt timestamp for the tenant
	//
	// required: false
	CreatedAt time.Time `json:"-" `

	// the UpdatedAt timestamp for the tenant
	//
	// required: false
	UpdatedAt time.Time `json:"-" `
}

// TenantCreate defines the structure for an API tenant create object
// swagger:model
type TenantCreate struct {

	// the name for this tenant
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this tenant
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`
}

// TenantUpdate defines the structure for an API tenant update object
// swagger:model
type TenantUpdate struct {
	// the id for the tenant
	//
	// required: true
	ID uuid.UUID `json:"id"  validate:"required"`

	// the name for this tenant
	//
	// required: true
	// max length: 255
	Name string `json:"name"  validate:"required"`

	// the description for this tenant
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (tenant *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tenant.ID = uuid

	return
}

// Tenants defines a slice of Tenant
type Tenants []*Tenant

// TenantNotFound error when tenant not found
type TenantNotFound struct{}

// ErrTenantNotFound is an error raised when a tenant can not be found in the database
var ErrTenantNotFound = fmt.Errorf("Tenant not found")

// TenantAlreadyExist error when tenant not found
type TenantAlreadyExist struct {
	Err     error
	FoundID string
}

func (e *TenantAlreadyExist) Error() string {
	return fmt.Sprintf("Error TenantAlreadyExist ID: '%s'", e.FoundID)
}

// GetTenants returns all tenants from the database
func GetTenants(db *database.DB) (Tenants, error) {

	var tenantList Tenants

	result := db.Client.Find(&tenantList)
	if result.Error != nil {
		return nil, result.Error
	}

	return tenantList, nil
}

// GetTenantByID returns a single tenant which matches the id from the
// database.
// If a tenant is not found this function returns a TenantNotFound error
func GetTenantByID(id uuid.UUID, db *database.DB) (*Tenant, error) {

	var tenant Tenant

	result := db.Client.First(&tenant, id)
	if result.Error != nil {
		return nil, ErrTenantNotFound
	}

	// check if result is empty
	if result.RowsAffected == 0 {
		return nil, ErrTenantNotFound
	}

	return &tenant, nil
}

// UpdateTenant replaces a tenant in the database with the given
// item.
// If a tenant with the given id does not exist in the database
// this function returns a TenantNotFound error
func UpdateTenant(p Tenant, db *database.DB) error {

	var tenant Tenant

	result := db.Client.Find(&tenant, p.ID)
	if result.Error != nil {
		return ErrTenantNotFound
	}

	// update fields
	tenant.Name = p.Name
	tenant.Description = p.Description

	updateresult := db.Client.Save(&tenant)
	if updateresult.Error != nil {
		return ErrTenantNotFound
	}

	return nil
}

// AddTenant adds a new tenant to the database
func AddTenant(t Tenant, db *database.DB) (*Tenant, error) {

	var temp Tenant

	searchres := db.Client.Where("name = ?", t.Name).First(&temp)
	if searchres.Error == nil {
		// setting tenant ID to object found in DB
		// for error handling
		fmt.Println("found %v", temp)

		return &temp, &TenantAlreadyExist{Err: errors.New("error creating tenant"), FoundID: temp.ID.String()}
	}

	// Create new entry
	result := db.Client.Create(&t)
	if result.Error != nil {
		// log.Fatal(result.Error)
		return nil, result.Error
	}
	return &t, nil
}

// DeleteTenant deletes a tenant from the database
func DeleteTenant(id uuid.UUID, db *database.DB) error {

	result := db.Client.Delete(&Tenant{}, id)
	if result.Error != nil {
		return ErrTenantNotFound
	}

	return nil
}

// findIndex finds the index of a tenant in the database
// returns -1 when no tenant can be found
// func findIndexByTenantID(id uint) (int, *TenantNotFound) {
// 	for i, p := range tenantList {
// 		if p.ID == id {
// 			return i, nil
// 		}
// 	}

// 	return 0, &TenantNotFound{}
// }
