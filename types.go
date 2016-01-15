package main

// Params are the parameters that the GitHub Release plugin can parse.
type Params struct {
	BaseURL   string   `json:"base_url"`
	UploadURL string   `json:"upload_url"`
	APIKey    string   `json:"api_key"`
	Files     []string `json:"files"`
	Checksums []string `json:"checksums"`
}