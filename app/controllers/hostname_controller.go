package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/aiyanhamid/mta-hosting-optimizer/app/models"
)

func GetHostnamesWithLessOrEqualActiveIPs(w http.ResponseWriter, r *http.Request) {
	activeIPsStr := os.Getenv("X")
	activeIPs, err := strconv.Atoi(activeIPsStr)
	if err != nil {
		activeIPs = 1
	}

	// Fetch IP configuration data (mocked)
	ipConfig := models.GetMockIPConfig()

	hostnames := make([]string, 0)

	for _, ip := range ipConfig.IPs {
		if len(ip.ActiveIPs) <= activeIPs {
			hostnames = append(hostnames, ip.Hostname)
		}
	}

	response := map[string]interface{}{
		"hostnames": hostnames,
	}

	json.NewEncoder(w).Encode(response)
}
