package aws

import (
	"testing"
)

func TestParse(t *testing.T) {
	var region awsRegion
	regions := parseRawRegions()

	for i := 0; i < len(regions); i++ {
		if regions[i].id == "us-east-1" {
			region = regions[i]
		}
	}

	if region == (awsRegion{}) {
		t.Error("Expected to find us-east-1, found nothing")
	}

	expectedName := "US East (N. Virginia)"
	if region.name != expectedName {
		t.Errorf("Expected to have name [%s], found [%s]", expectedName, region.name)
	}
	expectedUrl := "rds.us-east-1.amazonaws.com"
	if region.url != expectedUrl {
		t.Errorf("Expected to have url [%s], found [%s]", expectedUrl, region.url)
	}
}

func TestLatency(t *testing.T) {
	latency := latencyOf("localhost")
	if latency.Milliseconds() == 0 {
		t.Error("Expected to get a latency for localhost")
	}

	if latency.Milliseconds() > 100 {
		t.Error("Expected to have a low latency on localhost")
	}
}
