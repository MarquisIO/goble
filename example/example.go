// example
package main

import (
	"fmt"

	goble "github.com/MarquisIO/goble"
)

func main() {

	// Enable connection through serial port
	hm, err := goble.New("/dev/tty.usbserial", 9600)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Get firmwire version
	fmt.Println(
		"Software Version:",
		hm.GetSoftwareVersion().Result,
	)

	//Last connected device
	fmt.Println(
		"Last connected device",
		hm.GetLastConnectedDeviceAddress().Param,
	)

	// Get Device name
	fmt.Println(
		"Device Name:",
		hm.GetDeviceName().Param,
	)

	//Set device name
	fmt.Println(
		"Setting device name...",
		hm.SetDeviceName("Marquis").Result,
	)

	//Needs to restart to changes take effect
	fmt.Println(
		"Restarting...",
		hm.Restart().Result,
	)

	//Now device name will be Marquis
	//Diff between result / param
	//Result is raw response from device and param is converted result
	devName := hm.GetDeviceName()
	fmt.Println(
		"Device Name with result:",
		devName.Result,
		"\nDevice Name with param:",
		devName.Param,
	)

	//Get mode by definition of BleMode
	fmt.Println(
		"Mode:",
		hm.GetMode().Param,
	)

	//Get bound mode by definition of BleBoundMode
	fmt.Println(
		"Bond mode:",
		hm.GetBondMode().Param,
	)

	//Get PIO Pin status
	fmt.Println(
		"PIO 2 status(0-LOW, 1-HIGH):",
		hm.GetPIO(goble.Pin2).Param,
	)

	//Set PIO Pin to HIGH
	fmt.Println(
		"Setting PIO Pin 2 to HIGH (0-LOW, 1-HIGH)",
		hm.SetPIO(goble.Pin2, goble.High).Result,
	)

	//Now check the status
	fmt.Println(
		"PIO 2 status (0-LOW, 1-HIGH):",
		hm.GetPIO(goble.Pin2).Param,
	)

	//EnableBeaconMode
	fmt.Println(hm.FactoryReset().Result)
	fmt.Println(hm.SetBeaconMode(goble.IBeaconEnable).Result)
	fmt.Println(hm.GetBeaconMode().Param)
	fmt.Println(hm.SetAdvertisingInterval(goble.Ms546).Result)
	fmt.Println(hm.Restart().Result)

	//Close
	hm.Close()
}
