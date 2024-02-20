package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	// Setting the -i4 and -i6 flags
	ipV4Only := flag.Bool("i4", false, "Display only IPv4 addresses")
	ipV6Only := flag.Bool("i6", false, "Display only IPv6 addresses")
	flag.Parse()

	// Checks if the file name is provided
	if flag.NArg() < 1 {
		fmt.Println("Usage: go run resolve_domains.go [-i4 | -i6] <filename>")
		os.Exit(1)
	}

	filename := flag.Arg(0)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Map to store unique IPs
	uniqueIPs := make(map[string]bool)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		ips, err := net.LookupIP(domain)
		if err != nil {
//			fmt.Fprintf(os.Stderr, "Could not get IPs for %s: %v\n", domain, err)
			continue
		}
		for _, ip := range ips {
			if *ipV4Only && ip.To4() != nil {
				// Add IPv4 to the map
				uniqueIPs[ip.String()] = true
			} else if *ipV6Only && ip.To4() == nil {
				// Add IPv6 to the map
				uniqueIPs[ip.String()] = true
			} else if !*ipV4Only && !*ipV6Only {
				// If none of the -i4 or -i6 flags are used
				fmt.Printf("%s IN A %s\n", domain, ip.String())
			}
		}
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
	}

	// Print unique IPs if -i4 or -i6 flags are used
	if *ipV4Only || *ipV6Only {
		for ip := range uniqueIPs {
			fmt.Println(ip)
		}
	}
}
