package hyper

import "strings"

func (c *Context) IP() []string {
	var ips []string
	remoteAddr := strings.Split(c.Req.RemoteAddr, ":")
	if len(remoteAddr) == 2 {
		ips = append(ips, remoteAddr[0])
	}
	if len(c.Req.Header.Get("X-Forwarded-For")) > 0 {
		xForwardedFor := strings.Split(c.Req.Header.Get("X-Forwarded-For"), ",")
		for _, ip := range xForwardedFor {
			ips = append(ips, ip)
		}
	}
	return ips
}
