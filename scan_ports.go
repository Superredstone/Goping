package goping

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ScanPorts(IP string, min_port, max_port int, timeout time.Duration) ([]int, []int, error) {
	var OpenedPorts []int
	var ClosedPorts []int

	if min_port <= 0 {
		return OpenedPorts, ClosedPorts, fmt.Errorf("invalid port (min 1, max 65535)")
	}
	if max_port > 65535 {
		return OpenedPorts, ClosedPorts, fmt.Errorf("invalid port (min 1, max 65535)")
	}
	if net.ParseIP(IP) == nil {
		return OpenedPorts, ClosedPorts, fmt.Errorf("invalid ip address")
	}

	for i := min_port; i <= max_port; i++ {
		_, err := net.DialTimeout("tcp", IP+":"+strconv.Itoa(i), timeout)

		if err != nil {
			OpenedPorts = append(OpenedPorts, i)
		} else {
			ClosedPorts = append(ClosedPorts, i)
		}
	}

	return OpenedPorts, ClosedPorts, nil
}
