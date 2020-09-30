package main

func GetPublicIPAddress() {
	//Implement the function to return the public address of the pc.
	//Do not forget to handle errors if the pc isn't connected to a public network!
	// You will need to return a String.
}

func GetLocalIPAddress() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

func PortChecker() {
	//Implement this function to take in a list of Strings (port numbers) , check if those ports are available or in use.
	//Return a list of String of unavailable ports from the received list.
}
