package strutil

// RemoveEmpty removes empty strings from a string slice
func RemoveEmpty(s *[]string) {
	i := 0
	p := *s

	for _, sub := range p {
		if sub != "" {
			p[i] = sub
			i++
		}
	}

	*s = p[0:i]
}

// Select removes strings fror a string slice that does not pass a selector function
func Select(s *[]string, f func(string) bool) {
	i := 0
	p := *s

	for _, sub := range p {
		if f(sub) {
			p[i] = sub
			i++
		}
	}

	*s = p[0:i]
}