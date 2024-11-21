package ansi

import (
	"net/url"
)

// NotifyWorkingDirectory returns a sequence for notifying the program's
// current working directory.
//
//	OSC 7 ; Pt BEL
//
// Where Pt is a URL in the format "file://[host]/[path]".
// Set host to "localhost" if this is a path on the local computer.
//
// See: https://wezfurlong.org/wezterm/shell-integration.html#osc-7-escape-sequence-to-set-the-working-directory
// See: https://iterm2.com/documentation-escape-codes.html#:~:text=RemoteHost%20and%20CurrentDir%3A-,OSC%207,-%3B%20%5BPs%5D%20ST
func NotifyWorkingDirectory(host string, path string) string {
	u := &url.URL{
		Scheme: "file",
		Host:   host,
		Path:   path,
	}
	return "\x1b]7;" + u.String() + "\x07"
}
