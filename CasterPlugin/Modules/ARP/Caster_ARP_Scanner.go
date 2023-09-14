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
// Module specifies the functions required for starting a threaded scan on the network devices that were specified
//
//
//
package Caster_ARP

import (
	"fmt"
	CasterHelpingHands "main/Modules/HelpingHands"
	"net"
	"time"

	"github.com/google/gopacket/pcap"
)

func Caster_Scanner(inter *net.Interface) (x error) {
	var (
		ArpInf  ArpInfoSchem
		Handler *pcap.Handle
		Network *net.IPNet
	)
	if addresses, e := inter.Addrs(); e != nil {
		return e
	} else {
		for _, netAddress := range addresses {
			if internetP, ok := netAddress.(*net.IPNet); ok {
				if ipv4 := internetP.IP.To4(); ipv4 != nil {
					Network = &net.IPNet{
						IP:   ipv4,
						Mask: internetP.Mask[len(internetP.Mask)-4:],
					}
					break
				}
			}

		}
	}
	if !CasterHelpingHands.Caster_VerifyStabilityOfAddress(Network) {
		return
	}
	//fmt.Printf("\n\033[38;5;50m[\033[38;5;15mInformation\033[38;5;50m] Using network range (\033[38;5;15m%v\033[38;5;50m) for interface (\033[38;5;15m%v\033[38;5;50m) ", Network, inter.Name)
	Handler, x = pcap.OpenLive(inter.Name, 65536, true, pcap.BlockForever)
	if x != nil {
		return x
	}
	defer Handler.Close()
	ender := make(chan struct{})
	ArpInf.Caster_Handle = Handler
	ArpInf.Interface = inter
	go ArpInf.Caster_ReadIncoming(ender)
	defer close(ender)
	for {
		if x = ArpInf.Caster_Craft_ARP(Network); x != nil {
			fmt.Println("Error when writing packets -> ", ArpInf.Interface.Name, x)
			return x
		}
		// making 5 for right now!
		time.Sleep(time.Duration(5) * time.Second)
	}
}
