package project

// Project represents the root structure of a project file.
type Project struct {
	Name       string `json:"Name"`
	Department string `json:"Department"`
	Begin      Date   `json:"Begin"`
	End        Date   `json:"End"`
	Weeks      []Week `json:"Weeks"`
}
