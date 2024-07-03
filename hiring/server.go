/*
#################
 Leasing Server #
##################
#
# Assume the following snippet of *pseudo-code* implements a leasing server running 
# in its own process/container and exposing some sort of HTTP/gRPC/etc. endpoint, 
# that an external agent can use to lease resources from.
*/

package server

import "io"

var Resources = {
	"Nvidia H100": {
		capacity: 100
		leased:   0
	},
	"Nvidia H200": {
		capacity: 50
		leased:   10
	},
	...
	..
	.
}

func Acquire(name string) bool {
	if (name not in Resources) {
		return false
	}

	var resource = Resources[name]
	if (resource.leased >= resource.capacity) {
		return false
	}
	resource.leased += 1

	Resources[name] = resource
}

func Release(name string) bool {
	if (name not in Resources) {
		return false
	}

	var resource = Resources[name]
	resource.leased -= 1

	Resources[name] = resource
}

func DumpReport() {
	var report = ""
	for name in Resources {
		resource = Resources[name]
		report += "${name}: ${resource.leased} - ${resource.capacity}\n"
	}
	var err = io.writeFile("/var/log/resources.log", report)
}
