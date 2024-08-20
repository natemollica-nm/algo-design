package main

import (
	"algo-design/internal/ospf"
	"fmt"
	"net"
)

func main() {
	routerA := &ospf.Router{
		LoopbackIPs:  []net.IP{net.ParseIP("192.168.0.1")},
		InterfaceIPs: []net.IP{net.ParseIP("10.1.1.1"), net.ParseIP("172.16.1.1")},
		Links:        make(map[string]*ospf.Link),
		Area:         "0",
		LSDB:         make(map[string]*ospf.LSA),
		OSPFPriority: 10,
	}
	routerA.ID = "1.1.1.1"

	routerB := &ospf.Router{
		LoopbackIPs:  []net.IP{net.ParseIP("192.168.0.2")},
		InterfaceIPs: []net.IP{net.ParseIP("10.1.1.2")},
		Links:        make(map[string]*ospf.Link),
		Area:         "0",
		LSDB:         make(map[string]*ospf.LSA),
		OSPFPriority: 5,
	}
	routerB.ID = ospf.DetermineRouterID(routerB)

	routerC := &ospf.Router{
		LoopbackIPs:  nil, // No loopback IP
		InterfaceIPs: []net.IP{net.ParseIP("10.1.1.3"), net.ParseIP("172.16.1.3")},
		Links:        make(map[string]*ospf.Link),
		Area:         "0",
		LSDB:         make(map[string]*ospf.LSA),
		OSPFPriority: 1,
	}
	routerC.ID = ospf.DetermineRouterID(routerC)

	routers := []*ospf.Router{routerA, routerB, routerC}
	dr, bdr := ospf.ElectDRAndBDR(routers)

	fmt.Printf("Elected DR: %s with ID %s\n", dr.ID, dr.ID)
	fmt.Printf("Elected BDR: %s with ID %s\n", bdr.ID, bdr.ID)

	// Assume bandwidth in bps (e.g., 100Mbps)
	// Define links with bandwidths and network CIDRs
	routerA.Links[routerB.ID] = &ospf.Link{To: routerB, Bandwidth: 100e6, Network: "10.1.1.0/24"} // 100 Mbps
	routerB.Links[routerA.ID] = &ospf.Link{To: routerA, Bandwidth: 100e6, Network: "10.1.1.0/24"} // 100 Mbps

	routerB.Links[routerC.ID] = &ospf.Link{To: routerC, Bandwidth: 10e6, Network: "10.1.2.0/24"} // 10 Mbps
	routerC.Links[routerB.ID] = &ospf.Link{To: routerB, Bandwidth: 10e6, Network: "10.1.2.0/24"} // 10 Mbps

	routerA.Links[routerC.ID] = &ospf.Link{To: routerC, Bandwidth: 1e6, Network: "10.1.3.0/24"} // 1 Mbps
	routerC.Links[routerA.ID] = &ospf.Link{To: routerA, Bandwidth: 1e6, Network: "10.1.3.0/24"} // 1 Mbps

	area0 := &ospf.Area{
		ID: "0",
		Routers: map[string]*ospf.Router{
			routerA.ID: routerA,
			routerB.ID: routerB,
			routerC.ID: routerC,
		},
	}

	network := &ospf.Network{
		Routers: map[string]*ospf.Router{
			routerA.ID: routerA,
			routerB.ID: routerB,
			routerC.ID: routerC,
		},
		Areas: map[string]*ospf.Area{
			"0": area0,
		},
	}

	for _, router := range network.Routers {
		router.GenerateLSA()
	}

	area0.FloodLSAs()

	dist := ospf.DijkstraWithMultiPath(network, routerA.ID, "0")
	// Output the results including costs, router IDs, and network CIDRs
	for routerID, cost := range dist {
		if routerID != routerA.ID {
			for _, link := range routerA.Links {
				if link.To.ID == routerID {
					fmt.Printf("Cost from %s to %s: %f (Network: %s)\n", routerA.ID, routerID, cost, link.Network)
				}
			}
		}
	}

}
