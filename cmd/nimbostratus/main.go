package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gregoriokusowski/nimbostratus/aws"
)

const BASE_FORMAT = "%-28s %-16s"
const HEADER_FORMAT = BASE_FORMAT + "%10s\n"
const ZONE_FORMAT = BASE_FORMAT + "[%6dms]\n"

func main() {
	fmt.Printf(HEADER_FORMAT, "Zone", "Id", "Ping")
	ctx, _ := context.WithTimeout(context.TODO(), 500*time.Millisecond)
	zones := aws.GetZones(ctx)
	for _, zone := range zones {
		fmt.Printf(ZONE_FORMAT, zone.Name, zone.Id, zone.LatencyInMillis)
	}
}
