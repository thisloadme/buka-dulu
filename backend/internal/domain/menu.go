package domain

type Menu struct {
	ID               string  `db:"id" json:"id"`
	VentureID        string  `db:"venture_id" json:"venture_id"`
	Name             string  `db:"name" json:"name"`
	Description      string  `db:"description" json:"description,omitempty"`
	Status           string  `db:"status" json:"status"` // candidate, active, deferred, dropped
	IsHero           bool    `db:"is_hero" json:"is_hero"`
	ComplexityScore  float64 `db:"complexity_score" json:"complexity_score,omitempty"`
	ComplexityFactors string `db:"complexity_factors" json:"complexity_factors,omitempty"`
	SortOrder        int     `db:"sort_order" json:"sort_order"`
	CreatedAt        string  `db:"created_at" json:"created_at"`
	UpdatedAt        string  `db:"updated_at" json:"updated_at"`
}

type CreateMenuRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateMenuRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	IsHero      *bool   `json:"is_hero,omitempty"`
}
