package collector

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nihankhan/vex/internal/utils"
)

type MemoryCollector struct{}

func (m *MemoryCollector) Name() string { return "Memory" }

func (m *MemoryCollector) Collect() (*Info, error) {
	out, err := utils.RunCmd(`free -m | awk '/Mem:/ {print $2 " " $3}'`)
	if err != nil {
		return nil, err
	}

	parts := strings.Fields(strings.TrimSpace(out))
	if len(parts) < 2 {
		return nil, fmt.Errorf("failed to parse memory")
	}

	total := parts[0]
	used := parts[1]

	_, _ = strconv.Atoi(total)
	_, _ = strconv.Atoi(used)

	value := fmt.Sprintf("%sMiB / %sMiB", used, total)

	return &Info{
		Label: m.Name(),
		Value: value,
	}, nil
}
