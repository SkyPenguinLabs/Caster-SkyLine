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
// This module will help caster properly read and channel the responses from the ARP messages we sent out over
// the network
package External_Control_Protocol_Roku_Caster

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ControlMap = map[string]func(a NewECPSession){
	"up": func(a NewECPSession) {
		a.Remote_Up()
	},
	// controller move down
	"down": func(a NewECPSession) {
		a.Remote_Down()
	},
	// controller move left
	"left": func(a NewECPSession) {
		a.Remote_Left()
	},
	// controller move right
	"right": func(a NewECPSession) {
		a.Remote_Right()
	},
	// both select and ok do the same thing
	"select": func(a NewECPSession) {
		a.Remote_SelectOK()
	},
	"ok": func(a NewECPSession) {
		a.Remote_SelectOK()
	},
	// volume up
	"volup": func(a NewECPSession) {
		a.Remote_VolumeUp()
	},
	// volume down
	"voldown": func(a NewECPSession) {
		a.Remote_VolumeDown()
	},
	// mute
	"mute": func(a NewECPSession) {
		a.Remote_VolumeMute()
	},
	// pause
	"options": func(a NewECPSession) {
		a.Remote_Options()
	},
	// play
	"home": func(a NewECPSession) {
		a.Remote_Home()
	},
	"back": func(a NewECPSession) {
		a.Remote_Back()
	},
}

func CallControlInit(target string) {
	Remote := *Remote_Init(target)
	input := bufio.NewReader(os.Stdin)
	var Command string
	fmt.Print("\nRoku External Controller> ")
	for {
		Command, _ = input.ReadString('\n')
		Command = strings.Replace(Command, "\n", "", -1)
		if Command != "" {
			// Execute commands
			switch Command {
			case "exit":
				// do not quit the plugin
				// plugin returning | 0 will cause a segment fault
				break
			case "help":
				// print additional menu
				fmt.Println("Available commands:")
				fmt.Println("exit - Exit the program")
				fmt.Println("help - Show this help")
			default:
				if f, ok := ControlMap[strings.ToLower(Command)]; ok {
					f(Remote)
					fmt.Print("\nRoku External Controller> ")
				} else {
					fmt.Println("Unknown command")
				}
			}
		}
	}
}
