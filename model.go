package goble

import (
	"github.com/MarinX/serial"
)

// Test with HM-10 - Firmware v540

const (

	// Test connection
	test = "AT"

	// Device MAC Address
	getDeviceAddress = "AT+ADDR?"

	// Device Advertising Interval
	getAdvertisingInt = "AT+ADVI?"
	setAdvertisingInt = "AT+ADVI%s"

	// Device Advertising Type
	getAdvertisingType = "AT+ADTY?"
	setAdvertisingType = "AT+ADTY%d"

	// Device Battery Information
	getBatteryInfo = "AT+BATT?"

	// Device Baud Rate
	getBaudRate = "AT+BAUD?"
	setBaudRate = "AT+BAUD%d"

	// Device Characteristic
	getChar = "AT+CHAR?"
	setChar = "AT+CHAR0x%s"

	// Clear last connected device Address
	clear = "AT+CLEAR"

	// iBeacon Mode switch
	getBeacon = "AT+IBEA?"
	setBeacon = "AT+IBEA%d"

	// iBeacon UUID
	getUUID = "AT+IBE%d?"
	setUUID = "AT+IBE%d%s"

	// iBeacon Major
	getMajor = "AT+MARJ?"
	setMajor = "AT+MARJ0x%s"

	// iBeacon Minor
	getMinor = "AT+MINO?"
	setMinor = "AT+MINO0x%s"

	// iBeacon Measured power: Missing Now
	// AT+MEAS?
	// AT+MEAS%s

	// Device Mode
	getMode = "AT+MODE?"
	setMode = "AT+MODE%d"

	// Device Name
	getDeviceName = "AT+NAME?"
	setDeviceName = "AT+NAME%s"

	// Parity bit: Missing Now
	// AT+PARI?
	// AT+PARI%d --> need const

	// PIO1 output status (System Led)
	getStatus = "AT+PIO1?"
	setStatus = "AT+PIO1%d"

	// PIO pins outputs
	getPIO = "AT+PIO%s?"
	setPIO = "AT+PIO%s%d"

	// Device Pin code (000000-999999)
	getPin = "AT+PASS?"
	setPin = "AT+PASS%d"

	// Restore factory parameters
	factoryReset = "AT+RENEW"

	// Restart device
	reset = "AT+RESET"

	// Device Role
	setRole = "AT+ROLE%d"
	getRole = "AT+ROLE?"

	// Device RSSI
	rssi = "AT+RSSI?"

	// Last connected device address
	lastDeviceAddress = "AT+RADD?"

	// Device Bond Mode
	getBondMode = "AT+TYPE?"
	setBondMode = "AT+TYPE%d"

	// Device UUID
	setDeviceUUID = "AT+UUID?"
	getDeviceUUID = "AT+UUID%s"

	// Firmware version
	softwareVersion = "AT+VERS?"
)

// BleAdvertisingInt type
type BleAdvertisingInt string

// BleAdvertisingType type
type BleAdvertisingType int

// BleBaudRate type
type BleBaudRate int

// BleBeaconMode type
type BleBeaconMode int

// BleUUID type
type BleUUID int

// BleMode type
type BleMode int

// BlePIO type
type BlePIO int

// BlePIOPin type
type BlePIOPin string

// BleRole type
type BleRole int

// BleBondMode type
type BleBondMode int

// Ble advertising interval
const (
	Ms100  BleAdvertisingInt = "0"
	Ms152  BleAdvertisingInt = "1"
	Ms211  BleAdvertisingInt = "2"
	Ms318  BleAdvertisingInt = "3"
	Ms417  BleAdvertisingInt = "4"
	Ms546  BleAdvertisingInt = "5"
	Ms760  BleAdvertisingInt = "6"
	Ms852  BleAdvertisingInt = "7"
	Ms1022 BleAdvertisingInt = "8"
	Ms1285 BleAdvertisingInt = "9"
	Ms2000 BleAdvertisingInt = "A"
	Ms3000 BleAdvertisingInt = "B"
	Ms4000 BleAdvertisingInt = "C"
	Ms5000 BleAdvertisingInt = "D"
	Ms6000 BleAdvertisingInt = "E"
	Ms7000 BleAdvertisingInt = "F"
)

// Ble advertising type
const (
	AllinOne   BleAdvertisingType = 0 // Advertising, ScanResponse, Connectable
	LastDevice BleAdvertisingType = 1 // Only allow last device connect
	AdAndScan  BleAdvertisingType = 2 // Only allow Advertising and ScanResponse
	OnlyAd     BleAdvertisingType = 3 // Only allow advertising
)

// Ble baud rate
const (
	B9600   BleBaudRate = 0
	B19200  BleBaudRate = 1
	B38400  BleBaudRate = 2
	B57600  BleBaudRate = 3
	B115200 BleBaudRate = 4
	B4800   BleBaudRate = 5
	B2400   BleBaudRate = 6
	B1200   BleBaudRate = 7
	B230400 BleBaudRate = 8
)

//Ble beacon mode
const (
	IBeaconDisable BleBeaconMode = 0
	IBeaconEnable  BleBeaconMode = 1
)

// Ble UUIS
const (
	Part1 BleUUID = 0
	Part2 BleUUID = 1
	Part3 BleUUID = 2
	Part4 BleUUID = 3
)

//Ble mode
const (
	Transmission BleMode = 0 // Transmission mode
	PIO          BleMode = 1 // PIO collection mode + mode 0
	Remote       BleMode = 2 // Remote control mode + mode 0
)

//Ble pin output
const (
	Low  BlePIO = 0
	High BlePIO = 1
)

//Ble pio pins
const (
	Pin2 BlePIOPin = "2"
	Pin3 BlePIOPin = "3"
	Pin4 BlePIOPin = "4"
	Pin5 BlePIOPin = "5"
	Pin6 BlePIOPin = "6"
	Pin7 BlePIOPin = "7"
	Pin8 BlePIOPin = "8"
	Pin9 BlePIOPin = "9"
	PinA BlePIOPin = "A"
	PinB BlePIOPin = "B"
)

//Ble role
const (
	RoleSlaver BleRole = 0
	RoleMaster BleRole = 1
)

//Ble bond mode
const (
	NotNeedPinCode BleBondMode = 0
	AuthNotNeedPin BleBondMode = 1
	AuthWithPin    BleBondMode = 2
	AuthAndBond    BleBondMode = 3
)

// Ble struct
type Ble struct {
	fd *serial.Port
}

// BleResponse struct
type BleResponse struct {
	Error  error
	Result string
	Param  interface{}
}
