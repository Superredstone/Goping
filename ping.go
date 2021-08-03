package goping

import (
	"fmt"
	"net"
	"time"
)

func Ping(IP, port string, timeout time.Duration) (bool, error) {
	if net.ParseIP(IP) == nil {
		return false, fmt.Errorf("invalid ip address")
	}

	_, err := net.DialTimeout("tcp", IP+":"+port, timeout)
	if err != nil {
		return true, nil
	} else {
		return true, nil
	}
}
