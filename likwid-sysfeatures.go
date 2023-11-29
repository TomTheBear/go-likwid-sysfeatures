package main

import (
	"fmt"

	sysfeatures "example.com/go-likwid-sysfeatures/sysfeatures"
)

func main() {

	err := sysfeatures.SysFeaturesInit()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	list, err := sysfeatures.SysFeaturesList()
	if err == nil {
		for _, sf := range list {
			fmt.Println(sf.Category, sf.Name, sf.DevtypeName, sf.Description, sf.ReadOnly, sf.WriteOnly)
		}
	}

	dev, err := sysfeatures.LikwidDeviceCreateByName("hwthread", 1)
	if err != nil {
		fmt.Println(err.Error())
		sysfeatures.SysFeaturesClose()
		return
	}
	fmt.Println(dev.Devname, dev.Devtype, dev.Id)

	val, err := sysfeatures.SysFeaturesGetDevice("base_freq", dev)
	if err == nil {
		fmt.Println("base_freq", val)
	} else {
		fmt.Println("Failed to get base_freq")
	}

	sysfeatures.LikwidDeviceDestroy(dev)

	sysfeatures.SysFeaturesClose()
}
