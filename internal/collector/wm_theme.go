package collector

import (
	"strings"

	"github.com/nihankhan/vex/internal/utils"
)

type WMThemeCollector struct{}

func (w *WMThemeCollector) Name() string { return "WM Theme" }

func (w *WMThemeCollector) Collect() (*Info, error) {
	out, err := utils.RunCmd(
		`gsettings get org.gnome.desktop.wm.preferences theme`,
	)
	if err != nil {
		return nil, err
	}

	theme := strings.TrimSpace(out)

	// remove quotes
	theme = strings.Trim(theme, "'")

	if theme == "" {
		theme = "unknown"
	}

	return &Info{
		Label: w.Name(),
		Value: theme,
	}, nil
}
