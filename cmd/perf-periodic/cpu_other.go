//go:build !unix

package main

import "errors"

func cpuSeconds() (float64, error) { return 0, errors.New("process CPU sampling requires a Unix host") }
