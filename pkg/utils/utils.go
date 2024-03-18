package utils

import (
	"crypto/x509"

	"software.sslmate.com/src/go-pkcs12"
)

func DecodeP12File(p12Data []byte, password string) (privateKey interface{}, leaf *x509.Certificate, roots *x509.CertPool, err error) {

	privateKey, certificate, ca, err := pkcs12.DecodeChain(p12Data, password)
	if err != nil {
		return
	}

	leaf = certificate
	roots = x509.NewCertPool()
	for _, intermediateCert := range ca {
		if intermediateCert.IsCA {
			roots.AddCert(intermediateCert)
		}
	}

	return
}
