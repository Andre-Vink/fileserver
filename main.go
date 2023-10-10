package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", ":3000", "The port this file server is listening for connections.")
	flag.Parse()
	
	fmt.Printf("File server started at port %v.", *port)
	log.Fatal(http.ListenAndServe(*port, http.FileServer(http.Dir("."))))
}