package ospf

import (
	"net"
	"sort"
)

type Router struct {
	LoopbackIPs  []net.IP
	InterfaceIPs []net.IP
	Links        map[string]*Link
	Area         string
	ID           string
	LSDB         map[string]*LSA
	OSPFPriority int
}

type Link struct {
	To        *Router
	Bandwidth float64
	Network   string // Network CIDR
}

type Network struct {
	Routers map[string]*Router
	Areas   map[string]*Area
}

type Area struct {
	ID      string
	Routers map[string]*Router
}

type LSA struct {
	LinkID            string
	AdvertisingRouter string
	Cost              float64
	AreaID            string
}

func ElectDRAndBDR(routers []*Router) (*Router, *Router) {
	var dr, bdr *Router
	highestPriority := -1
	secondHighestPriority := -1

	for _, router := range routers {
		if router.OSPFPriority > highestPriority {
			bdr = dr
			dr = router
			secondHighestPriority = highestPriority
			highestPriority = router.OSPFPriority
		} else if router.OSPFPriority == highestPriority {
			if bytesCompare(net.ParseIP(dr.ID), net.ParseIP(router.ID)) < 0 {
				bdr = dr
				dr = router
			} else if bdr == nil || bytesCompare(net.ParseIP(bdr.ID), net.ParseIP(router.ID)) < 0 {
				bdr = router
			}
		} else if router.OSPFPriority > secondHighestPriority {
			bdr = router
			secondHighestPriority = router.OSPFPriority
		}
	}

	return dr, bdr
}

// CalculateCost
// Cisco Routers calculate cost: Cost = 10^8/Bandwidth(in bps)
func CalculateCost(bandwidth float64) float64 {
	return 1e8 / bandwidth
}

func DetermineRouterID(router *Router) string {
	var ips []net.IP

	// Prioritize loopback IP if present
	if router.LoopbackIPs != nil {
		ips = append(ips, router.LoopbackIPs...)
	} else {
		// Add all interface IPs
		ips = append(ips, router.InterfaceIPs...)
	}

	// Sort IP addresses in descending order
	sort.Slice(ips, func(i, j int) bool {
		return bytesCompare(ips[i], ips[j]) > 0
	})

	// The highest IP is the first in the sorted list
	if len(ips) > 0 {
		return ips[0].String()
	}

	return ""
}

// Helper function to compare two IP addresses
func bytesCompare(a, b net.IP) int {
	return bytesCompareBytes(a.To4(), b.To4())
}

func bytesCompareBytes(a, b []byte) int {
	for i := range a {
		if a[i] > b[i] {
			return 1
		}
		if a[i] < b[i] {
			return -1
		}
	}
	return 0
}

func (router *Router) GenerateLSA() {
	for neighborID, link := range router.Links {
		lsa := &LSA{
			LinkID:            neighborID,
			AdvertisingRouter: router.ID,
			Cost:              CalculateCost(link.Bandwidth),
			AreaID:            router.Area,
		}
		router.LSDB[neighborID] = lsa
	}
}

func (area *Area) FloodLSAs() {
	for _, router := range area.Routers {
		for _, lsa := range router.LSDB {
			for _, otherRouter := range area.Routers {
				if otherRouter.ID != lsa.AdvertisingRouter {
					otherRouter.LSDB[lsa.LinkID] = lsa
				}
			}
		}
	}
}
