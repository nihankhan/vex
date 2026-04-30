/*
Copyright © 2026 Nihan Khan nihan.khan@outlook.com
*/
package main

import (
	"github.com/nihankhan/vex/internal/app"
	"github.com/nihankhan/vex/internal/collector"
	"github.com/nihankhan/vex/internal/renderer"
)

func main() {
	collectors := []collector.Collector{
		// ── SYSTEM ──
		&collector.IdentityCollector{},
		&collector.OSCollector{},
		&collector.HostCollector{},
		&collector.KernelCollector{},
		&collector.UpTimeCollector{},

		// ── SOFTWARE / PACKAGE ──
		&collector.PackageCollector{},
		&collector.ShellCollector{},
		&collector.TerminalCollector{},

		// ── DISPLAY / UI ──
		&collector.ResolutionCollector{},
		&collector.DECollector{},
		&collector.WMCollector{},
		&collector.WMThemeCollector{},
		&collector.ThemeCollector{},
		&collector.IconCollector{},

		// ── HARDWARE ──
		&collector.CPUCollector{},
		&collector.GPUCollector{},
		&collector.MemoryCollector{},
	}

	// for _, c := range collectors {
	// 	info, err := c.Collect()
	// 	if err != nil {
	// 		fmt.Printf("%s: error: %v\n", c.Name(), err)
	// 		continue
	// 	}

	// 	fmt.Printf("%-10s : %s\n", info.Label, info.Value)
	// }

	detectedOS, infoLines := app.BuildSystemInfo(collectors)

	r := renderer.NewRenderer(detectedOS, infoLines)
	r.Render()
}
