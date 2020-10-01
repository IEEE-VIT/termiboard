package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

//func GetCpuTemperature() {
//	//Implement function to get the current temperature of CPU
//}
//

func GetCpuUsage() {
	if *showAll || *showCPUUsage {
		cpuPercent, err := cpu.Percent(time.Second, false)
		if err != nil {
			StandardPrinter(ErrorRedColor, "Could not retrieve CPU usage details.")
			panic(err) //Exit upon error, below code must not be executed
		}
		usedPercent := fmt.Sprintf("%.2f", cpuPercent[0])
		ResultPrinter("CPU Usage: ", usedPercent+"%")
	}
}

func GetCpuInfo() {
	if *showAll || *showCPUInfo {
		cpuInfo, err := cpu.Info()
		if err != nil {
			StandardPrinter(ErrorRedColor, "Could not retrieve CPU details.")
			panic(err) //Exit upon error, below code must not be executed
		}

		ResultPrinter("CPU Model: ", cpuInfo[0].ModelName)
		ResultPrinter("CPU Cores: ", cpuInfo[0].Cores)

	}
}
