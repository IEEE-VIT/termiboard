package main

import (
	"github.com/jaypipes/pcidb"
)

func GetPCIDevices() {
	ResultPrinter("PCI Bus Devices:", "")
	pci, err := pcidb.New()
	if err != nil {
		StandardPrinter(ErrorRedColor, "Error getting PCI Devices Data")
		panic(err)
	}

	for _, deviceClass := range pci.Classes {
		StandardPrinter(BoldWhiteColor, "\tDevice class: ["+string(deviceClass.Name)+"] "+string(deviceClass.ID))
		for _, deviceSubclass := range deviceClass.Subclasses {
			StandardPrinter(None, "\t\tDevice subclass:"+string(deviceSubclass.Name)+" "+string(deviceSubclass.ID))
			for _, programmingInterface := range deviceSubclass.ProgrammingInterfaces {
				StandardPrinter(BoldGreenColor, "\t\t\tProgramming interface"+string(programmingInterface.Name)+" "+string(programmingInterface.ID))
			}
		}
	}
}
