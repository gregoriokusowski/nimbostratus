package aws

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gregoriokusowski/nimbostratus"
)

// Based on https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html
const awsRawData = `
Region Name	Region	Endpoint	Protocol
US East (Ohio)	us-east-2	rds.us-east-2.amazonaws.com	HTTPS
US East (N. Virginia)	us-east-1	rds.us-east-1.amazonaws.com	HTTPS
US West (N. California)	us-west-1	rds.us-west-1.amazonaws.com	HTTPS
US West (Oregon)	us-west-2	rds.us-west-2.amazonaws.com	HTTPS
Africa (Cape Town)	af-south-1	rds.af-south-1.amazonaws.com	HTTPS
Asia Pacific (Hong Kong)	ap-east-1	rds.ap-east-1.amazonaws.com	HTTPS
Asia Pacific (Hyderabad)	ap-south-2	rds.ap-south-2.amazonaws.com	HTTPS
Asia Pacific (Jakarta)	ap-southeast-3	rds.ap-southeast-3.amazonaws.com	HTTPS
Asia Pacific (Malaysia)	ap-southeast-5	rds.ap-southeast-5.amazonaws.com	HTTPS
Asia Pacific (Melbourne)	ap-southeast-4	rds.ap-southeast-4.amazonaws.com	HTTPS
Asia Pacific (Mumbai)	ap-south-1	rds.ap-south-1.amazonaws.com	HTTPS
Asia Pacific (Osaka)	ap-northeast-3	rds.ap-northeast-3.amazonaws.com	HTTPS
Asia Pacific (Seoul)	ap-northeast-2	rds.ap-northeast-2.amazonaws.com	HTTPS
Asia Pacific (Singapore)	ap-southeast-1	rds.ap-southeast-1.amazonaws.com	HTTPS
Asia Pacific (Sydney)	ap-southeast-2	rds.ap-southeast-2.amazonaws.com	HTTPS
Asia Pacific (Thailand)	ap-southeast-7	rds.ap-southeast-7.amazonaws.com	HTTPS
Asia Pacific (Tokyo)	ap-northeast-1	rds.ap-northeast-1.amazonaws.com	HTTPS
Canada (Central)	ca-central-1	rds.ca-central-1.amazonaws.com	HTTPS
Canada West (Calgary)	ca-west-1	rds.ca-west-1.amazonaws.com	HTTPS
Europe (Frankfurt)	eu-central-1	rds.eu-central-1.amazonaws.com	HTTPS
Europe (Ireland)	eu-west-1	rds.eu-west-1.amazonaws.com	HTTPS
Europe (London)	eu-west-2	rds.eu-west-2.amazonaws.com	HTTPS
Europe (Milan)	eu-south-1	rds.eu-south-1.amazonaws.com	HTTPS
Europe (Paris)	eu-west-3	rds.eu-west-3.amazonaws.com	HTTPS
Europe (Spain)	eu-south-2	rds.eu-south-2.amazonaws.com	HTTPS
Europe (Stockholm)	eu-north-1	rds.eu-north-1.amazonaws.com	HTTPS
Europe (Zurich)	eu-central-2	rds.eu-central-2.amazonaws.com	HTTPS
Israel (Tel Aviv)	il-central-1	rds.il-central-1.amazonaws.com	HTTPS
Mexico (Central)	mx-central-1	rds.mx-central-1.amazonaws.com	HTTPS
Middle East (Bahrain)	me-south-1	rds.me-south-1.amazonaws.com	HTTPS
Middle East (UAE)	me-central-1	rds.me-central-1.amazonaws.com	HTTPS
South America (São Paulo)	sa-east-1	rds.sa-east-1.amazonaws.com	HTTPS
AWS GovCloud (US-East)	us-gov-east-1	rds.us-gov-east-1.amazonaws.com	HTTPS
AWS GovCloud (US-West)	us-gov-west-1	rds.us-gov-west-1.amazonaws.com	HTTPS
`

type awsRegion struct {
	id   string
	name string
	url  string
}

func GetRegions(ctx context.Context) []nimbostratus.Region {
	var regions []nimbostratus.Region
	var wg sync.WaitGroup
	regc := make(chan nimbostratus.Region)

	for _, option := range parseRawRegions() {
		wg.Add(1)
		go func(o awsRegion) {
			defer wg.Done()
			regc <- nimbostratus.Region{
				Id:      o.id,
				Name:    o.name,
				Latency: latencyOf(o.url),
			}
		}(option)
	}

	execution := make(chan bool)

	go func() {
		for {
			select {
			case region := <-regc:
				regions = append(regions, region)
			case <-ctx.Done():
				execution <- true
			}
		}
	}()

	go func() {
		wg.Wait()
		execution <- true
	}()

	<-execution
	sort.Sort(byLatency(regions))
	return regions
}

func latencyOf(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get("http://" + url)
	finish := time.Now()
	return finish.Sub(start)
}

type byLatency []nimbostratus.Region

func (a byLatency) Len() int      { return len(a) }
func (a byLatency) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byLatency) Less(i, j int) bool {
	return a[i].Latency.Nanoseconds() < a[j].Latency.Nanoseconds()
}

func parseRawRegions() []awsRegion {
	var regions []awsRegion
	for _, line := range strings.Split(strings.TrimSuffix(awsRawData, "\n"), "\n") {
		if len(line) > 0 && !strings.Contains(line, "Region") {
			values := strings.Split(line, "\t")
			regions = append(regions, awsRegion{
				name: strings.TrimSpace(values[0]),
				id:   strings.TrimSpace(values[1]),
				url:  strings.TrimSpace(values[2]),
			})
		}
	}
	return regions
}
