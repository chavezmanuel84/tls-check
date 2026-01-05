package main

type SSLResponse struct {
	Status    string     `json:"status"`
	Endpoints []Endpoint `json:"endpoints"`
	Certs     []Cert     `json:"certs"`
}

type Endpoint struct {
	Grade   string  `json:"grade"`
	Details Details `json:"details"`
}

type Details struct {
	Protocols []Protocol `json:"protocols"`
}

type Protocol struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Cert struct {
	NotAfter int64 `json:"notAfter"`
}
