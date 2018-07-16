package main

import (
	"fmt"
	//"html"
	"net/http"
	"os"
	//"runtime"
	"github.com/kavu/go_reuseport"
)

func main() {
	listener, err := reuseport.NewReusablePortListener("tcp", "0.0.0.0:80")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(os.Getgid())
		fmt.Println(r.RemoteAddr)
		//r.Header.Get("Remote_addr")
		fmt.Fprintf(w, "Hello, [%s]", r.RemoteAddr)
	})

	panic(server.Serve(listener))
}
