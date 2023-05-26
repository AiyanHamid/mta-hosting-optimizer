package models

type IPConfig struct {
	IPs []IP `json:"ips"`
}

type IP struct {
	Hostname  string   `json:"hostname"`
	ActiveIPs []string `json:"active_ips"`
}

// Mocked IP configuration data
func GetMockIPConfig() IPConfig {
	return IPConfig{
		IPs: []IP{
			{
				Hostname:  "mta1.example.com",
				ActiveIPs: []string{"10.0.0.1", "10.0.0.2"},
			},
			{
				Hostname:  "mta2.example.com",
				ActiveIPs: []string{"10.0.0.3", "10.0.0.4", "10.0.0.5"},
			},
			// Add more mocked data here
		},
	}
}
