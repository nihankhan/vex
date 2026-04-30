package collector

import (
	"strings"

	"github.com/nihankhan/vex/internal/utils"
)

type IconCollector struct{}

func (i *IconCollector) Name() string { return "Icons" }

func (i *IconCollector) Collect() (*Info, error) {
	out, err := utils.RunCmd(`gsettings get org.gnome.desktop.interface icon-theme`)
	if err != nil {
		return nil, err
	}

	icon := strings.TrimSpace(out)
	icon = strings.Trim(icon, "'")

	if icon == "" {
		icon = "unknown"
	}

	value := icon + " [GTK2/3]"

	return &Info{
		Label: i.Name(),
		Value: value,
	}, nil
}
