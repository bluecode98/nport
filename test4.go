package main

import (
	"syscall"
	"golang.org/x/sys/windows"
	"time"
	"os/exec"
	"io"
	"github.com/op/go-logging"

	"os"
)

var log = logging.MustGetLogger("goserver")
var format = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} > %{level:s} %{message}`,
)

func execMessage() error {
	cmdShell := exec.Command("cmd.exe")
	cmdShell.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
		CreationFlags: windows.STARTF_USESTDHANDLES,
	}

	ppReader, err := cmdShell.StdoutPipe()
	defer ppReader.Close()
	if err != nil {
		log.Debug("create read pipe error", err.Error())
	}

	ppWriter, err := cmdShell.StdinPipe()
	defer ppWriter.Close()
	if err != nil {
		log.Debug("create write pipe error", err.Error())
	}

	if err := cmdShell.Start(); err != nil {
		return err
	}


	// pipeReader
	go func() {
		buffer := make([]byte, 10240)

		for {
			// 从管道读取数据
			n, err := ppReader.Read(buffer)
			if err != nil {
				if err == io.EOF {
					log.Debug("pipi has closed")
					break
				} else {
					log.Debug("read content failed")
				}
			}

			log.Debug(string(buffer[:n]))
		}
	}()

	// pipeWriter
	go func() {
		for {
			time.Sleep(time.Duration(3)*time.Second)
			ppWriter.Write([]byte("time /t\r\n"))
		}
	}()

	time.Sleep(time.Duration(6)*time.Hour)
	return err
}

func main() {
	// init log config
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	logging.SetBackend(backend2Formatter)

	execMessage()

	time.Sleep(time.Duration(1)*time.Hour)
}
