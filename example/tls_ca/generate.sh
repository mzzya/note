#! /bin/bash

openssl genrsa -out ca.key 2048

# 默认-----BEGIN CERTIFICATE REQUEST-----
openssl req -new  -key ca.key -out ca.pem
openssl req -new  -key ca.key -out ca.csr

openssl x509 -req -in ca.pem -days 9999 -singkey=ca.key --out ca.crt