package collector

import (
	"os/exec"
	"strings"
)

type UpTimeCollector struct{}

func (u *UpTimeCollector) Name() string { return "Uptime" }

func (u *UpTimeCollector) Collect() (*Info, error) {
	args := `uptime -p`

	cmd := exec.Command("bash", "-c", args)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	uptime := strings.TrimSpace(string(out))

	return &Info{
		Label: u.Name(),
		Value: uptime,
	}, nil
}
