package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tls-check <domain>")
		return
	}

	domain := os.Args[1]

	for {
		result, err := fetchSSLStatus(domain)
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		fmt.Println("Status", result.Status)

		if result.Status == "READY" {
			fmt.Println("Analysis completed")
			break
		}

		fmt.Println("Waiting 10 seconds...")
		time.Sleep(10 * time.Second)
	}

	result, _ := fetchSSLStatus(domain)

	fmt.Println()
	fmt.Println("TLS Info for:", domain)
	fmt.Println("--------------------------------")

	if len(result.Endpoints) > 0 {
		ep := result.Endpoints[0]

		fmt.Println("✔ Grade:", ep.Grade)

		fmt.Print("✔ Protocols: ")
		for i, p := range ep.Details.Protocols {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(p.Name, " ", p.Version)
		}
		fmt.Println()
	}

	// Certificate
	if len(result.Certs) > 0 {
		expiration := time.Unix(result.Certs[0].NotAfter/1000, 0)
		fmt.Println("✔ Certificate valid until:", expiration.Format("2006-01-02"))
	}
}
