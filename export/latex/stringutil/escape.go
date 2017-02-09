package stringutil

var replacements = []converter{
	newStringConverter("{", "\\{"),
	newStringConverter("}", "\\}"),
	newStringConverter("\\", "\\textbackslash{}"),

	newStringConverter("&", "\\&"),
	newStringConverter("%", "\\%"),
	newStringConverter("$", "\\$"),
	newStringConverter("#", "\\#"),
	newStringConverter("_", "\\_"),
	newStringConverter("~", "\\textasciitilde{}"),
	newStringConverter("^", "\\textasciicircum{}"),
	newStringConverter("ß", "\\ss{}"),

	newSimpleRegexConverter(`"([^"]+)"`, `\enquote{$1}`),
}

// TexEscape modifies a string so it can be safely places in a LaTeX file
// without causing any errors due to special characters.
func TexEscape(s string) string {
	for _, replacer := range replacements {
		s = replacer.Process(s)
	}

	return s
}
