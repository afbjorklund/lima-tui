//go:build !windows

package services

import "syscall"

// setDetach configures SysProcAttr to detach the process on non-Windows platforms.
func setDetach(attr *syscall.SysProcAttr) {
	// On Unix-like systems, Setsid creates a new session to fully detach.
	attr.Setsid = true
}
