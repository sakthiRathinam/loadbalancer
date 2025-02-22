package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	startHTTPReverseProxy()
}

type server struct {
	address string
	headers map[string][]string
}

func devServers() []server {
	return []server{
		{address: "10.0.0.2:8000"},
		{address: "10.0.0.3:8000"},
		{address: "10.0.0.5:8000"},
	}
}
func startHTTPReverseProxy() {
	fmt.Println("Starting HTTP Reverse Proxy on port 8080")
	http.HandleFunc("/", proxyHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	servers := devServers()
	server := pickRandomServer(servers)
	backendURL := fmt.Sprintf("http://%s%s", server.address, r.RequestURI)

	req, err := http.NewRequest(r.Method, backendURL, r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	addHeaders(req, server.headers)

	addHeaders(req, r.Header)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	copyHeader(w.Header(), resp.Header)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func pickRandomServer(servers []server) server {
	return servers[rand.Intn(len(servers))]
}

func copyHeader(dst, src http.Header) {
	for key, value := range src {
		dst.Set(key, value[0])
	}
}
func addHeaders(req *http.Request, headers map[string][]string) {
	for key, value := range headers {
		for _, v := range value {
			req.Header.Add(key, v)
		}
	}
}
