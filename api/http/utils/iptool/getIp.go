package iptool

import "net/http"

func GetIPAddressFromRequest(r *http.Request) string {
	// If proxies before API layer, find the ip at the start of the array and return
	return r.RemoteAddr
}
