package cpu

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"termiboard/utils"
	"time"
)

//func GetCpuTemperature() {
//	//Implement function to get the current temperature of CPU
//}
//

func GetCpuUsage() {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		utils.StandardPrinter(utils.ErrorRedColor, "Could not retrieve CPU usage details.")
		utils.PrintVerbosePanic(err) //Exit upon error, below code must not be executed
	}
	usedPercent := fmt.Sprintf("%.2f", cpuPercent[0])
	utils.ResultPrinter("CPU Usage: ", usedPercent+"%")

}

func GetCpuInfo() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		utils.StandardPrinter(utils.ErrorRedColor, "Could not retrieve CPU details.")
		utils.PrintVerbosePanic(err) //Exit upon error, below code must not be executed
	}

	logicalCoresCount, err := cpu.Counts(true)
	if err != nil {
		utils.StandardPrinter(utils.ErrorRedColor, "Could not retrieve number of logical cpu cores.")
		utils.PrintVerbosePanic(err) //Exit upon error, below code must not be executed
	}

	physicalCoresCount, err := cpu.Counts(false)
	if err != nil {
		utils.StandardPrinter(utils.ErrorRedColor, "Could not retrieve number of physical cpu cores.")
		utils.PrintVerbosePanic(err) //Exit upon error, below code must not be executed
	}

	utils.ResultPrinter("CPU Model: ", cpuInfo[0].ModelName)
	utils.ResultPrinter("CPU Physical Cores: ", physicalCoresCount)
	utils.ResultPrinter("CPU Logical Cores: ", logicalCoresCount)
}
