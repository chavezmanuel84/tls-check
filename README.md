# TLS Check — SSL Labs CLI (Go)
This is a small command-line tool written in Go that checks the TLS/SSL security configuration of a given domain using the **Qualys SSL Labs API**.

## Overview

The tool analyzes a domain’s TLS configuration and reports:

- The SSL Labs grade
- Supported TLS protocol versions
- Certificate expiration date

Since SSL Labs analysis is asynchronous, the tool waits for the scan to complete before producing the final report.

## Requirements

- Go 1.20+
- Internet connection (to reach the SSL Labs API)

## Clone
```bash
git clone https://github.com/chavezmanuel84/tls-check.git
cd tls-check
```

## Build

Compile the binary:

```bash
go build -o tls-check
```
## Usage

```bash
./tls-check google.com
```

## Example Output:

```bash
Status READY
Analysis completed

TLS Info for: example.com
--------------------------------
✔ Grade: A
✔ Protocols: TLS 1.2, TLS 1.3
✔ Certificate valid until: 2026-02-25
```

## How it works

1. The tool sends a request to the SSL Labs analysis endpoint for the given domain.
2. Because the analysis is not immediate, it polls the API until the scan status is READY.
3. Once completed, relevant TLS information is extracted and displayed.

## Expected behavior

SSL Labs analyses are asynchronous. For domains that have not been recently analyzed, the scan may take several seconds or minutes to complete.

During this time, the tool will periodically display the current analysis status while polling the SSL Labs API. Once the analysis reaches the `READY` state, the final TLS report is printed.


## Project Structure

```text
tls-check/
├── main.go        # CLI flow, polling logic, output
├── sslabs.go      # SSL Labs API interaction
├── models.go      # JSON models
├── go.mod
└── README.md
```