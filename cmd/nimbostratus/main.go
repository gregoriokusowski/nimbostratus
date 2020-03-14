package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gregoriokusowski/nimbostratus/aws"
)

const BASE_FORMAT = "%-28s %-16s"
const HEADER_FORMAT = BASE_FORMAT + "%10s\n"
const REGION_FORMAT = BASE_FORMAT + "[%6dms]\n"

func main() {
	fmt.Printf(HEADER_FORMAT, "Region", "Id", "Ping")
	ctx, _ := context.WithTimeout(context.TODO(), 500*time.Millisecond)
	regions := aws.GetRegions(ctx)
	for _, region := range regions {
		fmt.Printf(REGION_FORMAT, region.Name, region.Id, region.Latency.Milliseconds())
	}
}
