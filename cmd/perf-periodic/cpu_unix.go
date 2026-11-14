//go:build unix

package main

import "syscall"

func cpuSeconds() (float64, error) {
	var usage syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &usage); err != nil {
		return 0, err
	}

	return float64(usage.Utime.Nano()+usage.Stime.Nano()) / 1e9, nil
}
