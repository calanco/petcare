package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/calanco/petcare/server"
)

var port string

func init() {
	flag.StringVar(&port, "port", "80", "Petcare listening port")
	flag.Parse()
}

func main() {
	petCareServer := server.Server()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), petCareServer))
}
