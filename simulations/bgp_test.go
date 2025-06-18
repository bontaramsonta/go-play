package bgp

import (
	"fmt"
	"testing"
)

func TestBGP(t *testing.T) {
	r1 := NewRouter("R1", 100)
	r2 := NewRouter("R2", 200)
	r3 := NewRouter("R3", 300)

	r1.AddPeer(r2)
	r2.AddPeer(r1)
	r2.AddPeer(r3)
	r3.AddPeer(r2)

	// R3 originates two prefixes
	r3.AdvertiseRoute("10.0.0.0/24")
	r3.AdvertiseRoute("10.0.1.0/25")

	// Show routing tables
	r1.ShowRoutes()

	// Simulate packet routing
	fmt.Println(r1.RoutePacket("10.0.0.1"))    // Matches 10.0.0.0/24
	fmt.Println(r1.RoutePacket("10.0.1.64"))   // Matches 10.0.1.0/25
	fmt.Println(r1.RoutePacket("192.168.1.1")) // No route
}
