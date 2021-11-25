package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	// build the static main page
	go func() {
		fs := http.FileServer(http.Dir("assets/"))
		http.Handle("/hello/", http.StripPrefix("/hello/", fs))
		http.ListenAndServe(":8080", nil)
	} ()

	// driver is listening to port 80
	addr, _ := net.ResolveTCPAddr("tcp4", ":80")
	listener, err := net.ListenTCP("tcp", addr)
	defer listener.Close()
	if err != nil {
		return
	}

	containerRouting()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection error")
		}
		conn.Close()
	}

}

func containerRouting() {
	// environment variables defined in driver YAML
	hadoop := os.Getenv("HADOOP") + ":9870"
	spark := os.Getenv("SPARK") + ":8080"
	jupyter := os.Getenv("JUPYTER") + ":8888"
	sonarqube := os.Getenv("SONARQUBE") + ":9000"

	// map the internal port of each container to an external IP
	go routingHelper(hadoop, ":81")
	go routingHelper(spark, ":82")
	go routingHelper(jupyter, ":83")
	go routingHelper(sonarqube, ":84")

}

func routingHelper(ip string, port string) {
	url, err := url.Parse(ip)
	if err != nil {
		return
	}
	http.ListenAndServe(port, httputil.NewSingleHostReverseProxy(url))
}
