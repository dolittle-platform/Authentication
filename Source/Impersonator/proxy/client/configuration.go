package client

type Configuration interface {
	ServiceAccountTokenPath() string
	CertificateAuthorityPath() string
}
