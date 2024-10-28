//go:build !windows

package main

import (
	"os"
	"syscall"
)

func detachProcess() {
	syscall.Umask(0)
	_ = os.Chdir("/")
}
