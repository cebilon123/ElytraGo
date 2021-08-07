package handler

import (
	"github.com/cebilon123/ElytraGo/mbinary"
	"github.com/cebilon123/ElytraGo/pckt"
	"net"
)


func GetFuncBasedOnPct(pct pckt.Packet) PacketHandlerFunc {
	if pct.Id == 0 {
		return handleHandshake
	}

	return nil
}

func handleHandshake(pct pckt.Packet, c net.Conn) {
	prtclVer, readI := mbinary.VarInt(pct.Data)
	println(prtclVer, readI)
}