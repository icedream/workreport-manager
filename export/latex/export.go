package latex

import (
	"io"
	"text/template"
	"time"

	"github.com/jinzhu/now"
	"github.com/nicksnyder/go-i18n/i18n"

	"git.dekart811.net/icedream/workreportmgr/export/latex/stringutil"
	"git.dekart811.net/icedream/workreportmgr/project"
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

	now.FirstDayMonday = prj.FirstDayMonday

	exportTemplate = exportTemplate.Funcs(template.FuncMap{
		"T": T,
		"beginofweek": func(date project.Date) project.Date {
			day := now.New(date.Time).BeginningOfWeek()
			if prj.OnlyShowWorkDays && !now.FirstDayMonday {
				day = day.Add(time.Hour * 24)
			}
			return project.Date{Time: day}
		},
		"endofweek": func(date project.Date) project.Date {
			day := now.New(date.Time).EndOfWeek()
			if prj.OnlyShowWorkDays {
				day = day.Add(time.Hour * -24)
				if now.FirstDayMonday {
					day = day.Add(time.Hour * -24)
				}
			}
			return project.Date{Time: day}
		},
	})

	data := struct {
		Project   *project.Project
		TexMarker TexMarker
		TexInputs []string
	}{
		Project:   prj,
		TexInputs: e.Inputs,
		TexMarker: e.Marker,
	}
	return exportTemplate.Execute(w, data)
}
