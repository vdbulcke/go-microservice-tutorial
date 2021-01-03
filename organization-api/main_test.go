package main

import (
	"fmt"
	"go-microservice-tutorial/organization-api/sdk/client"
	"go-microservice-tutorial/organization-api/sdk/client/tenants"
	"testing"
)

func TestClient(t *testing.T) {
	transportCfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	cli := client.NewHTTPClientWithConfig(nil, transportCfg)

	listTenantReqParam := tenants.NewListTenantsParams()

	tenantList, err := cli.Tenants.ListTenants(listTenantReqParam)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(tenantList.Error())

}
