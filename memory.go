package main

import (
	"fmt"
	"sort"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

func getReadableSize(sizeInBytes uint64) (readableSizeString string) {
	var (
		units = []string{"B", "KB", "MB", "GB", "TB", "PB"}
		size  = float64(sizeInBytes)
		i     = 0
	)
	for ; i < len(units) && size >= 1024; i++ {
		size = size / 1024
	}
	readableSizeString = fmt.Sprintf("%.2f %s", size, units[i])
	return
}

func GetRamUsage() {
	//Implement the function to fetch ram usage - Total Ram used , free ram available and used percentage
	if *showAll || *showRAM {
		m, err := mem.VirtualMemory()
		if err != nil {
			StandardPrinter(ErrorRedColor, "Could not retrieve RAM details.")
			panic(err) //Exit upon error, below code must not be executed
		}
		usedPercent := fmt.Sprintf("%f", m.UsedPercent)
		ResultPrinter("Ram Used: ", usedPercent+"%")
		ResultPrinter("Ram Available: ", getReadableSize(m.Free))
	}
}

// GetTopProcesses print out the top 5 process that are consuming most RAM
func GetTopProcesses() {
	if *showAll || *show5TopRAM {
		strOutput := ""
		processes, err := process.Processes()
		if err != nil {
			StandardPrinter(ErrorRedColor, "Could not retrieve running process list.")
			panic(err)
		}

		sort.Slice(processes, func(i, j int) bool {
			memoryPercentOfIthProcess, err := processes[i].MemoryPercent()
			if err != nil {
				StandardPrinter(ErrorRedColor, "Could not retrieve memory usage details.")
				panic(err)
			}
			memoryPercentOfJthProcess, err := processes[j].MemoryPercent()
			if err != nil {
				StandardPrinter(ErrorRedColor, "Could not retrieve memory usage details.")
				panic(err)
			}
			return memoryPercentOfIthProcess > memoryPercentOfJthProcess
		})

		for i := 0; i < 5; i++ {
			memoryPercentOfIthProcess, err := processes[i].MemoryPercent()
			if err != nil {
				StandardPrinter(ErrorRedColor, "Could not retrieve memory usage details.")
				panic(err)
			}
			strOutput += fmt.Sprintf("PID: %5d, memory %%: %2.1f\n", processes[i].Pid, memoryPercentOfIthProcess)
		}
		ResultPrinter("Top 5 processes by memory usage: \n", strOutput)
	}
}

func GetDiskUsage() {
	if *showAll || *showDisk {
		diskUsage, err := disk.Usage("/")
		if err != nil {
			StandardPrinter(ErrorRedColor, "Could not retrieve disk usage details.")
			panic(err) //Exit upon error, below code must not be executed
		}
		usedPercent := fmt.Sprintf("%.2f", diskUsage.UsedPercent)
		ResultPrinter("Disk Usage: ", usedPercent+"%")
		ResultPrinter("Disk Space Available: ", getReadableSize(diskUsage.Free))

	}
}
