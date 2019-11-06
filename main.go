package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/caiofilipini/certgen/cert"
	"github.com/caiofilipini/certgen/encoder"
	"github.com/caiofilipini/certgen/generator"
)

func main() {
	var (
		orgName        string
		keySize        uint64
		validityInDays uint64
		format         string
		jsonOutput     bool
		writeToDisk    bool
		outputDir      string
	)

	flag.StringVar(&orgName, "o", cert.DefaultOrgName, "organization name for the certificate")
	flag.Uint64Var(&keySize, "b", cert.DefaultKeySize, "number of bits in the generated key")
	flag.Uint64Var(&validityInDays, "e", cert.DefaultCertValidityInDays, "number of days the generated certificate will be valid for")

	flag.StringVar(&format, "f", encoder.FormatText, fmt.Sprintf("output format (%s, %s)", encoder.FormatText, encoder.FormatJSON))
	flag.BoolVar(&jsonOutput, "j", false, "shorthand flag for changing the output format to JSON")
	flag.BoolVar(&writeToDisk, "w", false, "write output to disk")
	flag.StringVar(&outputDir, "d", encoder.DefaultOutputDir, "directory where files will be written when using -w")

	flag.Parse()

	certgen := generator.New(
		generator.WithOrgName(orgName),
		generator.WithKeySize(keySize),
		generator.WithValidityInDays(validityInDays),
	)

	cert, err := certgen.NewTLS()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate TLS certificate: %v\n", err)
		os.Exit(1)
	}

	if jsonOutput {
		format = encoder.FormatJSON
	}

	encoder := encoder.New(encoder.Options{
		Format:      format,
		WriteToDisk: writeToDisk,
		OutputDir:   outputDir,
	})

	res, err := encoder.Encode(cert)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate output: %v\n", err)
		os.Exit(1)
	}

	for _, artifact := range res.Artifacts {
		fmt.Fprintf(os.Stderr, "%s written\n", artifact)
	}
}
