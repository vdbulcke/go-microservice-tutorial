package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTenantMissingNameReturnsErr(t *testing.T) {
	tenant := Tenant{
		Description: "FOD BOSA",
	}

	v := NewValidation()
	err := v.Validate(tenant)
	assert.Len(t, err, 1)
}

func TestTenantInvalidKBO(t *testing.T) {
	tenant := Tenant{
		Name:        "BOSA",
		Description: "FOD BOSA",
	}

	v := NewValidation()
	err := v.Validate(tenant)
	assert.Len(t, err, 1)

}
func TestTenantvalid(t *testing.T) {
	tenant := Tenant{
		Name:        "BOSA",
		Description: "FOD BOSA",
	}

	v := NewValidation()
	err := v.Validate(tenant)
	assert.Len(t, err, 0)

}
