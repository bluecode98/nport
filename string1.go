package main

import (
	//"fmt"
	//"golang.org/x/crypto/bcrypt"
	"os/exec"
	"io"
	"github.com/op/go-logging"
	"os"
	"time"
)

type Host struct {
	HostName string `json:"hostname"`
	OSName string `json:"osname"`
	OSVersion string `json:"osversion"`
	SystemType string `json:"systemtype"`
}


var log = logging.MustGetLogger("goserver")
var format = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} > %{level:s} %{message}`,
)
func makeShell(cmd string) error {
	cmdShell := exec.Command("sh", "-c", cmd)
	ppReader, err := cmdShell.StdoutPipe()
	if err != nil {
		log.Debug("create read pipe error", err.Error())
	}
	defer ppReader.Close()

	if err := cmdShell.Start(); err != nil {
		log.Debug(err.Error())
		return err
	}
	buffer := make([]byte, 10240)

	for {
		n, err := ppReader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Debug("pipi has closed")
				return err
			} else {
				log.Debug("read content failed")
			}
		}

		log.Debug(string(buffer[:n]))
	}
	return nil
}

func main(){
	// init log config
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	logging.SetBackend(backend2Formatter)

	//message := "Hello World!\n\tGod Bless You!\r"
	//key := "123456781234567812345678"
	//fmt.Printf(message)
	//log.Debug("key", key)
	//
	//hash, _ := bcrypt.GenerateFromPassword([]byte(message), bcrypt.DefaultCost)
	//log.Debug("password", string(hash))

	makeShell("pwd")

	time.Sleep(time.Duration(6)*time.Second)
}