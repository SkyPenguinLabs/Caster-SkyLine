// Framework Plugin By: Totally_Not_A_Haxxer
//
// ┌──────────────────────────────────────────────────────────────────────────────────────────┐
// │                    																					     │
// │  _____         _                     _____     _____    _____         _         _     _   _              │
// │ |     |___ ___| |_ ___ ___    ___   |     |___|_   _|  |     |___ ___|_|___ _ _| |___| |_|_|___ ___      │
// │ |   --| .'|_ -|  _| -_|  _|  |___|  |-   -| . | | |    | | | | .'|   | | . | | | | .'|  _| | . |   |     │
// │ |_____|__,|___|_| |___|_|           |_____|___| |_|    |_|_|_|__,|_|_|_|  _|___|_|__,|_| |_|___|_|_|     │
// │                                                                      |_|								 │
// └──────────────────────────────────────────────────────────────────────────────────────────┘
//
// This module will help caster properly read and channel the responses from the ARP messages we sent out over
// the network
package Caster_ARP

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"

	Caster_Exportable "main/Modules/Exportable"
	Caster_HelpingHands "main/Modules/HelpingHands"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var IpTrack []string
var MacTrack []string

func ArrContains(arr []string, obj string) bool {
	for _, v := range arr {
		if v == obj {
			return true
		}
	}
	return false
}

func (ArpInf *ArpInfoSchem) Caster_ReadIncoming(channel chan struct{}) {
	src := gopacket.NewPacketSource(ArpInf.Caster_Handle, layers.LayerTypeEthernet)
	pkttoread := src.Packets()
	for {
		var PKT gopacket.Packet
		select {
		// read channel
		case <-channel:
			println("Debug: Closing reader")
			return
		case PKT = <-pkttoread:
			layer := PKT.Layer(layers.LayerTypeARP)
			if layer != nil {
				layerinfo := layer.(*layers.ARP)
				if layerinfo.Operation != layers.ARPReply || bytes.Equal([]byte(ArpInf.Interface.HardwareAddr), layerinfo.SourceHwAddress) {
					continue
				}
				mac := fmt.Sprint(net.HardwareAddr(layerinfo.SourceHwAddress))
				ip := fmt.Sprint(net.IP((layerinfo.SourceProtAddress)))
				if ArrContains(IpTrack, ip) && ArrContains(MacTrack, mac) {
					continue
				} else {
					if mac != "" && fmt.Sprint(net.IP((layerinfo.SourceProtAddress))) != "" {
						DeviceName := Caster_HelpingHands.Caster_MacToManufac(mac)
						var database string
						if strings.Contains(DeviceName, "Roku") {
							database = "roku"
						} else if strings.Contains(DeviceName, "Apple") {
							database = "apple"
						} else if strings.Contains(DeviceName, "Google") {
							database = "google"
						}
						Caster_Exportable.ScanResults.Device = append(Caster_Exportable.ScanResults.Device, struct {
							Name       string    // Device name      | OUI
							IPA        string    // Device addr      | IP Address = IPA
							MAC        string    // Devide MAc       | MAC
							Timestamp  time.Time // Device Timestamp | Last seen at?
							State      string    // State of device  | Dead or Alive
							SkyLine_DB string    // Database for API | Skyline required
						}{
							Name:       DeviceName,
							IPA:        fmt.Sprint(net.IP((layerinfo.SourceProtAddress))),
							MAC:        mac,
							SkyLine_DB: database,
							Timestamp:  PKT.Metadata().Timestamp,
						})
						IpTrack = append(IpTrack, ip)
						MacTrack = append(MacTrack, mac)
						Caster_Exportable.ScanResults.Total_Hosts++
						Caster_Exportable.ScanResults.Total_MACS++
						time.Sleep(100 * time.Millisecond)
					}
				}
			}
		}
	}
}
