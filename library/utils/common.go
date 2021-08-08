package utils

import (
	"fmt"
	"net"
)

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func GetIpAddress() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err.Error())
	}

	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err.Error())
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Println(ip)
			// process IP address
		}
	}
}
