#!/bin/bash


echo "Get Tenants"
curl -sk -H 'Content-Type: application/json'  http://localhost:9090/api/beta2/tenants | jq .


echo "Create Test Tenants"

id=$(curl -sk -H 'Content-Type: application/json'  http://localhost:9090/api/beta2/tenants -X POST -d'{"name":"Test-Org", "description": "some test org"}' | jq -r .id) 

echo "updating test tenant"
curl -sk -H 'Content-Type: application/json'  http://localhost:9090/api/beta2/tenants -X PUT -d"{\"id\":\"${id}\", \"name\":\"Test-Org\", \"description\": \"some other  test org\"}" | jq . 

echo "getting new updated tenant"
curl -sk -H 'Content-Type: application/json'  "http://localhost:9090/api/beta2/tenants/${id}" | jq .


echo "Generate 2 license for tenant"
curl -sk -H 'Content-Type: application/json'  "http://localhost:9090/api/beta2/license/generate_license_for_tenant_id/${id}" | jq .
curl -sk -H 'Content-Type: application/json'  "http://localhost:9090/api/beta2/license/generate_license_for_tenant_id/${id}" | jq .

echo "List license for tenant"
curl -sk -H 'Content-Type: application/json'  "http://localhost:9090/api/beta2/license/get_licenses_by_tenant_id/${id}" | jq .