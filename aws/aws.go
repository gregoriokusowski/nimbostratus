package aws

import (
	"context"
	"net/http"
	"strings"
	"time"

	"kusowski.com/nimbostratus"
)

// Based on https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html
const awsRawData = `
Region Name	Region	Endpoint	Protocol
US East (Ohio)	us-east-2	rds.us-east-2.amazonaws.com	HTTPS
US East (N. Virginia)	us-east-1	rds.us-east-1.amazonaws.com	HTTPS
US West (N. California)	us-west-1	rds.us-west-1.amazonaws.com	HTTPS
US West (Oregon)	us-west-2	rds.us-west-2.amazonaws.com	HTTPS
Asia Pacific (Hong Kong)	ap-east-1	rds.ap-east-1.amazonaws.com	HTTPS
Asia Pacific (Mumbai)	ap-south-1	rds.ap-south-1.amazonaws.com	HTTPS
Asia Pacific (Osaka-Local)	ap-northeast-3	rds.ap-northeast-3.amazonaws.com	HTTPS
Asia Pacific (Seoul)	ap-northeast-2	rds.ap-northeast-2.amazonaws.com	HTTPS
Asia Pacific (Singapore)	ap-southeast-1	rds.ap-southeast-1.amazonaws.com	HTTPS
Asia Pacific (Sydney)	ap-southeast-2	rds.ap-southeast-2.amazonaws.com	HTTPS
Asia Pacific (Tokyo)	ap-northeast-1	rds.ap-northeast-1.amazonaws.com	HTTPS
Canada (Central)	ca-central-1	rds.ca-central-1.amazonaws.com	HTTPS
China (Beijing)	cn-north-1	rds.cn-north-1.amazonaws.com.cn	HTTPS
China (Ningxia)	cn-northwest-1	rds.cn-northwest-1.amazonaws.com.cn	HTTPS
Europe (Frankfurt)	eu-central-1	rds.eu-central-1.amazonaws.com	HTTPS
Europe (Ireland)	eu-west-1	rds.eu-west-1.amazonaws.com	HTTPS
Europe (London)	eu-west-2	rds.eu-west-2.amazonaws.com	HTTPS
Europe (Paris)	eu-west-3	rds.eu-west-3.amazonaws.com	HTTPS
Europe (Stockholm)	eu-north-1	rds.eu-north-1.amazonaws.com	HTTPS
Middle East (Bahrain)	me-south-1	rds.me-south-1.amazonaws.com	HTTPS
South America (Sao Paulo)	sa-east-1	rds.sa-east-1.amazonaws.com	HTTPS
AWS GovCloud (US-East)	us-gov-east-1	rds.us-gov-east-1.amazonaws.com	HTTPS
AWS GovCloud (US-West)	us-gov-west-1	rds.us-gov-west-1.amazonaws.com	HTTPS
`

type awsZone struct {
	id   string
	name string
	url  string
}

func GetZones(_ context.Context) []nimbostratus.Zone {
	var zones []nimbostratus.Zone
	for _, option := range parseRawZones() {
		zones = append(zones, nimbostratus.Zone{
			Id:              option.id,
			Name:            option.name,
			LatencyInMillis: latencyOf(option.url),
		})
	}
	return zones
}

func latencyOf(url string) int64 {
	start := time.Now()
	_, _ = http.Get("http://" + url)
	finish := time.Now()
	return finish.Sub(start).Milliseconds()
}

func parseRawZones() []awsZone {
	var zones []awsZone
	for _, line := range strings.Split(strings.TrimSuffix(awsRawData, "\n"), "\n") {
		if len(line) > 0 && !strings.Contains(line, "Region") {
			values := strings.Split(line, "\t")
			zones = append(zones, awsZone{
				name: strings.TrimSpace(values[0]),
				id:   strings.TrimSpace(values[1]),
				url:  strings.TrimSpace(values[2]),
			})
		}
	}
	return zones
}
