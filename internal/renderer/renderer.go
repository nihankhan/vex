package renderer

import (
	"fmt"
	"strings"

	"github.com/nihankhan/vex/internal/logo"
)

type Renderer struct {
	LogoLines []string
	InfoLines []string
}

func NewRenderer(logoName string, info []string) *Renderer {
	provider := logo.Get(logoName)

	return &Renderer{
		LogoLines: provider.Lines(),
		InfoLines: info,
	}
}

var ubuntuStart = RGB{R: 255, G: 85, B: 0}
var ubuntuEnd = RGB{R: 255, G: 200, B: 0}

func (r *Renderer) Render() {
	logoLines := r.LogoLines

	var infoLines []string

	// ── HEADER ──
	if len(r.InfoLines) > 0 {
		first := r.InfoLines[0]

		parts := strings.SplitN(first, ":", 2)

		// Identity: gotipath@backend
		if len(parts) == 2 && strings.TrimSpace(parts[0]) == "Identity" {
			identity := strings.TrimSpace(parts[1])

			infoLines = append(infoLines,
				BoldColorize(identity, Green),
				Colorize(strings.Repeat("-", len(identity)), White),
				"",
			)

			// remove Identity from normal table rendering
			r.InfoLines = r.InfoLines[1:]
		}
	}

	// ── INFO LINES ──
	for _, line := range r.InfoLines {
		infoLines = append(infoLines, line)

		// add palette right after Memory
		if strings.HasPrefix(line, "Memory:") {
			infoLines = append(infoLines,
				"",
				paletteBar(),
				paletteBar(),
			)
		}
	}

	colors := GenerateGradient(ubuntuStart, ubuntuEnd, len(logoLines))

	max := len(logoLines)
	if len(infoLines) > max {
		max = len(infoLines)
	}

	for i := 0; i < max; i++ {
		var l, rline string

		// ── LOGO ──
		if i < len(logoLines) {
			color := White

			if i < len(colors) {
				color = colors[i]
			}

			l = Colorize(logoLines[i], color)
		}

		// ── INFO ──
		if i < len(infoLines) {
			line := infoLines[i]

			// palette line
			if strings.Contains(line, "\033[48;5;") {
				rline = line
			} else {
				parts := strings.SplitN(line, ":", 2)

				if len(parts) == 2 {
					label := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])

					coloredLabel := BoldColorize(label, infoColor(label))

					rline = fmt.Sprintf("%-12s: %s", coloredLabel, value)
				} else {
					rline = line
				}
			}
		}

		fmt.Printf("%-55s %s\n", l, rline)
	}
}
