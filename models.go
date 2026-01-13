package main

type SSLResponse struct {
	Status        string     `json:"status"`
	StatusMessage string     `json:"statusMessage"`
	Endpoints     []Endpoint `json:"endpoints"`
	Certs         []Cert     `json:"certs"`
}

type Endpoint struct {
	IPAddress            string  `json:"ipAddress"`
	ServerName           string  `json:"serverName,omitempty"`
	Grade                string  `json:"grade,omitempty"`
	StatusMessage        string  `json:"statusMessage"`
	StatusDetails        string  `json:"statusDetails"`
	StatusDetailsMessage string  `json:"statusDetailsMessage"`
	Details              Details `json:"details"`
}

type Details struct {
	CertChains []CertChain `json:"certChains"`
	Protocols  []Protocol  `json:"protocols"`
}

type Protocol struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CertChain struct {
	Id      string   `json:"id"`
	CertIds []string `json:"certIds"`
}

type Cert struct {
	Id       string `json:"id"`
	NotAfter int64  `json:"notAfter"`
}
