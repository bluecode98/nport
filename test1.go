package main

import (
	"os/exec"
	"bytes"
	"fmt"
	"log"
)

func main() {
	cmd := exec.Command("ls", "-a", "-l")
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
