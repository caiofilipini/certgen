package generator

import (
	"bytes"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"math/rand"
	"time"

	"github.com/caiofilipini/certgen/cert"
)

type pemBlock string

const (
	pemBlockCertificate   = pemBlock("CERTIFICATE")
	pemBlockRSAPrivateKey = pemBlock("RSA PRIVATE KEY")
)

var (
	emptyCert = cert.Certificate{}
)

// Generator generates self-signed TLS certificates.
type Generator struct {
	keySize  uint64
	orgName  string
	validity time.Duration
}

// CertOpt is used to configure different options for the
// Generator.
type CertOpt func(g *Generator)

// WithOrgName configures the Generator to use the given orgName
// when generating certificates.
func WithOrgName(orgName string) CertOpt {
	return func(g *Generator) {
		g.orgName = orgName
	}
}

// WithKeySize configures the Generator to generate keys with
// `keySize` bits.
func WithKeySize(keySize uint64) CertOpt {
	return func(g *Generator) {
		g.keySize = keySize
	}
}

// WithValidityInDays configures the Generator to generate
// certificates that are valid for `validityInDays` days.
func WithValidityInDays(validityInDays uint64) CertOpt {
	return func(g *Generator) {
		g.validity = daysToDuration(validityInDays)
	}
}

// Newcreates and returns a new Generator based in the given `opts`.
func New(opts ...CertOpt) *Generator {
	gen := &Generator{
		orgName:  cert.DefaultOrgName,
		keySize:  cert.DefaultKeySize,
		validity: daysToDuration(cert.DefaultCertValidityInDays),
	}

	for _, opt := range opts {
		opt(gen)
	}

	return gen
}

// NewTLS generates a self-signed TLS certificate with a newly created
// private key, and returns both the cert and the private key PEM encoded.
func (g *Generator) NewTLS() (cert.Certificate, error) {
	validityStart := time.Now()
	validityEnd := validityStart.Add(g.validity)

	certTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(int64(randInt())),
		Subject: pkix.Name{
			Organization: []string{g.orgName},
		},
		NotBefore: validityStart,
		NotAfter:  validityEnd,
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
	}

	pk, pkPEM, err := genPrivateKey(g.keySize)
	if err != nil {
		return emptyCert, err
	}

	c, err := x509.CreateCertificate(crand.Reader, certTemplate, certTemplate, &pk.PublicKey, pk)
	if err != nil {
		return emptyCert, err
	}

	certPEM, err := toPEMString(pemBlockCertificate, c)
	if err != nil {
		return emptyCert, err
	}

	return cert.Certificate{
		OrgName:       g.orgName,
		CertPEM:       certPEM,
		PrivateKeyPEM: pkPEM,
		NotBefore:     validityStart,
		NotAfter:      validityEnd,
	}, nil
}

func genPrivateKey(size uint64) (*rsa.PrivateKey, string, error) {
	pk, err := rsa.GenerateKey(crand.Reader, int(size))
	if err != nil {
		return nil, "", err
	}

	pkPEM, err := toPEMString(pemBlockRSAPrivateKey, x509.MarshalPKCS1PrivateKey(pk))
	if err != nil {
		return nil, "", err
	}

	return pk, pkPEM, nil
}

func toPEMString(blockType pemBlock, b []byte) (string, error) {
	pb := &pem.Block{
		Type:  string(blockType),
		Bytes: b,
	}

	var buf bytes.Buffer
	if err := pem.Encode(&buf, pb); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func randInt() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int()
}

func daysToDuration(d uint64) time.Duration {
	if d == 0 {
		d = cert.DefaultCertValidityInDays
	}
	return time.Duration(d) * time.Hour * 24
}
