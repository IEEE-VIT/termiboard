package main

import (
	"github.com/shirou/gopsutil/cpu"
)

//func GetCpuTemperature() {
//	//Implement function to get the current temperature of CPU
//}
//
//func GetCpuUsage() {
//	//Implement function to get current usage of CPU
//}

func GetCpuInfo() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		StandardPrinter(ErrorRedColor, "Could not retrieve CPU details.")
	}
	ResultPrinter("CPU Model: ", cpuInfo[0].ModelName)
	ResultPrinter("CPU Cores: ", cpuInfo[0].Cores)

}
