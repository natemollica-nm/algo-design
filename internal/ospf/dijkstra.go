package ospf

import "math"

/*
	Dijkstra's Algorithm:
		1. Mark all nodes as unvisited. Create set of all unvisited nodes called unvisited set.
		2. Assign to every node a tentative distance value: set it to 0 for initial nodes and to ∞ (infinity)
		   for all other nodes. Set initial node as the current node.
		3. For the current node, consider all its unvisited neighbors and calculate their tentative distances
 		   (edges's weight) through the current node. Compare the newly calculated tentative distance to the
		   current assigned value and assign the smaller one.
		4. When done considering all unvisited neighbors of the current node, mark the current node as visited
		   and remove it from the unvisited set. A visited node will never be checked again.
		5. If the destination node has been marked visited (when planning a route between two specific nodes) or if the
		   smallest tentative distance among the nodes in the unvisited set is infinity (i.e., when planning a complete
		   traversal; when no connection between the current node and remaining unvisited nodes), then stop.
		6. The algorithm is finished. Otherwise select the unvisited node that is marked with the smallest tentative
		   distance, set it as the new "current node", and go back to step 3.

	Requirements:
		- Graphs must have weight values > 0


	Usages:
		- OSPF (Open Shortest Path First) Routing Protocol - Link State Routing Protocol
			- Classless Routing Protocol
			- Supports VLSM, CIDR, Manual Route Summarization, and Equal Cost LB
		 	- Incremental Updates Supported
		   	- Single Metric Parameter - Interface Cost used for edge connection weights
			- OSPF Route Administrative Distance (Default: 110)
			- Leverages Multicast Addresses - 224.0.0.5 and 224.0.0.6 for Routing Updates
		- OSPF Terminology
			- Router ID: Highest active IP address on Router
				1. Highest Loopback Address
				2. Highest IP Address on any active interface of Router (if no Loopbacks exist)
			- Router Priority: 8 bit value assigned to router operating OSPF that's used to elect DR and BDR in broadcast network
			- DR (Designated Router): Router elected to minimize the number of adjacencies formed
				- Distributes LSA (Link State Advertisements) to other routers
				- Elected in broadcast network shared by other routers for their DBD
				- DR responds to update requests from other routers for their DBDs
*/

func Dijkstra(network *Network, startID string) map[string]float64 {
	dist := make(map[string]float64)
	for routerID := range network.Routers {
		dist[routerID] = math.Inf(1)
	}
	dist[startID] = 0

	visited := make(map[string]bool)

	for len(visited) < len(network.Routers) {
		// Find the unvisited router with the smallest distance
		minDist := math.Inf(1)
		var minRouter *Router
		for routerID, router := range network.Routers {
			if !visited[routerID] && dist[routerID] < minDist {
				minDist = dist[routerID]
				minRouter = router
			}
		}

		// Visit the router
		visited[minRouter.ID] = true

		// Update distances to neighboring routers
		for neighborID, link := range minRouter.Links {
			if !visited[neighborID] {
				newDist := dist[minRouter.ID] + CalculateCost(link.Bandwidth)
				if newDist < dist[neighborID] {
					dist[neighborID] = newDist
				}
			}
		}
	}

	return dist
}

func DijkstraWithMultiPath(network *Network, startID string, areaID string) map[string]float64 {
	area := network.Areas[areaID]
	dist := make(map[string]float64)
	for routerID := range area.Routers {
		dist[routerID] = math.Inf(1)
	}
	dist[startID] = 0

	visited := make(map[string]bool)
	paths := make(map[string][]string)

	for len(visited) < len(area.Routers) {
		minDist := math.Inf(1)
		var minRouter *Router
		for routerID, router := range area.Routers {
			if !visited[routerID] && dist[routerID] < minDist {
				minDist = dist[routerID]
				minRouter = router
			}
		}

		visited[minRouter.ID] = true

		for neighborID, lsa := range minRouter.LSDB {
			if !visited[neighborID] {
				newDist := dist[minRouter.ID] + lsa.Cost
				if newDist < dist[neighborID] {
					dist[neighborID] = newDist
					paths[neighborID] = append(paths[minRouter.ID], minRouter.ID)
				} else if newDist == dist[neighborID] {
					paths[neighborID] = append(paths[neighborID], minRouter.ID)
				}
			}
		}
	}

	return dist
}
