package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

const dateFormat = "2006-01-02"

func main() {
	raw := flag.Bool("raw", false, "output raw JSON")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: tls-check [--raw] <domain>")
		os.Exit(1)
	}

	domain := flag.Arg(0)

	// Analyze host (polling handled in analyzeHost())
	result, err := analyzeHost(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if *raw {
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(data))
		return
	}

	fmt.Println()
	fmt.Println("TLS Report for:", domain)
	fmt.Println("--------------------------------")

	// Build cert lookup map
	certMap := make(map[string]Cert)
	for _, cert := range result.Certs {
		certMap[cert.Id] = cert
	}

	// Iterate all endpoints
	for _, ep := range result.Endpoints {
		fmt.Println()
		fmt.Println("Endpoint:", ep.IPAddress)

		if ep.StatusMessage != "" {
			fmt.Println("Status:", ep.StatusMessage)
		}

		if ep.Grade != "" {
			fmt.Println("Grade:", ep.Grade)
		}

		// Protocols
		if len(ep.Details.Protocols) > 0 {
			fmt.Print("Protocols: ")
			for i, p := range ep.Details.Protocols {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(p.Name, " ", p.Version)
			}
			fmt.Println()
		}

		// Certificate expiration (leaf cert)
		if len(ep.Details.CertChains) > 0 &&
			len(ep.Details.CertChains[0].CertIds) > 0 {

			leafCertID := ep.Details.CertChains[0].CertIds[0]
			if cert, ok := certMap[leafCertID]; ok {
				expiration := time.Unix(cert.NotAfter/1000, 0)
				fmt.Println("Certificate valid until:",
					expiration.Format("2006-01-02"))
			}
		}
	}
}
