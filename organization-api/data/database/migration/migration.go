package migration

import (
	"fmt"
	"go-microservice-tutorial/organization-api/data"
	"go-microservice-tutorial/organization-api/data/database"
	"log"
)

// DBMigration initialize DB
func DBMigration(db *database.DB) {

	// make sure Tenant is create in DB
	db.Client.AutoMigrate(&data.Tenant{})
	// make sure Tenant is create in DB
	db.Client.AutoMigrate(&data.License{})
}

// GenerateData insert some data in DB
func GenerateData(db *database.DB) {

	var tenantList = []*data.Tenant{
		&data.Tenant{
			Name:        "ACME",
			Description: "ACME Corp",
		},
	}

	// insert in DB
	for _, t := range tenantList {
		var temp data.Tenant

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
