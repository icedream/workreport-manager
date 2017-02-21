package project

// Project represents the root structure of a project file.
type Project struct {
	Name                  string `json:"Name"`
	Department            string `json:"Department"`
	FirstDayMonday        bool   `json:"First day is monday"`
	OnlyShowWorkDays      bool   `json:"Only show work days"`
	HideLegalRepSignField bool   `json:"Hide legal representative field"`
	Begin                 Date   `json:"Begin"`
	End                   Date   `json:"End"`
	Weeks                 []Week `json:"Weeks"`
}
