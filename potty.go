package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"code.google.com/p/gopacket"
	"code.google.com/p/gopacket/pcap"
)

func main() {
	var device = flag.String("device", "eth0", "Device to capture on")
	flag.Parse()

	handle, err := pcap.OpenLive(*device,
		60, // We really only need enough to glean IP information
		true,
		0)

	if err != nil {
		panic(err)
	}

	// Filter out IPs of ServiceNet
	// According to http://www.rackspace.com/knowledge_center/article/updating-servicenet-routes-on-cloud-servers-created-before-june-3-2013
	err = handle.SetBPFFilter("not net 10.176.0.0/12 and not net 10.208.0.0/12")

	if err != nil {
		panic(err)
	}

	bytesum := 0
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-c
		os.Stdout.Sync()
		log.Printf("Bytes: %v\n", bytesum)
		os.Exit(1)
	}()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		metadata := packet.Metadata()
		log.Println(metadata.Length)
		bytesum += metadata.Length
	}
}
