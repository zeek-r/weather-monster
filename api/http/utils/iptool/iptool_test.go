package iptool

import (
	"bytes"
	"net/http"
	"testing"
)

func TestIPTool(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost", bytes.NewBuffer([]byte{}))
	remoteAddr := "http://localhost"
	request.RemoteAddr = remoteAddr

	ip := GetIPAddressFromRequest(request)

	if ip != remoteAddr {
		t.Errorf("Incorrect value for ip, Got: %s Want %s", ip, remoteAddr)
	}
}
