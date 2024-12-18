package arp

import (
	"fmt"
	"net"
	"os/exec"
)

type ARPScan struct {
	Interface *net.Interface
	IPAddress net.IP
	Timeout   int
}

type IPRecord struct {
	IP  string
	MAC string
}

func NewARPScan(nic string) (*ARPScan, error) {
	// Check arping command exists
	_, err := exec.LookPath("arping")
	if err != nil {
		return nil, fmt.Errorf("failed to find arping command: %w", err)
	}

	ifi, err := net.InterfaceByName(nic)
	if err != nil {
		return nil, fmt.Errorf("failed to get interface %s: %w", nic, err)
	}

	// Get self IP address
	addrs, err := ifi.Addrs()
	if err != nil {
		return nil, fmt.Errorf("failed to get interface addresses: %w", err)
	}

	return &ARPScan{
		Interface: ifi,
		IPAddress: addrs[0].(*net.IPNet).IP,
		Timeout:   2,
	}, nil
}

func (a *ARPScan) Scan(ip net.IP) (*IPRecord, error) {
	// Execute arping command
	cmd := exec.Command("arping", "-I", a.Interface.Name, "-c", "1", "-w", fmt.Sprintf("%d", a.Timeout), ip.String())

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute arping: %w", err)
	}

	// Parse arping output
	hwaddr, err := net.ParseMAC(string(out))
	if err != nil {
		return &IPRecord{
			IP:  ip.String(),
			MAC: "",
		}, nil
	}

	return &IPRecord{
		IP:  ip.String(),
		MAC: hwaddr.String(),
	}, nil
}
