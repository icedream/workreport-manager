package util

import (
	"bytes"

	"github.com/russross/blackfriday"
)

/*
DocumentlessRenderer wraps an actual renderer and suppresses full-document
output.
*/
type DocumentlessRenderer struct {
	blackfriday.Renderer
}

/*
DocumentHeader does nothing as it is a bogus implementation to avoid full-
document output.
*/
func (r DocumentlessRenderer) DocumentHeader(out *bytes.Buffer) {
}

/*
DocumentFooter does nothing as it is a bogus implementation to avoid full-
document output.
*/
func (r DocumentlessRenderer) DocumentFooter(out *bytes.Buffer) {
}
