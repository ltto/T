package io

import (
	"os"
	"syscall"
)

func RemoveSyscallErrno(err error) (syscall.Errno, bool) {
	perr, ok := err.(*os.PathError)
	if !ok {
		return 0, false
	}
	errno, ok := perr.Err.(syscall.Errno)
	return errno, ok
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	} else {
		return !os.IsNotExist(err)
	}
}
