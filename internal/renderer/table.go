package renderer

import "fmt"

type Row struct {
	Label string
	Value string
}

func FormatTable(rows []Row) []string {
	var out []string

	maxLen := 0
	for _, r := range rows {
		if len(r.Label) > maxLen {
			maxLen = len(r.Label)
		}
	}

	for _, r := range rows {
		out = append(out,
			fmt.Sprintf("%-*s : %s", maxLen, r.Label, r.Value),
		)
	}

	return out
}
