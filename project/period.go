package project

// Period represents a school period in the weekly report.
type Period struct {
	Subject string   `json:"Subject"`
	Topics  []string `json:"Topics"`
}
