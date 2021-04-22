package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xorcat/cert"
)

var version = ""

func main() {
	var format string
	var template string
	var skipVerify bool
	var utc bool
	var timeout int
	var showVersion bool
	var cipherSuite string

	flag.StringVar(&format, "f", "simple table", "Output format. md: as markdown, json: as JSON. ")
	flag.StringVar(&format, "format", "simple table", "Output format. md: as markdown, json: as JSON. ")
	flag.StringVar(&template, "t", "", "Output format as Go template string or Go template file path.")
	flag.StringVar(&template, "template", "", "Output format as Go template string or Go template file path.")
	flag.BoolVar(&skipVerify, "k", false, "Skip verification of server's certificate chain and host name.")
	flag.BoolVar(&skipVerify, "skip-verify", false, "Skip verification of server's certificate chain and host name.")
	flag.BoolVar(&utc, "u", false, "Use UTC to represent NotBefore and NotAfter.")
	flag.BoolVar(&utc, "utc", false, "Use UTC to represent NotBefore and NotAfter.")
	flag.IntVar(&timeout, "s", 3, "Timeout seconds.")
	flag.IntVar(&timeout, "timeout", 3, "Timeout seconds.")
	flag.BoolVar(&showVersion, "v", false, "Show version.")
	flag.BoolVar(&showVersion, "version", false, "Show version.")
	flag.StringVar(&cipherSuite, "c", "", "Specify cipher suite. Refer to https://golang.org/pkg/crypto/tls/#pkg-constants for supported cipher suites.")
	flag.StringVar(&cipherSuite, "cipher", "", "Specify cipher suite. Refer to https://golang.org/pkg/crypto/tls/#pkg-constants for supported cipher suites.")
	flag.Parse()

	if showVersion {
		fmt.Println("cert version ", version)
		return
	}

	var certs cert.Certs
	var err error

	cert.SkipVerify = skipVerify
	cert.UTC = utc
	cert.TimeoutSeconds = timeout
	cert.CipherSuite = cipherSuite

	certs, err = cert.NewCerts(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if template == "" {
		switch format {
		case "md":
			fmt.Printf("%s", certs.Markdown())
		case "json":
			fmt.Printf("%s", certs.JSON())
		default:
			fmt.Printf("%s", certs)
		}
		return
	}

	if err := cert.SetUserTempl(template); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s", certs)
}
