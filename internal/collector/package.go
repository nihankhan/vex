package collector

import (
	"strings"

	"github.com/nihankhan/vex/internal/utils"
)

type PackageCollector struct{}

func (p *PackageCollector) Name() string { return "Packages" }

func (p *PackageCollector) Collect() (*Info, error) {
	dpkgCount, err := utils.RunCmd(`dpkg -l | wc -l`)
	if err != nil {
		return nil, err
	}
	snapCount, err := utils.RunCmd(`snap list 2>/dev/null | wc -l`)
	if err != nil {
		snapCount = "0"
	}

	value := strings.TrimSpace(dpkgCount) + " (dpkg), " +
		strings.TrimSpace(snapCount) + " (snap)"

	return &Info{
		Label: p.Name(),
		Value: value,
	}, nil
}
