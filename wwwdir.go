package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "p", 4040, "The port of the web server.")
	flag.Parse()
}

func main() {
	fmt.Println(banner())
	fmt.Println("wwwdir @ http://github.com/spcoder/wwwdir\n")
	fmt.Printf("Listening on port %d. Press ctrl+c to stop...\n", port)
	exec.Command("cmd", "/C", "start", fmt.Sprintf("http://localhost:%d", port)).Start()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir("."))))
}

func banner() string {
	return `                             
                     _ _     
 _ _ _ _ _ _ _ _ _ _| |_|___ 
| | | | | | | | | | . | |  _|
|_____|_____|_____|___|_|_| 
`
}
