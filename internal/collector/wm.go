package collector

import (
	"os"
	"strings"
)

type WMCollector struct{}

func (w *WMCollector) Name() string { return "WM" }

func (w *WMCollector) Collect() (*Info, error) {
	de := strings.ToLower(os.Getenv("XDG_CURRENT_DESKTOP"))

	var wm string

	switch {
	case strings.Contains(de, "gnome"):
		wm = "Mutter"
	case strings.Contains(de, "kde"):
		wm = "KWin"
	case strings.Contains(de, "xfce"):
		wm = "Xfwm4"
	case strings.Contains(de, "i3"):
		wm = "i3"
	case strings.Contains(de, "sway"):
		wm = "sway"
	default:
		wm = "unknown"
	}

	return &Info{
		Label: w.Name(),
		Value: wm,
	}, nil
}
