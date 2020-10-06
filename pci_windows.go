// Reference Documentation: https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-pnpentity
package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/StackExchange/wmi"
)

//Parameters are selected such that they match with Linux Counterpart as well
type PnPEntity struct {
	PNPClass     string
	Name         string
	Manufacturer string
}

func GetPCIDevices() {

	ResultPrinter("Host PCI devices:", "")

	//Create a Tab writer to STDout with padding of 3, space between column and Left Align column
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	//Always Flush the Tab Writer because it is line buffered
	defer w.Flush()

	//Create a slice for holding info about PCIDevices with PnPEntity properties
	var PCIDevices []PnPEntity
	//Query to select only the PCI Devices from Win32_PnPEntity
	err := wmi.Query("Select * from Win32_PnPEntity where PNPDeviceID like '%PCI%'", &PCIDevices)
	if err != nil {
		StandardPrinter(ErrorRedColor, "Could not retrieve PCI Devices")
		panic(err)
		return
	}

	//TODO: Implement StandardPrinter for tabwriter
	//Print Header Row of the table
	fmt.Fprintln(w, "Class\tManufacturer\tName\t")

	//TODO: Implement StandardPrinter for tabwriter
	//Print Seperator for Header and Data of the table
	fmt.Fprintln(w, "-----\t------------\t----\t")

	for _, PCIDevice := range PCIDevices {
		//TODO: Implement StandardPrinter for tabwriter
		//Print Data Row of the table
		fmt.Fprintln(w, PCIDevice.PNPClass+"\t"+PCIDevice.Manufacturer+"\t"+PCIDevice.Name+"\t")
	}

}
