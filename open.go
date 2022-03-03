//go:build linux
// +build linux

package main

import (
	"fmt"
	"os/exec"
)

func Open(url string) error {
	bin, err := exec.LookPath("xdg-open")
	if err != nil {
		return fmt.Errorf("can not find xdg-open, %s", err.Error())
	}
	return exec.Command(bin, url).Run()
}
