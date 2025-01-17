//go:build !windows
// +build !windows

package shell

import (
	"errors"
	"os"

	"golang.org/x/term"
)

// GetTerminalWidth checks if input is a terminal
func GetTerminalWidth(tty interface{}) (width int, err error) {
	var (
		fh     *os.File
		assert bool
	)

	if fh, assert = tty.(*os.File); !assert {
		width = standardTermWidth
		err = errors.New("TTY is not a files")
		return
	}

	width, _, err = term.GetSize(int(fh.Fd()))

	return
}
