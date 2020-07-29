// Package strutil gets your string manipulation covered! Features include escaping reserved 
// words, wildcard compare, and pattern extraction.
package strutil

import (
	"strings"
)

// Escape replaces all occurances of reserved with substitute in text, but makes sure that 
// any substitute already in text can be recovered using Unescape().
func Escape(text, reserved, subst string) string {
	if text == "" || reserved == "" || subst == "" {
		return ""
	}

	return strings.Replace(strings.Replace(text, subst, subst + "2", -1), reserved, subst + "1", -1)
}

// Unescape recovers text processed by the Escape() function.
func Unescape(text, reserved, subst string) string {
	if text == "" || reserved == "" || subst == "" {
		return ""
	}

	return strings.Replace(strings.Replace(text, subst + "1", reserved, -1), subst + "2", subst, -1)
}
