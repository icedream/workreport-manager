package stringutil

import (
	"strings"

	"github.com/icedream/workreportmgr/internal/util"
	"github.com/russross/blackfriday"
)

const (
	enabledExtensions = 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS
)

var additionalTexReplacements = []Converter{
	// NewStringConverter("{", "\\{"),
	// NewStringConverter("}", "\\}"),
	// NewStringConverter("\\", "\\textbackslash{}"),
	//
	// NewStringConverter("&", "\\&"),
	// NewStringConverter("%", "\\%"),
	// NewStringConverter("$", "\\$"),
	// NewStringConverter("#", "\\#"),
	// NewStringConverter("_", "\\_"),
	// NewStringConverter("~", "\\textasciitilde{}"),
	// NewStringConverter("^", "\\textasciicircum{}"),
	// NewStringConverter("ß", "\\ss{}"),

	NewSimpleRegexConverter(`"([^"]+)"`, `\enquote{$1}`),
	// NewSimpleRegexConverter("`([^`]+)`", "\\verb`$1`"),
}

/*
Latexize takes an input text as parsed from the value of any field in a project
file and turns it into LaTeX code.
*/
func Latexize(input string) string {
	renderer := util.DocumentlessRenderer{Renderer: blackfriday.LatexRenderer(0)}
	renderedString := string(blackfriday.Markdown([]byte(input), renderer, enabledExtensions))
	renderedString = CustomEscape(renderedString, additionalTexReplacements...)
	renderedString = strings.TrimSpace(renderedString)
	return renderedString
}
