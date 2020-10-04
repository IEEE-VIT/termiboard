// +build !windows

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jaypipes/ghw"
)

func GetPCIDevices() {

	ResultPrinter("Host PCI devices:", "")

	//Create a Tab writer to STDout with padding of 3, space between column and Left Align column
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	//Always Flush the Tab Writer because it is line buffered
	defer w.Flush()

	//TODO: Implement StandardPrinter for tabwriter
	//Print Header Row of the table
	fmt.Fprintln(w, "Class\tManufacturer\tName\t")

	////TODO: Implement StandardPrinter for tabwriter
	//Print Seperator for Header and Data of the table
	fmt.Fprintln(w, "-----\t------------\t----\t")

	pci, err := ghw.PCI()
	if err != nil {
		fmt.Printf("Error getting PCI info: %v", err)
	}

	devices := pci.ListDevices()
	if len(devices) == 0 {
		fmt.Printf("error: could not retrieve PCI devices\n")
		return
	}

	for _, device := range devices {
		vendor := device.Vendor
		vendorName := vendor.Name
		if len(vendor.Name) > 20 {
			//Product Name Shorten
			vendorName = string([]byte(vendorName)[0:17]) + "..."
		}
		product := device.Product
		productName := product.Name
		if len(product.Name) > 40 {
			//Product Name Shorten
			productName = string([]byte(productName)[0:37]) + "..."
		}
		//TODO: Implement StandardPrinter for tabwriter
		//Print Data Row of the table
		fmt.Fprintln(w, device.Class.Name+"\t"+vendorName+"\t"+productName+"\t")
	}

}
