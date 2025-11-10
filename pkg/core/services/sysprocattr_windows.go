//go:build windows

package services

import "syscall"

// setDetach is a no-op on Windows because syscall.SysProcAttr does not have Setsid.
func setDetach(attr *syscall.SysProcAttr) {
	// No detachment tweak needed; Windows uses different process group semantics.
}
