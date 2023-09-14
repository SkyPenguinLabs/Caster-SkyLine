package main

import (
	CasterARP "main/Modules/ARP"
	Caster_SkyLine "main/Modules/Conversions"
	Caster_Exportable "main/Modules/Exportable"
	"sync/atomic"
	"time"

	SkyEnvironment "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

var (
	CasterVer = "1.0.0 [BETA]"
	Network   []string
)

func Framework(env *SkyEnvironment.SkyLineEnvironment, args ...SkyEnvironment.SL_Object) SkyEnvironment.SL_Object {
	if len(args) > 1 {
		switch t := args[1].(type) {
		case *SkyEnvironment.SL_String:
			// string is the only one supported
			switch t.Value {
			case "StartArp":
				// parse arguments
				if len(args) >= 3 {
					// need at least three arguments
					// arg[1] = Use only specified interface
					// arg[2] = interfacename if arg[1] is true
					arp_driver_a1 := Caster_SkyLine.Caster_Convert_Boolean_Object(args[2])
					arp_driver_a2 := Caster_SkyLine.Caster_Convert_String_Object(args[3])
					if arp_driver_a1 {
						// we want to be able to trust that the developer verified this ( at least I did )
						// warn: zero trust in users
						go func() {
							CasterARP.Caster_Driver(true, arp_driver_a2)
						}()
						return &SkyEnvironment.SL_Boolean{Value: true}
					} else {
						// use multiple interfacaes and call driver
						go func() {
							CasterARP.Caster_Driver(false, "")
						}()
						return &SkyEnvironment.SL_Boolean{Value: true}
					}
				}
				return &SkyEnvironment.SL_Error{Message: "Thats weird, when calling `Framework()->StartArp` there was an argument error, make sure you specify more than 3 arguments or exactly 3 arguments"}
			case "WriteRoutine":
				if len(args) >= 3 {
					writer_a1 := Caster_SkyLine.Caster_Convert_String_Object(args[2])   // file to write
					writer_a2 := Caster_SkyLine.Caster_Convert_Integer8_Object(args[3]) // timer or amount in seconds
					timer := time.NewTicker(time.Second * time.Duration(writer_a2))
					var Atom7 atomic.Value
					Atom7.Store(func() {
						Caster_Exportable.WriteStruct(writer_a1)
					})
					go func() {
						for {
							select {
							case <-timer.C:
								fun := Atom7.Load().(func())
								fun()
							}
						}
					}()

				} else {
					return &SkyEnvironment.SL_Error{Message: "When calling `Framework()->WriteToFile` there was a filename that was missing...or rather a second argument to the command and a second count is also needed"}
				}
				return &SkyEnvironment.SL_NULL{}
			case "WriteToFile":
				if len(args) >= 2 {
					// parse the first value as an argument
					writer_a1 := Caster_SkyLine.Caster_Convert_String_Object(args[2])
					Caster_Exportable.WriteStruct(writer_a1)
				} else {
					return &SkyEnvironment.SL_Error{Message: "When calling `Framework()->WriteToFile` there was a filename that was missing...or rather a second argument to the command"}
				}
				return &SkyEnvironment.SL_NULL{}
			case "ExtractHostCount":
				return &SkyEnvironment.SL_Integer{Value: Caster_Exportable.Caster_CapturedHosts}
			case "ExtractMacCount":
				return &SkyEnvironment.SL_Integer{Value: len(Caster_Exportable.Caster_MACS)}
			case "VersionPlugin":
				return &SkyEnvironment.SL_String{Value: CasterVer}
			case "ExtractHosts":
				elems := make([]SkyEnvironment.SL_Object, 0)
				for i := 0; i < len(Caster_Exportable.Caster_IPS); i++ {
					elems = append(elems, &SkyEnvironment.SL_String{Value: Caster_Exportable.Caster_IPS[i]})
				}
				return &SkyEnvironment.SL_Array{Elements: elems}
			case "ExtractMACS":
				elems := make([]SkyEnvironment.SL_Object, 0)
				for i := 0; i < len(Caster_Exportable.Caster_MACS); i++ {
					elems = append(elems, &SkyEnvironment.SL_String{Value: Caster_Exportable.Caster_MACS[i]})
				}
				return &SkyEnvironment.SL_Array{Elements: elems}
			case "ExtractArrNet":
				NetElem := make([]SkyEnvironment.SL_Object, 0)
				for n := 0; n < len(Network); n++ {
					NetElem = append(NetElem, &SkyEnvironment.SL_String{Value: Network[n]})
				}
				return &SkyEnvironment.SL_Array{Elements: NetElem}
			default:
				return &SkyEnvironment.SL_Error{Message: "Caster Framework: Parameter / Argument is not exactly supported?"}
			}
		default:
			return &SkyEnvironment.SL_Error{Message: "Caster Framework: Sorry, this data type is not supported, NEED=String"}
		}
	} else {
		return &SkyEnvironment.SL_NULL{}
	}
}

func main() {
	println("hello world!")
}
