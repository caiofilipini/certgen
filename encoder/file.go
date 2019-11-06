package encoder

import (
	"fmt"
	"os"
	"strings"

	"github.com/caiofilipini/certgen/cert"
)

const (
	// DefaultOutputDir is the default directory all files will be
	// written in when a custom directory is not specified.
	DefaultOutputDir = "./"

	timestampFmt = "2006-01-02T150405"
)

func fileNamePrefixFor(c cert.Certificate, dir string) string {
	if strings.HasSuffix(dir, "/") {
		dir = strings.TrimSuffix(dir, "/")
	}
	return fmt.Sprintf("%s/%s-%s", dir, c.OrgName, c.NotBefore.Format(timestampFmt))
}

func openFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0600)
}
