package main

import (
  "io/ioutil"
  "net/http"
)

func GetPublicIPAddress() string {
  URL := "https://api.ipify.org" //Using Third Party Service to Ping
  resp, err := http.Get(URL) //Get the JSON Response
  if err != nil {
    StandardPrinter(ErrorRedColor, "Could not get response from " + URL)
    StandardPrinter(WarningYellowColor, "Check your Internet Connection")
    panic(err) //Exit upon error, below code must not be executed
  }
  defer resp.Body.Close() //The client must close the response body when finished with it
  IP, err := ioutil.ReadAll(resp.Body) //Reading all data from a io.Reader until EOF
  if err != nil {
    StandardPrinter(ErrorRedColor, "Could not find IPv4 Address")
    panic(err) //Exit upon error, below code must not be executed
  }

  return string(IP) //Cast []bByte to String and return
}

//func GetLocalIPAddress() {
//	//Implement to function to return the local IP Address of the pc.
//	// You will need to return a String.
//}
//
//func PortChecker() {
//	//Implement this function to take in a list of Strings (port numbers) , check if those ports are available or in use.
//	//Return a list of String of unavailable ports from the received list.
//}
