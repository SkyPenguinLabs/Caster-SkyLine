package Caster_Exportable

import "time"

// Scan results
var ScanResults Scan

type Scan struct {
	Device      Device
	Total_Hosts int // Amount of hosts
	Total_MACS  int // Amount scanned and responded
	Total_OUIS  int // Amount properly traced
}

type Device []struct {
	Name       string    // Device name      | OUI
	IPA        string    // Device addr      | IP Address = IPA
	MAC        string    // Devide MAc       | MAC
	Timestamp  time.Time // Device Timestamp | Last seen at?
	State      string    // State of device  | Dead or Alive
	SkyLine_DB string    // Database for APIs
}

var ScanDevs Device
