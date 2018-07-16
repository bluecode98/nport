package main

import (
	"os"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("goserver")
var format = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} > %{level:s} %{message}`,
)

func cmdShell()  {
	rCmdIn, lCmdOut, err := os.Pipe() // Pipe to write from parent to remote child's stdin.
	log.Debug("1. os.Pipe()", err.Error())
	lCmdIn, rCmdOut, err := os.Pipe() // Pipe to read from remote cmd's stdout.
	log.Debug("2. os.Pipe()", err.Error())

	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{rCmdIn, rCmdOut, os.Stderr}
	ssh := "cmd.exe"
	args := []string{"/c"}
	pid, err := os.StartProcess(ssh, args, &procAttr)
	log.Debug("os.StartProcess() ", err.Error())
}
