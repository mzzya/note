openssl genrsa -out ca.key 2048

# 默认-----BEGIN CERTIFICATE REQUEST-----
openssl req -new  -key ca.key -out ca.pem


openssl x509 -req -in ca.pem -days 9999 -signkey ca.key -out ca.crt