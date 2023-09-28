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
// This module is apart of the Caster-IoT-Manipulation-Framework and is dedicated to URL parsing for SSDP module
// discovery.
package Caster_HelpingHands

import (
	"net/url"
	"strings"
)

func Caster_ExtractUUID_From_URL(U string, ssdp bool) string {
	l, x := url.Parse(U)
	if x != nil {
		// panic for now
		panic(x)
	}
	if ssdp {
		port := l.Port()
		host := l.Host
		if port == "60000" {
			split := strings.Split(l.Path, "/")
			if len(split) >= 3 {
				return host + "@" + split[3]
			} else {
				return "Host was empty"
			}
		}
	}
	return "port was not SSDP_MOD_UUID"
}
