package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

//func GetCpuTemperature() {
//	//Implement function to get the current temperature of CPU
//}

func GetCpuUsage() {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		StandardPrinter(ErrorRedColor, "Could not retrieve CPU usage details.")
	}
	usedPercent := fmt.Sprintf("%f", cpuPercent[0])
	ResultPrinter("CPU Usage: ", usedPercent+"%")
}

func GetCpuInfo() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		StandardPrinter(ErrorRedColor, "Could not retrieve CPU details.")
	}
	ResultPrinter("CPU Model: ", cpuInfo[0].ModelName)
	ResultPrinter("CPU Cores: ", cpuInfo[0].Cores)

}
