package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func GetRamUsage() {
	//Implement the function to fetch ram usage - Total Ram used , free ram available and used percentage
	m, err := mem.VirtualMemory()
	if err != nil {
		StandardPrinter(ErrorRedColor, "Could not retrieve RAM details.")
	}
	usedPercent := fmt.Sprintf("%f", m.UsedPercent)
	ResultPrinter("Ram Used: ", usedPercent+"%")
	ResultPrinter("Ram Available: ", m.Used)
}

//func GetTopProcesses() {
//	//Implement this function to return top 5 process that are consuming the most ram
//}
//
//func GetDiskUsage() {
//	//Implement this function to display disk usage
//}
