package collector

import (
	"os"
	"strings"

	"github.com/nihankhan/vex/internal/utils"
)

type DECollector struct{}

func (d *DECollector) Name() string { return "DE" }

func (d *DECollector) Collect() (*Info, error) {
	de := os.Getenv("XDG_CURRENT_DESKTOP")
	if de == "" {
		de = "unknown"
	}

	de = strings.TrimSpace(de)

	var version string

	if strings.Contains(strings.ToLower(de), "gnome") {
		out, err := utils.RunCmd(`gnome-shell --version`)
		if err != nil {
			return nil, err
		}
		version = parseGNOME(out)
	} else {
		version = ""
	}

	value := de
	if version != "" {
		value = de + " " + version
	}

	return &Info{
		Label: d.Name(),
		Value: value,
	}, nil
}

func parseGNOME(out string) string {
	parts := strings.Fields(out)
	if len(parts) >= 3 {
		return parts[2]
	}
	return ""
}
