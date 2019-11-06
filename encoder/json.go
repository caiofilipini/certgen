package encoder

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caiofilipini/certgen/cert"
)

// jsonEncoder is an implementation of Encoder that encodes
// Certificate objects as JSON.
type jsonEncoder struct {
	opts Options
}

// Encode takes a Certificate object and encodes it in JSON
// format. If the Encoder has been configured to write to disk,
// then the resulting JSON will be written to a file; otherwise,
// it will be printed to STDOUT.
//
// It returns a non-nil error in case it fails to encode the given
// Certificate into JSON, or if it fails to write to disk.
func (e jsonEncoder) Encode(cert cert.Certificate) (Result, error) {
	var artifacts []string
	out := os.Stdout

	if e.opts.WriteToDisk {
		jsonFile, err := openFile(fmt.Sprintf("%s.json", fileNamePrefixFor(cert, e.opts.OutputDir)))
		if err != nil {
			return errResult, err
		}
		defer jsonFile.Close()

		artifacts = append(artifacts, jsonFile.Name())
		out = jsonFile
	}

	if err := json.NewEncoder(out).Encode(cert); err != nil {
		return errResult, fmt.Errorf("failed to generate JSON output: %w", err)
	}

	return Result{
		Success:   true,
		Artifacts: artifacts,
	}, nil
}
