package events

type Events struct {
	Data  []event `json:"data"`
	Count int32   `json:"count"`
	Page  int32   `json:"page"`
}

type event struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Organizer   string `json:"organizer"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Website     string `json:"website"`
	Email       string `json:"email"`
	Venue       string `json:"venue"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Screenshot  string `json:"screenshot"`
}
