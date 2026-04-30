package collector

import (
	"strings"

	"github.com/nihankhan/vex/internal/utils"
)

type TerminalCollector struct{}

func (t *TerminalCollector) Name() string { return "Terminal" }

func (t *TerminalCollector) Collect() (*Info, error) {
	out, err := utils.RunCmd(`
		term=""
		pid=$PPID

		while [ "$pid" -ne 1 ]; do
			name=$(ps -p "$pid" -o comm= 2>/dev/null | xargs)

			case "$name" in
				# ignore shells + your app
				bash|zsh|fish|sh|dash|tmux|screen|sudo|su|gf)
					;;

				# terminals
				gnome-terminal*|gnome-terminal-server*)
					term="gnome-terminal"
					break
					;;

				kitty*)
					term="kitty"
					break
					;;

				alacritty*)
					term="alacritty"
					break
					;;

				wezterm*)
					term="wezterm"
					break
					;;

				konsole*)
					term="konsole"
					break
					;;

				tilix*)
					term="tilix"
					break
					;;

				xterm*)
					term="xterm"
					break
					;;

				code|code-insiders)
					term="vscode"
					break
					;;
			esac

			pid=$(ps -p "$pid" -o ppid= 2>/dev/null | xargs)

			[ -z "$pid" ] && break
		done

		echo "$term"
	`)
	if err != nil {
		return nil, err
	}

	term := strings.TrimSpace(out)

	if term == "" {
		term = "unknown"
	}

	return &Info{
		Label: t.Name(),
		Value: term,
	}, nil
}
