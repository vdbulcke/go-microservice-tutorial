#!/bin/bash

grpcurl --plaintext -d '{"TenantID": "42", "OrganizationName": "USD"}' localhost:9092 LicenseService/NewLicense