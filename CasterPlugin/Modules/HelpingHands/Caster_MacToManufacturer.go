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
// The manufac function allows Caster to get the OUI of a given MAC address quite quickly!
//
//
//
//
package Caster_HelpingHands

import (
	"bufio"
	"os"
	"strings"
)

var MACDB = make(map[string]string)

func init() {
	file, x := os.Open("Backend/Core/Data/DB.txt")
	if x != nil {
		panic(x)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) >= 2 {
			oui := strings.ToLower(parts[0])
			vendor := parts[1]
			MACDB[oui] = vendor
		}
	}
	if x = scanner.Err(); x != nil {
		panic(x)
	}
}

func Caster_MacToManufac(address string) (OUI string) {
	if len(address) >= 9 {
		if vendor, ok := MACDB[string(address[:8])]; ok {
			return vendor
		} else {
			return "UnknownInDB"
		}
	} else {
		return "InvalidMAC"
	}
}
