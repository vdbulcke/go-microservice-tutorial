package migration

import (
	"fmt"
	tenantdata "go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/data/database"
	"log"
)

// DBMigration initialize DB
func DBMigration(db *database.DB) {

	// make sure Tenant is create in DB
	db.Client.AutoMigrate(&tenantdata.Tenant{})

}

// GenerateData insert some data in DB
func GenerateData(db *database.DB) {

	var tenantList = []*tenantdata.Tenant{
		&tenantdata.Tenant{
			Name:        "ACME",
			Description: "ACME Corp",
		},
	}

	// insert in DB
	for _, t := range tenantList {
		var temp tenantdata.Tenant

		searchres := db.Client.Where("name = ?", t.Name).First(&temp)
		if searchres.Error == nil {
			log.Println("Record already exists %v", t)
			continue
		}

		result := db.Client.Create(t)
		if result.Error != nil {
			log.Println(result.Error)
		}

		fmt.Println("Inserted %v", t)
	}
}
