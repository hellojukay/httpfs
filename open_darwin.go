//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"os/exec"
)

func Open(url string) error {
	bin, err := exec.LookPath("open")
	if err != nil {
		return fmt.Errorf("can not find open, %s", err.Error())
	}
	return exec.Command(bin, url).Run()
}
