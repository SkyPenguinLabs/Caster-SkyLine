package Caster_ARP

import (
	"fmt"
	"net"
	"sync"
)

func Caster_Driver(single bool, inter string) {
	var wg sync.WaitGroup
	interfaces, x := net.Interfaces()
	if x != nil {
		panic(x)
	} else {
		if !single {
			for _, interfacename := range interfaces {
				wg.Add(1)
				go func(interfacename net.Interface) {
					defer wg.Done()
					if x := Caster_Scanner(&interfacename); x != nil {
						fmt.Println("[-] Error Scanning on interface -> ", interfacename.Name, x)
					}

				}(interfacename)
			}
		} else {
			for _, interfacename := range interfaces {
				if interfacename.Name == inter {
					wg.Add(1)
					go func(name net.Interface) {
						defer wg.Done()
						if x := Caster_Scanner(&name); x != nil {
							fmt.Println("[-] Error Scanning on interface -> ", interfacename.Name, x)
						}
					}(interfacename)
				}
			}
		}
	}
	wg.Wait()
}
