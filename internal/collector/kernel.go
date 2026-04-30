package collector

import (
	"os/exec"
	"strings"
)

type KernelCollector struct{}

func (k *KernelCollector) Name() string { return "Kernel" }

func (k *KernelCollector) Collect() (*Info, error) {
	args := `uname -r`

	cmd := exec.Command("bash", "-c", args)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	kernel := strings.TrimSpace(string(out))

	return &Info{
		Label: k.Name(),
		Value: kernel,
	}, nil
}
