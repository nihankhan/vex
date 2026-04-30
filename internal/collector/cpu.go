package collector

import (
	"os"
	"strings"
)

type CPUCollector struct{}

func (c *CPUCollector) Name() string { return "CPU" }

func (c *CPUCollector) Collect() (*Info, error) {
	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "model name") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return &Info{Label: c.Name(), Value: strings.TrimSpace(parts[1])}, nil
			}
		}
	}

	return &Info{Label: c.Name(), Value: "Unknown"}, nil
}

// lspci | grep -Ei 'VGA|3D|Display' | sed -E 's/.*\[AMD\/ATI\] .* \[(.*)\].*/AMD ATI \1/'
