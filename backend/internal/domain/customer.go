package domain

type CustomerSegment struct {
	ID               string `db:"id" json:"id"`
	VentureID        string `db:"venture_id" json:"venture_id"`
	Name             string `db:"name" json:"name"`
	AgeRange         string `db:"age_range" json:"age_range,omitempty"`
	BuyContext       string `db:"buy_context" json:"buy_context,omitempty"`
	BudgetRange      string `db:"budget_range" json:"budget_range,omitempty"`
	Location         string `db:"location" json:"location,omitempty"`
	ConsumptionMoment string `db:"consumption_moment" json:"consumption_moment,omitempty"`
	IsTooGeneral     bool   `db:"is_too_general" json:"is_too_general"`
	IsLocked         bool   `db:"is_locked" json:"is_locked"`
	CreatedAt        string `db:"created_at" json:"created_at"`
	UpdatedAt        string `db:"updated_at" json:"updated_at"`
}

type CreateCustomerSegmentRequest struct {
	Name              string `json:"name"`
	AgeRange          string `json:"age_range,omitempty"`
	BuyContext        string `json:"buy_context,omitempty"`
	BudgetRange       string `json:"budget_range,omitempty"`
	Location          string `json:"location,omitempty"`
	ConsumptionMoment string `json:"consumption_moment,omitempty"`
}
