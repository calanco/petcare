package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/calanco/petcare/server"
	"github.com/sirupsen/logrus"
)

var port string

func init() {
	flag.StringVar(&port, "port", "80", "Petcare listening port")
	flag.Parse()
}

func main() {
	petCareServer := server.Server()

	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), petCareServer))
}
