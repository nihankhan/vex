package app

import (
	"fmt"
	"strings"

	"github.com/nihankhan/vex/internal/collector"
)

func BuildSystemInfo(collectors []collector.Collector) (string, []string) {
	var infoLines []string
	detectedOS := "ubuntu"

	for _, c := range collectors {
		info, err := c.Collect()
		if err != nil {
			continue
		}

		if strings.ToLower(info.Label) == "os" {
			value := strings.ToLower(info.Value)

			switch {
			case strings.Contains(value, "arch"):
				detectedOS = "arch"
			case strings.Contains(value, "debian"):
				detectedOS = "debian"
			default:
				detectedOS = "ubuntu"
			}
		}

		infoLines = append(infoLines,
			fmt.Sprintf("%-12s %s", info.Label+":", info.Value),
		)
	}

	return detectedOS, infoLines
}
