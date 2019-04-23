package util

import (
	"net"
	"net/http"
)

func MustGetIp(r *http.Request) (ip, port string) {
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		panic(err)
	}
	return ip, port
}
