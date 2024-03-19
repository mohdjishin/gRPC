#!/bin/bash

rm -f certs/client/*.pem certs/server/*.pem certs/CA/*.pem certs/client/*.p12 certs/server/*.p12

openssl req -x509 -newkey rsa:4096 -days 365 -nodes \
    -keyout certs/CA/ca-key.pem -out certs/CA/ca-cert.pem \
    -subj "/C=FR/ST=Occitanie/L=Toulouse/O=Tech School/OU=Education/CN=*.test.in/emailAddress=testgmail.com"

echo "CA's self-signed certificate"
openssl x509 -in certs/CA/ca-cert.pem -noout -text

openssl req -newkey rsa:4096 -nodes \
    -keyout certs/server/server-key.pem -out certs/server/server-req.pem \
    -subj "/C=FR/ST=Ile de France/L=Paris/O=PC GOCART/OU=Computer/CN=*.test.com/emailAddress=test@gmail.com"

openssl x509 -req -in certs/server/server-req.pem -days 60 \
    -CA certs/CA/ca-cert.pem -CAkey certs/CA/ca-key.pem -CAcreateserial \
    -out certs/server/server-cert.pem -extfile certs/server/server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in certs/server/server-cert.pem -noout -text

openssl pkcs12 -export -out certs/server/server.p12 -inkey certs/server/server-key.pem -in certs/server/server-cert.pem -certfile certs/CA/ca-cert.pem -passout pass:yourpassword

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes \
    -keyout certs/client/client-key.pem -out certs/client/client-req.pem \
    -subj "/C=FR/ST=Alsace/L=Strasbourg/O=PC Client/OU=Computer/CN=*.test.com/emailAddress=test@gmail.com"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in certs/client/client-req.pem -days 60 \
    -CA certs/CA/ca-cert.pem -CAkey certs/CA/ca-key.pem -CAcreateserial \
    -out certs/client/client-cert.pem -extfile certs/client/client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in certs/client/client-cert.pem -noout -text

# Convert client's key and certificate to PKCS#12 format
openssl pkcs12 -export -out certs/client/client.p12 -inkey certs/client/client-key.pem -in certs/client/client-cert.pem -certfile certs/CA/ca-cert.pem -passout pass:yourpassword
