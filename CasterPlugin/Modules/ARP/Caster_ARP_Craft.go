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
// Before we can send the packet data out, we need to make sure that the program can properly craft a packet
package Caster_ARP

import (
	"net"

	CasterHelpingHands "main/Modules/HelpingHands"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type ArpInfoSchem struct {
	Caster_Handle *pcap.Handle
	Interface     *net.Interface
}

func (ArpInf *ArpInfoSchem) Caster_Craft_ARP(Addr *net.IPNet) (x error) {
	Caster_Ethernet_Craft := layers.Ethernet{
		SrcMAC: ArpInf.Interface.HardwareAddr,
		DstMAC: net.HardwareAddr{
			0xff, // 1.
			0xff, // 1.
			0xff, // 1.
			0xff, // 1.
			0xff, // 1.
			0xff, // 1. - ff.ff.ff.ff.ff.ff <- Destination
		},
		EthernetType: layers.EthernetTypeARP,
	}
	Caster_ARP_Craft := layers.ARP{
		DstHwAddress: []byte{
			0, 0, 0,
			0, 0, 0,
		},
		SourceProtAddress: []byte(
			Addr.IP,
		),
		SourceHwAddress: []byte(
			ArpInf.Interface.HardwareAddr,
		),
		Operation:       layers.ARPRequest,
		ProtAddressSize: 4,
		HwAddressSize:   6,
		Protocol:        layers.EthernetTypeIPv4,
		AddrType:        layers.LinkTypeEthernet,
	}

	// Structure buffers
	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	// put to use under loop
	for _, address := range CasterHelpingHands.Caster_NetworkParse(Addr) {
		Caster_ARP_Craft.DstProtAddress = []byte(address)
		gopacket.SerializeLayers(buffer, options, &Caster_Ethernet_Craft, &Caster_ARP_Craft)
		if x = ArpInf.Caster_Handle.WritePacketData(buffer.Bytes()); x != nil {
			return x
		}
	}
	return x
}
