#!/bin/bash


echo "Get Tenants"
curl -sk -H 'Content-Type: application/json'  http://localhost:9090/api/beta1/tenants | jq .


