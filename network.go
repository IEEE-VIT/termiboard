package main

import (
	"io/ioutil"
	"net"
	"net/http"
)

//func GetPublicIPAddress() string {
//	//Implement the function to return the public address of the pc.
//	//Do not forget to handle errors if the pc isn't connected to a public network!
//	// You will need to return a String.
//}
//
func GetLocalIPAddress() {
	if *showAll || *showLocalIP {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			StandardPrinter(ErrorRedColor, "Unfortunately it is not possible to get your local IP")
		}
		for _, address := range addrs {
			// check the address type and if it is not a loopback the display it
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ResultPrinter("Local IP Address: ", ipnet.IP.String())
				}
			}
		}

	}
}

func GetPublicIPAddress() {
	if *showAll || *showPublicIP {
		URL := "https://api.ipify.org" //Using Third Party Service to Ping
		resp, err := http.Get(URL)     //Get the JSON Response
		if err != nil {
			StandardPrinter(ErrorRedColor, "Could not get response from "+URL)
			StandardPrinter(WarningYellowColor, "Check your Internet Connection")
			panic(err) //Exit upon error, below code must not be executed
		}
		defer resp.Body.Close()              //The client must close the response body when finished with it
		IP, err := ioutil.ReadAll(resp.Body) //Reading all data from a io.Reader until EOF
		if err != nil {
			StandardPrinter(ErrorRedColor, "Could not find IPv4 Address")
			panic(err) //Exit upon error, below code must not be executed
		}
		ResultPrinter("Public IPv4 Address: ", string(IP)) //Cast []bByte to String and return
	}
}

//
//func PortChecker() {
//	//Implement this function to take in a list of Strings (port numbers) , check if those ports are available or in use.
//	//Return a list of String of unavailable ports from the received list.
//}
