package latex

import (
	"io"
	"text/template"

	"github.com/nicksnyder/go-i18n/i18n"

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
	T, err := i18n.Tfunc(e.Locale)
	if err != nil {
		return
	}

	exportTemplate.Funcs(template.FuncMap{
		"T": T,
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
