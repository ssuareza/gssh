package gssh

import (
	"testing"
)

func TestGetIP(t *testing.T) {
	instances := []Server{
		{
			Name: "server1",
			Values: map[string]string{
				"instance-id": "i-0a541b4374af7d920",
				"private-ip":  "172.16.0.180",
				"public-ip":   "52.204.253.208",
			},
		},
		{
			Name: "server2",
			Values: map[string]string{
				"instance-id": "i-1a432b4374dc7d123",
				"private-ip":  "172.16.0.181",
				"public-ip":   "50.204.253.210",
			},
		},
	}

	ip, err := GetIP("i-0a541b4374af7d920", instances, "public")
	if err != nil {
		t.Error("Not able to get ip")
	}

	expected := "52.204.253.208"
	if ip != expected {
		t.Errorf("IP is %s and should be %s", ip, expected)
	}
}
