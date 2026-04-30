package collector

import (
	"bytes"
	"os/exec"
	"strings"
)

type GPUCollector struct{}

func (g *GPUCollector) Name() string { return "GPU" }

func (g *GPUCollector) Collect() (*Info, error) {
	args := `lspci | grep -Ei 'VGA|3D|Display' | sed -E 's/.*\[AMD\/ATI\] .* \[(.*)\].*/AMD ATI \1/'`
	cmd := exec.Command("bash", "-c", args)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	gpu := strings.TrimSpace(out.String())

	return &Info{
		Label: g.Name(),
		Value: gpu,
	}, nil
}
