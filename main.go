package main

import "fmt"

const (
	InfoOrangeColor    = "\033[1;34m%s\033[0m"
	BannerBlueColor    = "\033[1;36m%s\033[0m"
	WarningYellowColor = "\033[1;33m%s\033[0m"
	ErrorRedColor      = "\033[1;31m%s\033[0m"
)

func main() {
	printBanner()
	StandardPrinter(WarningYellowColor, "v1.0")
	GetCpuInfo()
	GetCpuUsage()
	GetRamUsage()
	GetDiskUsage()
	GetLocalIPAddress()
	GetPublicIPAddress()
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
