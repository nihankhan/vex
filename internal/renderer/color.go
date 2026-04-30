package renderer

import (
	"fmt"
	"strings"
)

type Color string

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	White  Color = "\033[37m"
	Bold   Color = "\033[1m"
)

// RGB color structure
type RGB struct {
	R, G, B int
}

// Convert RGB → ANSI escape code
func rgb(r, g, b int) Color {
	return Color(fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b))
}

// clean integer lerp (recommended)
func lerpInt(a, b int, t float64) int {
	return int(float64(a) + float64(b-a)*t)
}

func Colorize(text string, color Color) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}

// GenerateGradient creates smooth color steps
func GenerateGradient(start, end RGB, steps int) []Color {
	if steps <= 0 {
		return nil
	}

	colors := make([]Color, steps)

	for i := 0; i < steps; i++ {
		t := float64(i) / float64(steps-1)

		r := lerpInt(start.R, end.R, t)
		g := lerpInt(start.G, end.G, t)
		b := lerpInt(start.B, end.B, t)

		colors[i] = rgb(r, g, b)
	}

	return colors
}

func infoColor(label string) Color {
	switch label {
	case "OS":
		return Cyan

	case "Host":
		return Yellow

	case "Kernel":
		return Green

	case "Uptime":
		return Red

	case "CPU":
		return Blue

	case "GPU":
		return Purple

	case "Memory":
		return Cyan

	case "Packages":
		return Yellow

	case "Shell":
		return Green

	case "Terminal":
		return Blue

	case "Resolution":
		return Purple

	case "DE":
		return Cyan

	case "WM":
		return Yellow

	case "WM Theme":
		return Green

	case "Theme":
		return Blue

	case "Icons":
		return Purple

	default:
		return White
	}
}

func colorBlock(color int) string {
	return fmt.Sprintf("\033[48;5;%dm  \033[0m", color)
}

func paletteBar() string {
	colors := []int{
		0, 8, // black
		1, 9, // red
		2, 10, // green
		3, 11, // yellow
		4, 12, // blue
		5, 13, // magenta
		6, 14, // cyan
		7, 15, // white
	}

	var out strings.Builder

	for _, c := range colors {
		out.WriteString(colorBlock(c))
	}

	return out.String()
}

func BoldColorize(text string, color Color) string {
	return fmt.Sprintf("%s%s%s%s", Bold, color, text, Reset)
}
