package collector

import (
	"os"
	"strings"

	"github.com/nihankhan/vex/internal/utils"
)

type ThemeCollector struct{}

func (t *ThemeCollector) Name() string { return "Theme" }

func (t *ThemeCollector) Collect() (*Info, error) {
	gtk3, err := utils.RunCmd(`gsettings get org.gnome.desktop.interface gtk-theme`)
	if err != nil {
		return nil, err
	}
	gtk3 = strings.Trim(strings.TrimSpace(gtk3), "'")

	gtk2 := getGTK2Theme()

	if gtk2 == "" {
		gtk2 = gtk3
	}

	value := gtk3 + " [GTK2/3]"

	return &Info{
		Label: t.Name(),
		Value: value,
	}, nil
}

func getGTK2Theme() string {
	home, _ := os.UserHomeDir()

	out, err := utils.RunCmd(`grep gtk-theme ` + home + `/.gtkrc-2.0 2>/dev/null | cut -d"=" -f2`)
	if err != nil {
		return ""
	}
	out = strings.TrimSpace(out)
	out = strings.Trim(out, "\"")

	return out
}
