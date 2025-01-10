package main

import (
	"fmt"
	"net"
)

// GetIPsInCIDR returns all IP addresses in a given CIDR block.
func GetIPsInCIDR(cidrBlock string) ([]string, error) {
	// Parse the CIDR block
	_, network, err := net.ParseCIDR(cidrBlock)
	if err != nil {
		return nil, err
	}

	// Calculate the list of IPs in the CIDR block
	ipAddresses := []string{}
	for currentIP := network.IP.Mask(network.Mask); network.Contains(currentIP); {
		ipAddresses = append(ipAddresses, currentIP.String())

		// Increment the IP address by 1
		for byteIndex := len(currentIP) - 1; byteIndex >= 0; byteIndex-- {
			currentIP[byteIndex]++
			if currentIP[byteIndex] != 0 {
				break
			}
		}
	}

	return ipAddresses, nil
}

func main() {
	cidrBlock := "192.168.1.0/29"
	ipAddresses, err := GetIPsInCIDR(cidrBlock)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print all IP addresses in the CIDR block
	for _, ip := range ipAddresses {
		fmt.Println(ip)
	}
}
