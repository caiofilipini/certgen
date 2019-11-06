package cert

import "time"

const (
	// DefaultOrgName is the default organization name used in
	// certificates when a custom organization is not configured.
	DefaultOrgName = "org"

	// DefaultKeySize is the default number of bits in the generated
	// private key for the certificate when a custom size is not
	// configured.
	DefaultKeySize = 2048

	// DefaultCertValidityInDays is the default number of days a
	// generated certificate will be valid for when a custom validity
	// is not configured.
	DefaultCertValidityInDays = 1
)

// Certificate represents a generated TLS certificate.
type Certificate struct {
	// OrgName is the organization name used in the certificate.
	OrgName string `json:"org_name"`

	// CertPEM is the PEM-encoded certificate.
	CertPEM string `json:"certificate"`

	// PrivateKeyPEM is the PEM-encoded private key associated
	// with the certificate.
	PrivateKeyPEM string `json:"private_key"`

	// NotBefore is the timestamps of when the certificate was
	// generated and started being valid.
	NotBefore time.Time `json:"not_before"`

	// NotAfter is the timestamps of when the certificate will
	// expire.
	NotAfter time.Time `json:"not_after"`
}
