package structfunctions

import (
	"fmt"
	"strings"
)

type StaticRoute struct {
	Spec   StaticRouteSpec   `json:"spec,omitempty"`
	Status StaticRouteStatus `json:"status,omitempty"`
}

type StaticRouteStatus struct {
	Nodes []string `json:"nodes,omitempty"`
}

type StaticRouteSpec struct {
	AddressFamily string            `json:"addressFamily,omitempty"`
	DestRoute     string            `json:"destRoute,omitempty"`
	ViaDev        string            `json:"viaDev,omitempty"`
	NextHop       string            `json:"nextHop,omitempty"`
	NodeSelector  map[string]string `json:"nodeSelector"`
}

func GenStaticRoute() StaticRoute {
	nodeSelector1 := map[string]string{
		"multi-node-cluster-worker":  "true",
		"multi-node-cluster-worker2": "false",
	}
	staticRouteSpec1 := StaticRouteSpec{AddressFamily: "IPv4",
		DestRoute:    "192.168.123.32/32",
		ViaDev:       "eth0",
		NextHop:      "1.1.1.1",
		NodeSelector: nodeSelector1,
	}
	staticRouteStatus1 := StaticRouteStatus{Nodes: []string{"worker1", "worker2", "worker3"}}
	staticRoute := StaticRoute{Spec: staticRouteSpec1,
		                       Status: staticRouteStatus1,
	}
	return staticRoute
}

func PrintNodeSelector(staticRoute StaticRoute) {
	for noodeName, routeNeeded := range staticRoute.Spec.NodeSelector {
		if strings.Contains(noodeName, "worker") && strings.Contains(routeNeeded, "true") {
			fmt.Println("DEBUG4 --- :", "noodeName", noodeName, "routeNeeded", routeNeeded)
		}
	}
}

func PrintNodeSelectorOfStaticRoute() {
	staticRoute1 := GenStaticRoute()
	PrintNodeSelector(staticRoute1)
}