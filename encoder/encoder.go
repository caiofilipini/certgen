package encoder

import (
	"github.com/caiofilipini/certgen/cert"
)

const (
	// FormatText is used to configure the Encoder to generate
	// its output in text format.
	FormatText = "text"

	// FormatJSON is used to configure the Encoder to generate
	// its output in JSON format.
	FormatJSON = "json"
)

var (
	errResult = Result{
		Success: false,
	}
)

// Encoder defines the operations each encoder implementation
// needs to satisfy.
type Encoder interface {
	// Encode takes a Certificate object, encodes it,
	// and returns a corresponding Result object.
	//
	// It returns a non-nil error in case it fails to
	// encode the given Certificate.
	Encode(cert.Certificate) (Result, error)
}

// Options configures the behavior of the different
// types of Encoder.
type Options struct {
	// Format configures the output format of the encoder
	// (e.g. "text" or "json").
	Format string

	// WriteToDisk controls whether or not the resulting
	// artifact should be written to disk.
	WriteToDisk bool

	// OutputDir configures the directory in which files
	// whould be written when WriteToDisk is set to `true`.
	OutputDir string
}

// Result contains information about the encoding results.
type Result struct {
	// Success is set to `true` when encoding is successful.
	Success bool

	// Artifacts will contain the names of files written
	// when the Encoder is configures to write to disk.
	Artifacts []string
}

// New takes an Options object and returns a corresponding
// Encoder to satisfy the given Options.
func New(opts Options) Encoder {
	if opts.Format == FormatJSON {
		return jsonEncoder{
			opts: opts,
		}
	}

	return textEncoder{
		opts: opts,
	}
}
