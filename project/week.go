package project

// Week represents a week in the report.
type Week struct {
	Date                Date
	Number              int      `json:"Number,omitempty"`
	Periods             []Period `json:"Professional school"`
	WorkActivities      []string `json:"Operational activities"`
	WorkActivityDetails string   `json:"Operational instructions"`
}
