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
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		StandardPrinter(ErrorRedColor, "Could not retrieve CPU usage details.")
		panic(err) //Exit upon error, below code must not be executed
	}
	usedPercent := fmt.Sprintf("%.2f", cpuPercent[0])
	ResultPrinter("CPU Usage: ", usedPercent+"%")

}

func GetCpuInfo() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		StandardPrinter(ErrorRedColor, "Could not retrieve CPU details.")
		panic(err) //Exit upon error, below code must not be executed
	}

	logical_cores_count, err := cpu.Counts(true)
	if err != nil {
		StandardPrinter(ErrorRedColor, "Could not retrieve number of logical cpu cores.")
		panic(err) //Exit upon error, below code must not be executed
	}

	physical_cores_count, err := cpu.Counts(false)
	if err != nil {
		StandardPrinter(ErrorRedColor, "Could not retrieve number of physical cpu cores.")
		panic(err) //Exit upon error, below code must not be executed
	}

	ResultPrinter("CPU Model: ", cpuInfo[0].ModelName)
	ResultPrinter("CPU Physical Cores: ", physical_cores_count)
	ResultPrinter("CPU Logical Cores: ", logical_cores_count)
}
