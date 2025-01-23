# Nimbostratus

[![Go Version](https://img.shields.io/github/go-mod/go-version/gregoriokusowski/nimbostratus)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/github/license/gregoriokusowski/nimbostratus)](LICENSE)

> A nimbostratus cloud is a multi-level, gray, often dark, amorphous, nearly uniform cloud that usually produces continuous rain, snow, or sleet but no lightning or thunder. [Wikipedia](https://en.wikipedia.org/wiki/Nimbostratus_cloud)

Nimbostratus is a quick (and not so reliable) tool to check which cloud region you should use based on latency.
Currently only supports AWS.

## Requirements

- Go 1.16 or higher

## Usage

The tool can be used in two different ways: via CLI or as a Go Library.

### CLI

```bash
go install github.com/gregoriokusowski/nimbostratus/cmd/nimbostratus@latest
```

Run the command to see a sorted list of AWS regions by latency from your current location:

```bash
nimbostratus
# => Region                       Id                    Ping
# => Europe (Frankfurt)           eu-central-1    [    26ms]
# => Europe (Paris)               eu-west-3       [    28ms]
# => Europe (Stockholm)           eu-north-1      [    37ms]
# => ...
```

The output shows:
- Region: Human readable AWS region name
- Id: AWS region identifier
- Ping: Measured latency in milliseconds

### Library

First, add the dependency to your project:

```bash
go get github.com/gregoriokusowski/nimbostratus
```

Then use it in your code:

```go
package main

import (
	"context"
	"time"

	"github.com/gregoriokusowski/nimbostratus/aws"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	regions := aws.GetRegions(ctx)
	for _, region := range regions {
		...
	}
}
```

## Information

By _not so reliable_ I mean it's just checking a few requests and their response times.
If you plan to deploy something to production you should not use this to support your assumptions.

### Why?

I personally need this for another project, and this could be easily extracted. So... why not?

If you want something more complete:
* https://github.com/reoim/pingcloud-cli

If you want something that you can run from your browser:
* https://www.cloudping.info/

### How?

Nimbostratus relies on HTTP requests to public RDS endpoints.

# License

nimbostratus is released under The MIT License (MIT).
