package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pocari/vscode-remote-go-template/src/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("%s%s\n", hello.Greet(), "!!"))
}

func main() {
	portNumber := "9080"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
