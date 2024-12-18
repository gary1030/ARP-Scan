package main

import (
	"fmt"
	"os"

	"github.com/gary1030/ARP-Scan/config"
	"github.com/gary1030/ARP-Scan/internal/arp"
	report "github.com/gary1030/ARP-Scan/internal/html-report"
)

func main() {
	// Get config path from command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./arp-scan <config-path>")
		os.Exit(1)
	}
	configPath := os.Args[1]

	c, err := config.ReadConfig(configPath)
	if err != nil {
		panic(err)
	}

	arpScan, err := arp.NewARPScan(c.NIC)
	if err != nil {
		panic(err)
	}

	rows := []report.Row{}
	for _, ip := range c.IPs {
		fmt.Println(ip)
		_, err := arpScan.Scan(ip)
		if err != nil {
			fmt.Println("Failed to scan IP:", ip)
			rows = append(rows, report.Row{
				IP:     ip.String(),
				Status: "Unused",
			})
		} else {
			rows = append(rows, report.Row{
				IP:     ip.String(),
				Status: "In use",
			})
		}
	}

	// Generate HTML report
	err = report.GenerateHTMLReport(c.ReportPath, rows)
	if err != nil {
		panic(err)
	}

	fmt.Println("ARP scan completed successfully!")
}
