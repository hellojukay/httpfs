//go:build windows
// +build windows

package main

import (
	"fmt"
	"os/exec"
	"time"
)

// start up cmd.exe then run start command to open default browser
func Open(url string) error {
	bin, err := exec.LookPath("cmd")
	if err != nil {
		return fmt.Errorf("can not find start, %s", err.Error())
	}
	c := exec.Cmd{
		Path: bin,
	}
	w, err := c.StdinPipe()
	if err != nil {
		return err
	}
	c.Start()
	w.Write([]byte(fmt.Sprintf("start %s\n", url)))
	time.Sleep(time.Second)
	w.Close()
	return c.Wait()
}
