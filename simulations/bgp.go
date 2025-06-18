package bgp

import (
	"fmt"
	"slices"
	"strings"
)

// Route represents a network prefix and the AS path to reach it.
type Route struct {
	Prefix string
	ASPath []int
}

// Router simulates a BGP router
type Router struct {
	Name  string
	ASN   int
	Peers []*Router
	RIB   map[string]Route // Routing Information Base
}

// NewRouter creates a new BGP router
func NewRouter(name string, asn int) *Router {
	return &Router{
		Name: name,
		ASN:  asn,
		RIB:  make(map[string]Route),
	}
}

// AddPeer connects two routers as BGP peers
func (r *Router) AddPeer(peer *Router) {
	r.Peers = append(r.Peers, peer)
}

// AdvertiseRoute simulates advertising a route to peers
func (r *Router) AdvertiseRoute(prefix string) {
	route := Route{
		Prefix: prefix,
		ASPath: []int{r.ASN},
	}
	r.RIB[prefix] = route
	for _, peer := range r.Peers {
		peer.ReceiveRoute(route, r)
	}
}

// ReceiveRoute handles an incoming BGP update from a peer
func (r *Router) ReceiveRoute(route Route, from *Router) {
	// Prevent loop
	if slices.Contains(route.ASPath, r.ASN) {
		return
	}

	// Append own ASN to AS Path
	newPath := append([]int{r.ASN}, route.ASPath...)
	newRoute := Route{
		Prefix: route.Prefix,
		ASPath: newPath,
	}

	// Select route if it's new or shorter
	existing, exists := r.RIB[route.Prefix]
	if !exists || len(newPath) < len(existing.ASPath) {
		r.RIB[route.Prefix] = newRoute
		for _, peer := range r.Peers {
			if peer != from {
				peer.ReceiveRoute(newRoute, r)
			}
		}
	}
}

// ShowRoutes prints the current routing table
func (r *Router) ShowRoutes() {
	fmt.Printf("Router %s (AS%d) Routing Table:\n", r.Name, r.ASN)
	for _, route := range r.RIB {
		fmt.Printf("  %s via %s\n", route.Prefix, formatASPath(route.ASPath))
	}
	fmt.Println()
}

func formatASPath(path []int) string {
	strs := make([]string, len(path))
	for i, as := range path {
		strs[i] = fmt.Sprintf("AS%d", as)
	}
	return strings.Join(strs, " → ")
}

// ipInPrefix checks if an IP belongs to a given CIDR prefix like "10.0.0.0/24"
func ipInPrefix(ip string, prefix string) bool {
	ipParts := strings.Split(ip, ".")
	prefixParts := strings.Split(prefix, "/")
	baseIP := strings.Split(prefixParts[0], ".")
	maskLength := atoi(prefixParts[1])

	ipBin := toBinary(ipParts)
	baseBin := toBinary(baseIP)

	return ipBin[:maskLength] == baseBin[:maskLength]
}

func toBinary(parts []string) string {
	var b strings.Builder
	for _, p := range parts {
		num := atoi(p)
		b.WriteString(fmt.Sprintf("%08b", num))
	}
	return b.String()
}

func atoi(s string) int {
	n := 0
	for _, ch := range s {
		n = n*10 + int(ch-'0')
	}
	return n
}

func (r *Router) RoutePacket(ip string) string {
	bestMatch := ""
	bestMask := -1
	var selected Route

	for prefix, route := range r.RIB {
		if ipInPrefix(ip, prefix) {
			maskLen := atoi(strings.Split(prefix, "/")[1])
			if maskLen > bestMask {
				bestMask = maskLen
				bestMatch = prefix
				selected = route
			}
		}
	}

	if bestMatch == "" {
		return fmt.Sprintf("No route found for IP %s in router %s (AS%d)", ip, r.Name, r.ASN)
	}

	return fmt.Sprintf("Packet to %s will be routed via %s (AS Path: %s)", ip, bestMatch, formatASPath(selected.ASPath))
}
