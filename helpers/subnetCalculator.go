package helpers

import (
	"fmt"
	"math"
	"net"
	"os"
	"strings"
)

func CalculateSubnet(cidr string) map[string]interface{} {
	// Parse CIDR
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Printf("Invalid CIDR notation: %v\n", err)
		os.Exit(2)
	}

	// Get mask size
	ones, bits := ipnet.Mask.Size()

	// Calculate total hosts and usable hosts
	totalHosts := math.Pow(2, float64(bits-ones))
	usableHosts := totalHosts - 2
	if ones >= bits-1 { // For /31 and /32 networks
		usableHosts = totalHosts
	}

	// Get network and broadcast addresses
	networkIP := ipnet.IP

	// Calculate broadcast address
	broadcastIP := make(net.IP, len(networkIP))
	copy(broadcastIP, networkIP)
	for i := 0; i < len(broadcastIP); i++ {
		broadcastIP[i] = networkIP[i] | ^ipnet.Mask[i]
	}

	// Calculate first and last usable IP
	firstUsableIP := make(net.IP, len(networkIP))
	copy(firstUsableIP, networkIP)
	if ones < bits-1 { // Not /31 or /32
		firstUsableIP[len(firstUsableIP)-1]++
	}

	lastUsableIP := make(net.IP, len(broadcastIP))
	copy(lastUsableIP, broadcastIP)
	if ones < bits-1 { // Not /31 or /32
		lastUsableIP[len(lastUsableIP)-1]--
	}

	// Get subnet mask in dotted decimal
	mask := ipnet.Mask
	subnetMask := fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])

	// Determine IP class
	ipClass := getIPClass(ip)

	// Determine IP type
	ipType := getIPType(ip)

	return map[string]interface{}{
		"cidr":          cidr,
		"networkAddr":   networkIP,
		"firstUsableIP": firstUsableIP,
		"lastUsableIP":  lastUsableIP,
		"broadcast":     broadcastIP,
		"totalHosts":    totalHosts,
		"usableHosts":   usableHosts,
		"subnetMask":    subnetMask,
		"ipClass":       ipClass,
		"ipType":        ipType,
	}
}

// getIPClass determines the IP class (A, B, C, D, or E)
func getIPClass(ip net.IP) string {
	// Convert to IPv4 if it's IPv6-mapped IPv4
	if ip.To4() != nil {
		ip = ip.To4()
	}

	// Check if IPv6
	if len(ip) == 16 && !strings.Contains(ip.String(), "::ffff:") {
		return "IPv6 (No Class)"
	}

	// IPv4 classes
	firstOctet := int(ip[0])
	switch {
	case firstOctet < 128:
		return "A"
	case firstOctet < 192:
		return "B"
	case firstOctet < 224:
		return "C"
	case firstOctet < 240:
		return "D (Multicast)"
	default:
		return "E (Reserved)"
	}
}

// getIPType determines if the IP is private, public, etc.
func getIPType(ip net.IP) string {
	if ip.IsLoopback() {
		return "Loopback"
	} else if ip.IsMulticast() {
		return "Multicast"
	} else if ip.IsLinkLocalUnicast() {
		return "Link-local Unicast"
	} else if ip.IsPrivate() {
		return "Private"
	} else if ip.IsUnspecified() {
		return "Unspecified"
	} else {
		return "Public"
	}
}

func IsIPInCIDR(ipStr, cidrStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		fmt.Println("Invalid IP address format")
		os.Exit(2)
	}
	_, ipnet, err := net.ParseCIDR(cidrStr)
	if err != nil {
		fmt.Printf("Invalid CIDR notation: %v\n", err)
		os.Exit(2)
	}
	return ipnet.Contains(ip)
}
