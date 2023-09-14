package Caster_Exportable

import (
	"encoding/json"
	"io/ioutil"
)

// Write and marshal structure to file

func WriteStruct(Filename string) bool {
	Data, x := json.Marshal(ScanResults)
	if x != nil {
		panic(x)
	}
	x = ioutil.WriteFile(Filename, Data, 0644)
	if x != nil {
		panic(x)
	}
	return true
}
