package main

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
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
		usedMessage := fmt.Sprintf(
			"%s (%.2f%%)",
			getReadableSize(m.Used),
			m.UsedPercent,
		)
		ResultPrinter("Ram Total: ", getReadableSize(m.Total))
		ResultPrinter("Ram Used: ", usedMessage)
		ResultPrinter("Ram Available: ", getReadableSize(m.Available))
		ResultPrinter("Ram Free: ", getReadableSize(m.Free))
	}
}

//func GetTopProcesses() {
//	//Implement this function to return top 5 process that are consuming the most ram
//}
//
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
