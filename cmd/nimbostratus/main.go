package main

import (
	"context"
	"fmt"

	"kusowski.com/nimbostratus/aws"
)

func main() {
	zones := aws.GetZones(context.TODO())
	for _, zone := range zones {
		fmt.Println("%+v", zone)
	}
}
