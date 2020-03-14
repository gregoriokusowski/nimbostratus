package main

import (
	"context"
	"fmt"

	"github.com/gregoriokusowski/nimbostratus/aws"
)

const BASE_FORMAT = "%-28s %-16s"
const HEADER_FORMAT = BASE_FORMAT + "%10s\n"
const ZONE_FORMAT = BASE_FORMAT + "[%6dms]\n"

func main() {
	fmt.Printf(HEADER_FORMAT, "Zone", "Id", "Ping")
	zones := aws.GetZones(context.TODO())
	for _, zone := range zones {
		fmt.Printf(ZONE_FORMAT, zone.Name, zone.Id, zone.LatencyInMillis)
	}
}
