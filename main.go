package main

import (
	"flag"
	"fmt"
	"termiboard/cpu"
	"termiboard/memory"
	"termiboard/network"
	"termiboard/utils"
)

var (
	// Provisioned by ldflags
	version   string
	buildDate string
	commit    string
)

var versionString = fmt.Sprintf("v1.0 tag: %s, date: %s, commit: %s", version, buildDate, commit)

var (
	showCPUInfo  *bool
	showCPUUsage *bool
	showRAM      *bool
	showDisk     *bool
	showLocalIP  *bool
	showPublicIP *bool
	showAll      *bool
	show5TopRAM  *bool
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
	showAll = flag.Bool("all", false, "Show all stats")
	flag.Parse()

	if flag.NFlag() == 0 {
		*showAll = true
	}
	var functionsWithConditions = []struct {
		condition bool
		function  func()
	}{
		{true, utils.PrintBanner},
		{true, func() { utils.StandardPrinter(utils.WarningYellowColor, versionString) }},
		{*showCPUInfo, cpu.GetCpuInfo},
		{*showCPUUsage, cpu.GetCpuUsage},
		{*showRAM, memory.GetRamUsage},
		{*showDisk, memory.GetDiskUsage},
		{*showLocalIP, network.GetLocalIPAddress},
		{*showPublicIP, network.GetPublicIPAddress},
		{*show5TopRAM, memory.GetTopProcesses},
	}
	for _, pair := range functionsWithConditions {
		if *showAll || pair.condition {
			pair.function()
		}
	}
}
