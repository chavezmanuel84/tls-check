# TLS Check — SSL Labs CLI (Go)
This is a small command-line tool written in Go that checks the TLS/SSL security configuration of a given domain using the [Qualys SSL Labs API](https://github.com/ssllabs/ssllabs-scan/blob/master/ssllabs-api-docs-v3.md).

## Overview

The tool analyzes a domain’s TLS configuration and reports for each domain's endpoint:

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
./tls-check example.com
```

### Example Output:

```bash
Status: DNS
Status: IN_PROGRESS
Waiting for SSL Labs analysis to complete...
Waiting for SSL Labs analysis to complete...
Status: READY

TLS Report for: example.com
--------------------------------

Endpoint: 93.184.216.34
Status: Ready
Grade: A
Protocols: TLS 1.2, TLS 1.3
Certificate valid until: 2026-06-02

Endpoint: 93.184.216.91
Status: Ready
Grade: A
Protocols: TLS 1.2, TLS 1.3
Certificate valid until: 2026-06-02
```

### Error example
```bash
Status: ERROR
Error: SSL Labs error: Unable to resolve domain name
```
### Raw JSON output (Optional)

To print the full SSL Labs JSON response use the --raw flag:

```bash
./tls-check --raw example.com
```

## How it works

1. The tool sends a request to the SSL Labs **Analyze API** for the given domain.
2. Because the analysis is not immediate, it polls the API until the scan status is `READY`.
3. Once completed, relevant TLS information is extracted and displayed.
4. The tool polls the API for a limited time (approximately 5 minutes). If the analysis does not complete within this window, the program exits with an error.

## Expected behavior

Since SSL Labs analyses are asynchronous, the scan may take several seconds or minutes to complete for domains that have not been recently analyzed.

During this time, the tool will periodically display the current analysis status while polling the SSL Labs API. Once the analysis reaches the `READY` state, the final TLS report is printed.

If the SSL Labs API returns an `ERROR` status, the tool will stop polling and display the error message provided by the API.

## Project Structure

```text
tls-check/
├── main.go        # CLI entrypoint, flags, output formatting
├── sslabs.go      # SSL Labs API client + polling/status handling
├── models.go      # JSON response models
├── go.mod
└── README.md
```