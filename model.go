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
	setChar = "AT+CHAR0x%s" // 1-65534

	// Clear last connected device Address
	clear = "AT+CLEAR"

	// Set iBeacon deploy mode
	setDeployMode = "AT+DELO%d"

	// iBeacon Mode switch
	getBeacon = "AT+IBEA?"
	setBeacon = "AT+IBEA%d"

	// iBeacon UUID
	getUUID = "AT+IBE%d?"
	setUUID = "AT+IBE%d%s" //1-4294967294

	// iBeacon Major
	getMajor = "AT+MARJ?"
	setMajor = "AT+MARJ0x%s" // 1-65534

	// iBeacon Minor
	getMinor = "AT+MINO?"
	setMinor = "AT+MINO0x%s" // 1-65534

	// iBeacon Measured power
	getIPower = "AT+MEA??"
	setIPower = "AT+MEA0x%s" //1-255 /!\ DANGER /!\

	// Device Work Mode
	getMode = "AT+MODE?"
	setMode = "AT+MODE%d"

	// Device Name
	getDeviceName = "AT+NAME?"
	setDeviceName = "AT+NAME%s"

	// Parity bit
	getParity = "AT+PARI?"
	setParity = "AT+PARI%d"

	// PIO1 output status (System Led)
	getStatus = "AT+PIO1?"
	setStatus = "AT+PIO1%d"

	// PIO pins outputs
	getPIO = "AT+PIO%s?"
	setPIO = "AT+PIO%s%d"

	// Device Pin code (000000-999999)
	getPin = "AT+PASS?"
	setPin = "AT+PASS%d"

	// Device power
	getPower = "AT+POWE?"
	setPower = "AT+POWE%d"

	// Restore factory parameters
	factoryReset = "AT+RENEW"

	// Restart device
	restart = "AT+RESET"

	// Device Role
	getRole = "AT+ROLE?"
	setRole = "AT+ROLE%d"

	// Last connected device address
	lastDeviceAddress = "AT+RADD?"

	// Device Bond Mode
	getBondMode = "AT+TYPE?"
	setBondMode = "AT+TYPE%d"

	// Device UUID
	getDeviceUUID = "AT+UUID?"
	setDeviceUUID = "AT+UUID%s"

	// Firmware version
	getSoftwareVersion = "AT+VERS?"
)

// BleAdvertisingInt type
type BleAdvertisingInt string

// BleAdvertisingType type
type BleAdvertisingType int

// BleBaudRate type
type BleBaudRate int

// BleDeployMode type
type BleDeployMode int

// BleBeaconMode type
type BleBeaconMode int

// BleUUID type
type BleUUID int

// BleMode type
type BleMode int

// BleParity type
type BleParity int

// BlePIO type
type BlePIO int

// BlePIOPin type
type BlePIOPin string

// BlePower type
type BlePower int

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

// Ble iBeacon Deploy Mode
const (
	BroadAndScan BleDeployMode = 1
	BroadOnly    BleDeployMode = 2
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

// Ble parity mode
const (
	None BleParity = 0
	Even BleParity = 1
	Odd  BleParity = 2
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

// Ble device power
const (
	P0 BlePower = 0 // -23 dbm
	P1 BlePower = 1 // -6 dbm
	P2 BlePower = 2 // 0 dbm (default)
	P3 BlePower = 3 // 6 dbm
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
