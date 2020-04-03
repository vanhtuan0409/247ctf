package main

import (
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	f, err := os.Create("res.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	handle, err := pcap.OpenOffline("./error_reporting.pcap")
	if err != nil {
		panic(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	source.NextPacket() // skip 1st packet (request packet)
	for packet := range source.Packets() {
		msg, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), packet.Data())
		if err != nil {
			continue
		}

		body, ok := msg.Body.(*icmp.Echo)
		if !ok {
			panic("parse failed")
		}
		f.Write(body.Data[34:]) // tricky, dont know why Go parse extra header into body
	}
}
