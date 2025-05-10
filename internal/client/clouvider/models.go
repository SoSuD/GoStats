package clouvider

type Service struct {
	ID           string `json:"id"`
	Domain       string `json:"domain"`
	Total        string `json:"total"`
	Status       string `json:"status"`
	BillingCycle string `json:"billingcycle"`
	NextDue      string `json:"next_due"`
	Category     string `json:"category"`
	CategoryURL  string `json:"category_url"`
	Name         string `json:"name"`
}

type ListResponse struct {
	Services []Service `json:"services"`
}

type BalanceModel struct {
	Success bool `json:"success"`
	Details struct {
		Currency   string      `json:"currency"`
		AccBalance interface{} `json:"acc_balance"`
		AccCredit  string      `json:"acc_credit"`
	} `json:"details"`
}
