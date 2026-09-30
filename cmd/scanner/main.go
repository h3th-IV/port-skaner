package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Some Shit")

	target := flag.String("target", "", "Target host or IP address (required)")
	ports := flag.String("port", "1-1024", "Ports to scan (e.g. 80,443 or 1-1000)")
	timeout := flag.Duration("timeout", 2*time.Second, "Connection timeout")
	workers := flag.Int("workers", 100, "Number of concurrent workers")
	output := flag.String("output", "table", "Output format: table | json")
	banner := flag.Bool("banner", true, "Attempt to grab banners on open ports")
	verbose := flag.Bool("v", false, "Verbose output")

	flag.Parse()

	// Validation
	if *target == "" {
		fmt.Println("Erorr: -target is required")
		flag.Usage()
		os.Exit(1)
	}

	// temp confirmation
	fmt.Printf("Starting scan...\n")
	fmt.Printf("  Target   : %s\n", *target)
	fmt.Printf("  Ports    : %s\n", *ports)
	fmt.Printf("  Timeout  : %v\n", *timeout)
	fmt.Printf("  Workers  : %d\n", *workers)
	fmt.Printf("  Output   : %s\n", *output)
	fmt.Printf("  Banner   : %v\n", *banner)
	fmt.Printf("  Verbose  : %v\n", *verbose)
}
