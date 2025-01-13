package models

import (
	"fmt"
	"regexp"
	"strings"
)

type ExpectRow struct {
	Text   string
	Colors []string
}

func (row ExpectRow) Render() string {

	r := regexp.MustCompile(`<\|(.*?)\|>`)

	textMatches := r.FindAllStringSubmatch(row.Text, -1)
	var texts []string
	for _, m := range textMatches {
		texts = append(texts, m[1])
	}

	for i, c := range row.Colors {
		if strings.HasPrefix(c, "rgb") {
			var r, g, b int
			fmt.Sscanf(c, "rgb(%d,%d,%d)", &r, &g, &b)
			texts[i] = fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, texts[i])
		}
		if strings.HasPrefix(c, "ansi") {
			var code int
			fmt.Sscanf(c, "ansi(%d)", &code)
			texts[i] = fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", code, texts[i])
		}
	}

	indexMatches := r.FindAllStringIndex(row.Text, -1)
	result := []byte(row.Text)
	offset := 0

	for i, match := range indexMatches {
		if i >= len(texts) {
			break
		}

		start := match[0] + offset
		end := match[1] + offset

		result = append(result[:start], append([]byte(texts[i]), result[end:]...)...)
		offset += len(texts[i]) - (end - start)
	}

	return string(result)
}