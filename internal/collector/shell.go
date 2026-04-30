package collector

import (
	"os"
	"strings"

	"github.com/nihankhan/vex/internal/utils"
)

type ShellCollector struct{}

func (s *ShellCollector) Name() string { return "Shell" }

func (s *ShellCollector) Collect() (*Info, error) {
	shellPath := os.Getenv("SHELL")
	shell := "unknown"

	if strings.Contains(shellPath, "bash") {
		out, _ := utils.RunCmd(`bash --version | head -n1`)
		shell = parseBashVersion(out)
	} else if strings.Contains(shellPath, "zsh") {
		out, _ := utils.RunCmd(`zsh --version`)
		shell = strings.TrimSpace(out)
	} else if strings.Contains(shellPath, "fish") {
		out, _ := utils.RunCmd(`fish --version`)
		shell = strings.TrimSpace(out)
	} else {
		shell = strings.TrimSpace(shellPath)
	}

	return &Info{
		Label: s.Name(),
		Value: shell,
	}, nil
}

func parseBashVersion(out string) string {
	parts := strings.Fields(out)
	for i, p := range parts {
		if strings.Contains(p, "version") && i+1 < len(parts) {
			return "bash " + strings.Split(parts[i+1], "(")[0]
		}
	}
	return "bash unknown"
}
