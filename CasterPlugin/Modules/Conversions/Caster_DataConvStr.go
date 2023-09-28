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
// Since this module is all about data conversions, its important that we can explain it in some shape or form.
//
// When working with the FFI framework in SkyLine, you need to be able to export and import arguments into the plugin
//
// and from the plugin. The issue with this is that despite an object in SkyLine being an alias to a say string type in Go
//
// on the backend, Go does not see it as that type and rather sees it as its own direct structure which makes perfect sense!
//
// in order to properly import and export data, we need to define conversion functions internally. The reason the SkyLine script
//
// does not need to define these is because the results being put into the plugin or imported into the plugin are being converted
//
// to golang data types whilst the exported data is being converted directly to SkyLine Objects before the data is even returned.
//
package Caster_SkyConversions

import (
	SkyEnvironment "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func Caster_Convert_String_Object(Object SkyEnvironment.SL_Object) string {
	if v, ok := Object.(*SkyEnvironment.SL_String); ok {
		return v.Value
	}
	return "err"
}
