package app

import "strings"

// reflowBody joins indented continuation lines within list items so that
// goldmark does not render them as hard line breaks.
func reflowBody(body string) string {
	lines := strings.Split(body, "\n")
	out := make([]string, 0, len(lines))

	prevBlank := false

	for _, line := range lines {
		blank := strings.TrimSpace(line) == ""
		if len(out) > 0 && !prevBlank && !blank && strings.HasPrefix(line, "  ") && !isListItem(line) {
			out[len(out)-1] += " " + strings.TrimSpace(line)
		} else {
			out = append(out, line)
		}

		prevBlank = blank
	}

	return strings.Join(out, "\n")
}

func isListItem(s string) bool {
	t := strings.TrimLeft(s, " \t")

	return strings.HasPrefix(t, "- ") || strings.HasPrefix(t, "* ") || strings.HasPrefix(t, "+ ")
}
