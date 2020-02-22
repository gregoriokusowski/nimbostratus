package aws

import (
	"testing"
)

func TestParse(t *testing.T) {
	var zone awsZone
	zones := parseRawZones()

	for i := 0; i < len(zones); i++ {
		if zones[i].id == "us-east-1" {
			zone = zones[i]
		}
	}

	if zone == (awsZone{}) {
		t.Error("Expected to find us-east-1, found nothing")
	}

	expectedName := "US East (N. Virginia)"
	if zone.name != expectedName {
		t.Errorf("Expected to have name [%s], found [%s]", expectedName, zone.name)
	}
	expectedUrl := "rds.us-east-1.amazonaws.com"
	if zone.url != expectedUrl {
		t.Errorf("Expected to have url [%s], found [%s]", expectedUrl, zone.url)
	}
}

func TestLatency(t *testing.T) {
	latencyInMillis := latencyOf("localhost")
	if latencyInMillis == 0 {
		t.Error("Expected to get a latency for localhost")
	}

	if latencyInMillis > 100 {
		t.Error("Expected to have a low latency on localhost")
	}
}
