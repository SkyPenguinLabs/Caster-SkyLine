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
	"fmt"

	ECP_Control "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/BrandNames/Roku/SmartDevices"
)

func (ECP_Session *NewECPSession) Remote_DevUp() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_DEVUP,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_DevDown() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_DEVDOWN,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_Left() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_LEFT,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_Right() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_RIGHT,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_SelectOK() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_OK,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_Back() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_BACK,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_Up() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_UP,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_Down() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_DOWN,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_VolumeDown() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_VDOWN,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_VolumeUp() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_VUP,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_VolumeMute() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_MUTE,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_Options() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_OPTIONS,
			ECP_Session.Target,
		),
	)
}

func (ECP_Session *NewECPSession) Remote_Home() {
	EmptyPost(
		fmt.Sprintf(
			ECP_Control.ROKU_KEYPRESS_HOME,
			ECP_Session.Target,
		),
	)
}
