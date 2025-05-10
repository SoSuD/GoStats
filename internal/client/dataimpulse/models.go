package dataimpulse

type TrafficModel struct {
	TotalTraffic int64       `json:"total_traffic"`
	TrafficUsed  int64       `json:"traffic_used"`
	TrafficLeft  int64       `json:"traffic_left"`
	UsedThreads  int         `json:"used_threads"`
	Login        string      `json:"login"`
	Status       string      `json:"status"`
	Message      interface{} `json:"message"`
	Elapsed      string      `json:"elapsed"`
}
