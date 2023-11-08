package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	cmd := request.URL.Query().Get("cmd")
	runCmd(cmd)
	test(cmd)
}

func test(cmd string) {
	runCmd(cmd)
}

func runCmd(cmd string) {
	exec.Command(cmd).Run()
}
