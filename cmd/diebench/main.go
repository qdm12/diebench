package main

import "syscall"

func main() {

}

func setProcessAffinityMask(h syscall.Handle, mask uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(syscall.NewLazyDLL("kernel32.dll").NewProc("SetProcessAffinityMask").Addr(), 2, uintptr(h), mask, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
