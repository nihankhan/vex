package collector

import (
	"os/exec"
	"strings"
)

type OSCollector struct {
}

func (o *OSCollector) Name() string { return "OS" }

func (o *OSCollector) Collect() (*Info, error) {
	args := `
	source /etc/os-release
	arch=$(uname -m)
	echo "$PRETTY_NAME $arch"
	`

	cmd := exec.Command("bash", "-c", args)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	osName := strings.TrimSpace(string(out))

	return &Info{
		Label: o.Name(),
		Value: osName,
	}, nil
}
