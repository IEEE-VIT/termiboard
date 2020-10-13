package main

import (
	"flag"
	"fmt"
)

var (
	// Provisioned by ldflags
	version   string
	buildDate string
	commit    string
)

var versionString = fmt.Sprintf("v1.0 tag: %s, date: %s, commit: %s", version, buildDate, commit)

const (
	InfoOrangeColor    = "\033[1;34m%s\033[0m"
	BannerBlueColor    = "\033[1;36m%s\033[0m"
	WarningYellowColor = "\033[1;33m%s\033[0m"
	ErrorRedColor      = "\033[1;31m%s\033[0m"
	BoldWhite          = "\033[1;37m%s\033[0m"
	None               = "\033[0m%s\033[0m"
)

var (
	showCPUInfo             *bool
	showCPUUsage            *bool
	showRAM                 *bool
	showDisk                *bool
	showLocalIP             *bool
	showPublicIP            *bool
	showAll                 *bool
	showNTopRAM             *bool
	numberOfProcessesToShow *int
)

func main() {
	//init flags
	showCPUInfo = flag.Bool("cpu-info", false, "Show CPU information")
	showCPUUsage = flag.Bool("cpu-usage", false, "Show CPU usage")
	showRAM = flag.Bool("ram", false, "Show RAM usage")
	showDisk = flag.Bool("disk", false, "Show disk usage")
	showLocalIP = flag.Bool("local-ip", false, "Show local IP address")
	showPublicIP = flag.Bool("public-ip", false, "Show public IP address")
	showNTopRAM = flag.Bool("top-ram", false, "Show top n process that consume the most memory")
	numberOfProcessesToShow = flag.Int("n-proc", 5, "number of processes to show")
	showAll = flag.Bool("all", false, "Show all stats")
	flag.Parse()

	if flag.NFlag() == 0 {
		*showAll = true
	}
	var functionsWithConditions = []struct {
		condition bool
		function  func()
	}{
		{true, printBanner},
		{true, func() { StandardPrinter(WarningYellowColor, versionString) }},
		{*showCPUInfo, GetCpuInfo},
		{*showCPUUsage, GetCpuUsage},
		{*showRAM, GetRamUsage},
		{*showDisk, GetDiskUsage},
		{*showLocalIP, GetLocalIPAddress},
		{*showPublicIP, GetPublicIPAddress},
		{*showNTopRAM, func() { GetTopProcesses(*numberOfProcessesToShow) }},
	}
	for _, pair := range functionsWithConditions {
		if *showAll || pair.condition {
			pair.function()
		}
	}
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
