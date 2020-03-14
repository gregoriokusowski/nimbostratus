# Nimbostratus

> A nimbostratus cloud is a multi-level, gray, often dark, amorphous, nearly uniform cloud that usually produces continuous rain, snow, or sleet but no lightning or thunder. [Wikipedia](https://en.wikipedia.org/wiki/Nimbostratus_cloud)

Nimbostratus is a quick (and not so reliable) tool to check which cloud region you should use based on latency.
Currently only supports AWS.

## Usage

The tool can be used in two different ways: via CLI or as a Go Library.

### CLI

Install

```
$ go get github.com/gregoriokusowski/nimbostratus/cmd/nimbostratus
```

Run

```
$ nimbostratus
Region                       Id                    Ping
Europe (Frankfurt)           eu-central-1    [    26ms]
Europe (Paris)               eu-west-3       [    28ms]
Europe (Stockholm)           eu-north-1      [    37ms]
...
```

### Library

Import `github.com/gregoriokusowski/nimbostratus` via your favorite dependency module.

And use it

```go
package whatever

import (
	"github.com/gregoriokusowski/nimbostratus/aws"
)

func bla() {
  regions := aws.GetRegions(context.TODO()) // or configure a timeout with context.WithTimeout
  // ...
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
