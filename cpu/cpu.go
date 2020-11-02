package cpu

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"termiboard"
	"time"
)

//func GetCpuTemperature() {
//	//Implement function to get the current temperature of CPU
//}
//

func GetCpuUsage() {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		main.StandardPrinter(main.ErrorRedColor, "Could not retrieve CPU usage details.")
		panic(err) //Exit upon error, below code must not be executed
	}
	usedPercent := fmt.Sprintf("%.2f", cpuPercent[0])
	main.ResultPrinter("CPU Usage: ", usedPercent+"%")

}

func GetCpuInfo() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		main.StandardPrinter(main.ErrorRedColor, "Could not retrieve CPU details.")
		panic(err) //Exit upon error, below code must not be executed
	}

	logicalCoresCount, err := cpu.Counts(true)
	if err != nil {
		main.StandardPrinter(main.ErrorRedColor, "Could not retrieve number of logical cpu cores.")
		panic(err) //Exit upon error, below code must not be executed
	}

	physicalCoresCount, err := cpu.Counts(false)
	if err != nil {
		main.StandardPrinter(main.ErrorRedColor, "Could not retrieve number of physical cpu cores.")
		panic(err) //Exit upon error, below code must not be executed
	}

	main.ResultPrinter("CPU Model: ", cpuInfo[0].ModelName)
	main.ResultPrinter("CPU Physical Cores: ", physicalCoresCount)
	main.ResultPrinter("CPU Logical Cores: ", logicalCoresCount)
}
