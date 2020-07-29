package strutil

import (
	"strings"
)

// Prefix returns the specified number of characters from the start of a string. If 
// the number of characters specified is more than the length of the string, the entire string 
// is returned. If the length is negative, characters are returned from the end of the string.
func Prefix(s string, length int) string {
	if s == "" {
		return ""
	}

	if length < 0 {
		return Suffix(s, length * -1)
	}

	if len(s) < length {
		return string(s)
	}

	return s[:length]
}

// Suffix returns the specified number of characters from the end of a string. If 
// the number of characters specified is more than the length of the string, the entire string 
// is returned. If the length is negative, characters are returned from the start of the string.
func Suffix(s string, length int) string {
	if s == "" {
		return ""
	}

	if length < 0 {
		return Prefix(s, length * -1)
	}

	if len(s) < length {
		return s
	}

	return s[len(s) - length:]
}

// Before returns all characters before the the specified substring
func Before(s, substr string) string {
	if s == "" || substr == "" {
		return s
	}

	cursor := strings.Index(s, substr)
	if cursor < 0 {
		return s
	}

	if cursor == 0 {
		return ""
	}

	return s[0:cursor]
}

// After returns all characters after the the specified substring
func After(s, substr string) string {
	if s == "" || substr == "" {
		return s
	}

	cursor := strings.Index(s, substr)
	if cursor < 0 {
		return s
	}

	return s[cursor + len(substr):]
}

// Between returns all substrings sandwiched between the specified head and tail
func Between(s string, head, tail []string, first int) []string {
	if first == 0 {
		return []string{}
	}

	cursor := 0
	found := 0
	x, y := -1, -1

	output := []string{}

	for true {
		// find head tags
		for i := 0; i < len(head); i++ {
			found = strings.Index(s[cursor:], head[i])
			if found == -1 {
				break
			} else {
				cursor = cursor + found + len(head[i])
			}
		}

		if found == -1 {
			return output
		}

		x = cursor

		// find tail tags
		for i := 0; i < len(tail); i++ {
			found = strings.Index(s[cursor:], tail[i])
			if found == -1 {
				break
			} else {
				cursor = cursor + found + len(tail[i])
			}

			// exit if not all head tags found, else get target start index
			if found == -1 {
				return output
			}
		}

		y = x + strings.Index(s[x:], tail[0])

		// add to output
		output = append(output, s[x:y])
		if first > -1 && len(output) >= first {
			return output
		}
	}

	return output
}
