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
// Since this is a plugin that will be compiled and loaded into SkyLine, the FFI framework within SkyLine allows
// the developers of the plugin and primary program to declare variables that are exportable. This makes it much
// much more easier on the development end of things
//
//
package Caster_Exportable

import (
	SkyEnvironment "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

var (
	Caster_IndividalInterface string
	Caster_CapturedHosts      int
	Caster_TotalOUIs          int
	Caster_MACS               []string // Macs
	Caster_IPS                []string
)

func Caster_ExportMacs() SkyEnvironment.SL_Object {
	arr := make([]SkyEnvironment.SL_Object, 0)
	for i := 0; i < len(Caster_MACS); i++ {
		arr = append(arr, &SkyEnvironment.SL_String{Value: Caster_MACS[i]})
	}
	return &SkyEnvironment.SL_Array{Elements: arr}
}

func Caster_ExportIPS() SkyEnvironment.SL_Object {
	arr := make([]SkyEnvironment.SL_Object, 0)
	for i := 0; i < len(Caster_IPS); i++ {
		arr = append(arr, &SkyEnvironment.SL_String{Value: Caster_IPS[i]})
	}
	return &SkyEnvironment.SL_Array{Elements: arr}
}

func Caster_ExportCountHosts() SkyEnvironment.SL_Object {
	return &SkyEnvironment.SL_Integer{Value: Caster_CapturedHosts}
}

func Caster_ExportCountOUIS() SkyEnvironment.SL_Object {
	return &SkyEnvironment.SL_Integer{Value: Caster_TotalOUIs}
}
