package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to echo.")

	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)

	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}

	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
	} else {
		userIP := net.ParseIP(ip)
		if userIP == nil {
			fmt.Fprintf(w, "userip: %q is not IP:port\n", r.RemoteAddr)
		}
		fmt.Fprintf(w, "IP: %s\n", userIP)
		fmt.Fprintf(w, "Port: %s\n", port)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
