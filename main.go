package main

import (
	"flag"
	"fmt"
)

const (
	InfoOrangeColor    = "\033[1;34m%s\033[0m"
	BannerBlueColor    = "\033[1;36m%s\033[0m"
	WarningYellowColor = "\033[1;33m%s\033[0m"
	ErrorRedColor      = "\033[1;31m%s\033[0m"
	BoldGreenColor     = "\033[1;32m%s\033[0m"
	BoldWhiteColor     = "\033[1;37m%s\033[0m"
	None               = "\033[0m%s\033[0m"
)

var (
	showCPUInfo    *bool
	showCPUUsage   *bool
	showRAM        *bool
	showDisk       *bool
	showLocalIP    *bool
	showPublicIP   *bool
	showAll        *bool
	show5TopRAM    *bool
	showPCIDevices *bool
)

func main() {
	//init flags
	showCPUInfo = flag.Bool("cpu-info", false, "Show CPU information")
	showCPUUsage = flag.Bool("cpu-usage", false, "Show CPU usage")
	showRAM = flag.Bool("ram", false, "Show RAM usage")
	showDisk = flag.Bool("disk", false, "Show disk usage")
	showLocalIP = flag.Bool("local-ip", false, "Show local IP address")
	showPublicIP = flag.Bool("public-ip", false, "Show public IP address")
	show5TopRAM = flag.Bool("top5-ram", false, "Show top 5 process that consume the most memory")
	showPCIDevices = flag.Bool("pci-devices", false, "Show all PCI Devices")
	showAll = flag.Bool("all", false, "Show all stats")
	flag.Parse()

	if flag.NFlag() == 0 {
		*showAll = true
	}

	printBanner()
	StandardPrinter(WarningYellowColor, "v1.0")
	GetCpuInfo()
	GetCpuUsage()
	GetRamUsage()
	GetDiskUsage()
	GetLocalIPAddress()
	GetPublicIPAddress()
	GetTopProcesses()
	GetPCIDevices()
}

func printBanner() {
	fmt.Printf(BannerBlueColor, " _                      _ _                         _\n| |_ ___ _ __ _ __ ___ (_) |__   ___   __ _ _ __ __| |\n| __/ _ \\ '__| '_ ` _ \\| | '_ \\ / _ \\ / _` | '__/ _` |\n| ||  __/ |  | | | | | | | |_) | (_) | (_| | | | (_| |\n \\__\\___|_|  |_| |_| |_|_|_.__/ \\___/ \\__,_|_|  \\__,_|")
	fmt.Println("")
}

func StandardPrinter(color string, message string) {
	fmt.Printf(color, message)
	fmt.Println("")
}

func ResultPrinter(title string, result interface{}) {
	fmt.Printf(InfoOrangeColor, title)
	fmt.Printf("%v", result)
	fmt.Println("")
}
