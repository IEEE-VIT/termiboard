package main

import (
	"net"
)

//func GetPublicIPAddress() {
//	//Implement the function to return the public address of the pc.
//	//Do not forget to handle errors if the pc isn't connected to a public network!
//	// You will need to return a String.
//}
//
//func GetLocalIPAddress() {
//	//Implement to function to return the local IP Address of the pc.
//	// You will need to return a String.
//}

// Returns a list of unavailable ports from the received list.
func PortChecker(ports []string) []string {
	var unavailable []string
	for _, port := range ports {
		if isPortUnavailable(port) {
			unavailable = append(unavailable, port)
		}
	}

	return unavailable
}

func isPortUnavailable(port string) bool {
	ln, err := net.Listen("tcp", ":"+port)
	if ln != nil {
		defer ln.Close()
	}

	return err != nil
}
