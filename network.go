package main

import "net"

//func GetPublicIPAddress() string {
//	//Implement the function to return the public address of the pc.
//	//Do not forget to handle errors if the pc isn't connected to a public network!
//	// You will need to return a String.
//}
//
func GetLocalIPAddress() {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        StandardPrinter(ErrorRedColor, "Could not retrieve your local ip")
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
		    ResultPrinter("Local ip addres: ", ipnet.IP.String())
            }
        }
    }

}
//
//func PortChecker() {
//	//Implement this function to take in a list of Strings (port numbers) , check if those ports are available or in use.
//	//Return a list of String of unavailable ports from the received list.
//}
