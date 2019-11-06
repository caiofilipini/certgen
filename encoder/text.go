package encoder

import (
	"fmt"
	"os"

	"github.com/caiofilipini/certgen/cert"
)

// jsonEncoder is an implementation of Encoder that encodes
// Certificate objects as text.
type textEncoder struct {
	opts Options
}

// Encode takes a Certificate object and encodes it in text
// format. If the Encoder has been configured to write to disk,
// it will write two files: one for the certificate, and another
// for the private key; otherwise, both will be printed to STDOUT.
//
// It returns a non-nil error in case it fails to write any of
// the files.
func (e textEncoder) Encode(cert cert.Certificate) (Result, error) {
	if e.opts.WriteToDisk {
		var artifacts []string
		prefix := fileNamePrefixFor(cert, e.opts.OutputDir)

		crtFile, err := openFile(fmt.Sprintf("%s.crt", prefix))
		if err != nil {
			return errResult, err
		}
		defer crtFile.Close()

		if _, err := fmt.Fprintf(crtFile, cert.CertPEM); err != nil {
			return errResult, fmt.Errorf("failed to write certificate to file: %w", err)
		}

		artifacts = append(artifacts, crtFile.Name())

		keyFile, err := openFile(fmt.Sprintf("%s.key", prefix))
		if err != nil {
			return errResult, err
		}
		defer keyFile.Close()

		if _, err := fmt.Fprintf(keyFile, cert.PrivateKeyPEM); err != nil {
			return errResult, fmt.Errorf("failed to write private key to file: %w", err)
		}

		artifacts = append(artifacts, keyFile.Name())

		return Result{
			Success:   true,
			Artifacts: artifacts,
		}, nil
	}

	fmt.Fprintln(os.Stdout, cert.CertPEM)
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout, cert.PrivateKeyPEM)

	return Result{
		Success:   true,
		Artifacts: []string{},
	}, nil
}
