package external

import (
	"syscall"
)

type (
	ErrNo         int
	ErrNoExtended int

	Error struct {
		Code         ErrNo
		ExtendedCode ErrNoExtended
		SystemErrno  syscall.Errno
		Err          string
	}
)

func (err Error) Error() string {
	return "error sqlite"
}
