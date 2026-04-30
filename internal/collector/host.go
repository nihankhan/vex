package collector

import (
	"os/exec"
	"strings"
)

type HostCollector struct{}

func (h *HostCollector) Name() string { return "Host" }

func (h *HostCollector) Collect() (*Info, error) {
	args := `
	product=$(cat /sys/devices/virtual/dmi/id/product_name)
	version=$(cat /sys/devices/virtual/dmi/id/product_version)
	echo "$product $version"
	`

	cmd := exec.Command("bash", "-c", args)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	host := strings.TrimSpace(string(out))

	return &Info{
		Label: h.Name(),
		Value: host,
	}, nil
}
