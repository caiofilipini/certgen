# certgen

Simple TLS self-signed certificates generator written in Go.

**DISCLAIMER:** even though I use this tool in a bunch of test/development scenarios, it was written mostly as a learning exercise and TLS certificates generated with it **SHOULD NOT** be used in any production setting.

## Installation

Since this tool is written in Go, you can install it by running:

```
$ go install github.com/caiofilipini/certgen
```

It will be built locally, and the resulting binary will be available in `$GOPATH/bin`.

## Usage

```
$ certgen --help
Usage of certgen:
  -b uint
        number of bits in the generated key (default 2048)
  -d string
        directory where files will be written when using -w (default "./")
  -e uint
        number of days the generated certificate will be valid for (default 1)
  -f string
        output format (text, json) (default "text")
  -j    shorthand flag for changing the output format to JSON
  -o string
        organization name for the certificate (default "org")
  -w    write output to disk
```

### Examples

A quick example that will generate a TLS certificate and its corresponding private key and write both files into `/tmp`:

```
$ certgen -w -d /tmp
/tmp/org-2019-11-07T170027.crt written
/tmp/org-2019-11-07T170027.key written
```

And in order to generate certificate valid for 30 days and encode it to a JSON file:

```
$ certgen -e 30 -w -j -d /tmp
/tmp/org-2019-11-07T170246.json written

$ cat /tmp/org-2019-11-07T170246.json | jq '.not_after'
"2019-12-07T17:02:46.795701-05:00"
```
