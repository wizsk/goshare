package main

import (
	"net"
	"os"
	"strconv"
)

// it looks for form start to including end
func findFreePort(start, end int) string {
	if p := os.Getenv("PORT"); p != "" {
		if c, err := net.Listen("tcp", "0.0.0.0:"+p); err == nil &&
			c.Close() == nil {
			return p
		}
	}

	for i := start; i <= end; i++ {
		p := strconv.Itoa(i)
		if c, err := net.Listen("tcp", "0.0.0.0:"+p); err == nil &&
			c.Close() == nil {
			return p
		}
	}

	return ""
}
