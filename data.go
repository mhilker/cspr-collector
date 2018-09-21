package main

import "time"

type CSPReport struct {
	DocumentUri        string    `json:"document-uri"`
	Referrer           string    `json:"referrer"`
	ViolatedDirective  string    `json:"violated-directive"`
	EffectiveDirective string    `json:"effective-directive"`
	OriginalPolicy     string    `json:"original-policy"`
	Disposition        string    `json:"disposition"`
	BlockedUri         string    `json:"blocked-uri"`
	StatusCode         int       `json:"status-code"`
	ScriptSample       string    `json:"script-sample"`
	Occurred           time.Time `json:"occurred"`
}

type CSPRequest struct {
	Report CSPReport `json:"csp-report"`
}

func NewCSPRequest() CSPRequest {
	report := CSPRequest{
		Report: CSPReport{
			Occurred: time.Now(),
		},
	}

	return report
}
