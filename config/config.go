package config

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	IPs        []net.IP
	Timeout    time.Duration
	NIC        string
	ReportPath string
}

type arpConfig struct {
	IPRange []string `yaml:"ipRange"`
	Timeout int      `yaml:"timeout"`
	NIC     string   `yaml:"nic"`
}

type rawConfig struct {
	ARP        arpConfig `yaml:"arp"`
	ReportPath string    `yaml:"reportPath"`
}

func ReadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rawConfig rawConfig
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&rawConfig)
	if err != nil {
		return nil, err
	}

	fmt.Println("Parsed config successfully!")

	var ips []net.IP
	for _, ipRange := range rawConfig.ARP.IPRange {
		rangeIPs, err := parseIPRange(ipRange)
		if err != nil {
			return nil, err
		}
		ips = append(ips, rangeIPs...)
	}

	return &Config{
		IPs:        ips,
		Timeout:    time.Duration(rawConfig.ARP.Timeout) * time.Second,
		ReportPath: rawConfig.ReportPath,
		NIC:        rawConfig.ARP.NIC,
	}, nil
}

func parseIPRange(ipRange string) ([]net.IP, error) {
	parts := strings.Split(ipRange, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid IP range format: %s", ipRange)
	}

	startIP := net.ParseIP(strings.TrimSpace(parts[0]))
	endIP := net.ParseIP(strings.TrimSpace(parts[1]))

	// check endIP is greater than startIP
	for i := 0; i < len(startIP); i++ {
		if startIP[i] > endIP[i] {
			return nil, fmt.Errorf("invalid IP range: %s", ipRange)
		}
	}

	if startIP == nil || endIP == nil {
		return nil, fmt.Errorf("invalid IP address in range: %s", ipRange)
	}

	var ips []net.IP
	for ip := cloneIP(startIP); !ip.Equal(endIP); ip = nextIP(ip) {
		ips = append(ips, cloneIP(ip))
	}
	ips = append(ips, cloneIP(endIP))

	return ips, nil
}

func nextIP(ip net.IP) net.IP {
	nextIP := cloneIP(ip)
	for j := len(nextIP) - 1; j >= 0; j-- {
		nextIP[j]++
		if nextIP[j] > 0 {
			break
		}
	}
	return nextIP
}

func cloneIP(ip net.IP) net.IP {
	clone := make(net.IP, len(ip))
	copy(clone, ip)
	return clone
}
