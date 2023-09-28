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
// This module is dedicated to data type manipulation or rather data manipulation!
//
//
//
//
package Caster_HelpingHands


func Type_Array_ExterminateExtra(src []string) []string {
	k := make(map[string]bool)
	j := []string{}
	for _, y := range src {
		if _, u := k[y]; !u {
			k[y] = true
			j = append(j, y)
		}
	}
	return j
}
