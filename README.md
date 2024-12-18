# ARP-Scan

ARP-Scan is a tool designed to scan IP usage using the arping command on Linux systems. This project provides a simple and efficient way to perform ARP (Address Resolution Protocol) scans within specified IP ranges.

## Prerequisites

- Linux operating system
- arping command-line tool
- Go programming language (for compilation)

## Installation

1. Install the arping tool:

```bash
apt-get install arping
```

2. Clone the repository:

```bash
git clone https://github.com/gary1030/ARP-Scan.git
cd arp-scan
```

3. Install dependencies:

```bash
go mod tidy
```

4. Compile the project:

```bash
go build -o arp-scan ./cmd/main.go
```

## Configuration

Create a YAML configuration file with the following structure:

```yaml
arp:
  nic: eth0
  ipRange:
    - 172.25.75.50-172.25.75.55
    - 172.25.75.60-172.25.75.65
  timeout: 3

reportPath: ./index.html
```

- `nic`: Specify the network interface to use for scanning
- `ipRange`: List of IP ranges to scan
- `timeout`: Timeout in seconds for ARP requests
- `reportPath`: Path to save the HTML report

## Usage

Run the compiled binary with the path to your configuration file:

```bash
./arp-scan /path/to/config.yaml
```

The scan results will be saved in an HTML report at the specified `reportPath`.
