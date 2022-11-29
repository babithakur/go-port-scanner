package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"os"
)

func main() {
	var port uint
	var port_range uint
	open_ports := 0
	ip := flag.String("i", "", "Target IP to scan")
	flag.UintVar(&port_range, "p", 100, "Port range")
	flag.Parse()
	if len(os.Args) < 3 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	ip_check := strings.Split(*ip, ".")
	if len(ip_check) == 4 {
		fmt.Println("Scanning started...")
		for port = 0; port < port_range; port++ {
			addr := fmt.Sprintf("%s:%d", *ip, port)
			conn, err := net.Dial("tcp", addr)
			if err == nil {
				fmt.Println("Found open port:", port)
				conn.Close()
				open_ports += 1
			}
		}
		fmt.Println("================Scan Completed===================")
		fmt.Println("Scanned", port, "ports, found", open_ports, "open!")
		fmt.Println("==================================================")
	} else {
		fmt.Println("Invalid IP!")
		os.Exit(1)
	}
}
