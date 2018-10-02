package latex

import (
	"io"
	"text/template"
	"time"

	"github.com/jinzhu/now"
	"github.com/nicksnyder/go-i18n/i18n"

	"github.com/icedream/workreport-manager/export/latex/stringutil"
	"github.com/icedream/workreport-manager/project"
)

// TexMarker represents information about a TeX document to be read in by the
// respective compiler. It contains information for example about which TeX
// variant to use.
type TexMarker struct {
	Program string
}

// Exporter provides functionality to export a workreports project to a LaTeX
// file.
type Exporter struct {
	Locale string
	Inputs []string
	Marker TexMarker
}

// Export generates LaTeX code from the given project and writes it to the given
// writer.
func (e *Exporter) Export(prj *project.Project, w io.Writer) (err error) {
	originalT, err := i18n.Tfunc(e.Locale)
	if err != nil {
		return
	}

	// Wrap translations with latexize just to be safe
	// type TranslateFunc func(translationID string, args ...interface{}) string
	T := func(translationID string, args ...interface{}) string {
		return stringutil.Latexize(originalT(translationID, args...))
	}

	// TODO: Reflect WeekStartDay exactly in project file, too
	if prj.FirstDayMonday {
		now.WeekStartDay = time.Monday
	}

	exportTemplate = exportTemplate.Funcs(template.FuncMap{
		"T": T,
		"beginofweek": func(date project.Date) project.Date {
			day := now.New(date.Time).BeginningOfWeek()
			if prj.OnlyShowWorkDays && now.WeekStartDay != time.Monday {
				day = day.Add(time.Hour * 24)
			}
			return project.Date{Time: day}
		},
		"endofweek": func(date project.Date) project.Date {
			day := now.New(date.Time).EndOfWeek()
			if prj.OnlyShowWorkDays {
				day = day.Add(time.Hour * -24)
				if now.WeekStartDay == time.Monday {
					day = day.Add(time.Hour * -24)
				}
			}
			return project.Date{Time: day}
		},
	})

	// Generate work report numbers
	actualNumber := 0
	weeknums := map[int]int{}
	for i, week := range prj.Weeks {
		actualNumber++

		if week.Number > 0 {
			// Use custom number for this week's report
			actualNumber = week.Number
		}

		weeknums[i] = actualNumber
	}

	data := struct {
		Project     *project.Project
		TexMarker   TexMarker
		TexInputs   []string
		WeekNumbers map[int]int
	}{
		Project:     prj,
		TexInputs:   e.Inputs,
		TexMarker:   e.Marker,
		WeekNumbers: weeknums,
	}
	return exportTemplate.Execute(w, data)
}
