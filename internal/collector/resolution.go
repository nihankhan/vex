package collector

import (
	"strings"

	"github.com/nihankhan/vex/internal/utils"
)

type ResolutionCollector struct{}

func (r *ResolutionCollector) Name() string { return "Resolution" }

func (r *ResolutionCollector) Collect() (*Info, error) {
	out, err := utils.RunCmd(`xrandr | grep ' connected' | grep -o '[0-9]\+x[0-9]\+'`)
	if err != nil {
		return nil, err
	}

	resolutions := strings.Fields(strings.TrimSpace(out))

	value := strings.Join(resolutions, ", ")

	return &Info{
		Label: r.Name(),
		Value: value,
	}, nil
}
