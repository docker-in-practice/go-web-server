package main

import (
	"fmt"
	"net/http"
	"flag"
    "io/ioutil"
)

//Server port
var port string

func check(msg string, e error) {
    if e != nil {
        panic(msg + e.Error())
    }
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	page, err := ioutil.ReadFile("page.html")
	check("defaultHandler:", err)
	fmt.Fprintln(w, string(page))
}

func handle_args() {
	flag.StringVar(&port, "port","8080","TCP port to serve on")
	flag.Parse()
}

func main() {

	handle_args()

	http.HandleFunc("/", defaultHandler)

	fmt.Println("Serving from port: " + port)
	err := http.ListenAndServe(":" + port, nil)
	check("ListenAndServe: ", err)
}
