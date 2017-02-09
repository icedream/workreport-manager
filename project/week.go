package project

// Week represents a week in the report.
type Week struct {
	Date                Date
	WorkActivities      []string `json:"Operational activities"`
	WorkActivityDetails string   `json:"Operational instructions"`
	Periods             []Period `json:"Professional school"`
}
