#!/bin/bash
# Generate Certificate Authority
openssl genrsa -out ssl/server.key 2048
openssl req -new -x509 -sha256 -key ssl/server.key -out ssl/server.crt -days 3650 -subj "/C=US/ST=Philadelphia/L=PHL/O=Drexel CCI/OU=CS Dept/CN=localhost"

#Gernerate Key and Sign
openssl req -new -sha256 -key ssl/server.key -out ssl/server.csr  -subj "/C=US/ST=Philadelphia/L=PHL/O=Drexel CCI/OU=CS Dept/CN=localhost"
openssl x509 -req -sha256 -in ssl/server.csr -signkey ssl/server.key -out ssl/server.crt -days 3650