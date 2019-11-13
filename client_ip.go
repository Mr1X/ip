package ip

import (
	"net"
	"net/http"
	"strings"
)

// ClientIP get client ip
func ClientIP(r *http.Request) string {
	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP != "" {
		clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
		if clientIP == "" {
			clientIP = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
		}
		if clientIP != "" {
			return clientIP
		}
	}

	if addr := r.Header.Get("X-Appengine-Remote-Addr"); addr != "" {
		return addr
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
