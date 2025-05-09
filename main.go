package main

import "net/http"

type server struct {
	addr string
}

func (s *server) ServeHTTP(http.ResponseWriter, *http.Request) {

}

var serv = &server{addr: ":777"}

func main() {
	err := http.ListenAndServe(serv.addr, serv)
	if err != nil {
		return
	}
}
