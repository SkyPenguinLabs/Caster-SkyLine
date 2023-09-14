//
// Framework Plugin By: Totally_Not_A_Haxxer
//
//
//
//┌──────────────────────────────────────────────────────────────────────────────────────────┐
//│                    																					     │
//│  _____         _                     _____     _____    _____         _         _     _   _              │
//│ |     |___ ___| |_ ___ ___    ___   |     |___|_   _|  |     |___ ___|_|___ _ _| |___| |_|_|___ ___      │
//│ |   --| .'|_ -|  _| -_|  _|  |___|  |-   -| . | | |    | | | | .'|   | | . | | | | .'|  _| | . |   |     │
//│ |_____|__,|___|_| |___|_|           |_____|___| |_|    |_|_|_|__,|_|_|_|  _|___|_|__,|_| |_|___|_|_|     │
//│                                                                      |_|								 │
//└──────────────────────────────────────────────────────────────────────────────────────────┘
//
// Working with the network and calculating the required data
//
//
//
package Caster_HelpingHands

import (
	"encoding/binary"
	"net"
)

func Caster_NetworkParse(network *net.IPNet) (DataHolder []net.IP) {
	Caster_NetNum := binary.BigEndian.Uint32([]byte(
		network.IP,
	))
	Caster_NetMask := binary.BigEndian.Uint32([]byte(
		network.Mask,
	))
	Caster_Netwhole := Caster_NetNum & Caster_NetMask
	Caster_Netbroad := Caster_Netwhole | ^Caster_NetMask
	for Caster_Netwhole++; Caster_Netwhole < Caster_Netbroad; Caster_Netwhole++ {
		var buf [4]byte // - 1.1.1.1  | Holding 4 octets
		binary.BigEndian.PutUint32(buf[:], Caster_Netwhole)
		DataHolder = append(DataHolder, net.IP(buf[:]))
	}
	return DataHolder
}
