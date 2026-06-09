package domain

type Ingredient struct {
	ID        string  `db:"id" json:"id"`
	VentureID string  `db:"venture_id" json:"venture_id"`
	MenuID    string  `db:"menu_id" json:"menu_id"`
	Name      string  `db:"name" json:"name"`
	Unit      string  `db:"unit" json:"unit"`
	Quantity  float64 `db:"quantity" json:"quantity"`
	UnitPrice float64 `db:"unit_price" json:"unit_price"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
}

type PackagingCost struct {
	ID              string  `db:"id" json:"id"`
	VentureID       string  `db:"venture_id" json:"venture_id"`
	MenuID          string  `db:"menu_id" json:"menu_id"`
	Name            string  `db:"name" json:"name"`
	UnitPrice       float64 `db:"unit_price" json:"unit_price"`
	QuantityPerPcs  float64 `db:"quantity_per_pcs" json:"quantity_per_pcs"`
	CreatedAt       string  `db:"created_at" json:"created_at"`
}

type CostSummary struct {
	ID             string  `db:"id" json:"id"`
	VentureID      string  `db:"venture_id" json:"venture_id"`
	MenuID         string  `db:"menu_id" json:"menu_id"`
	HppPerPorsi    float64 `db:"hpp_per_porsi" json:"hpp_per_porsi"`
	SuggestedPrice float64 `db:"suggested_price" json:"suggested_price"`
	TargetMargin   float64 `db:"target_margin" json:"target_margin"`
	GrossMargin    float64 `db:"gross_margin" json:"gross_margin"`
	MarginStatus   string  `db:"margin_status" json:"margin_status"` // sehat, tipis, berbahaya
	BreakEvenUnit  int     `db:"break_even_unit" json:"break_even_unit"`
	LaborPerUnit   float64 `db:"labor_per_unit" json:"labor_per_unit,omitempty"`
	OverheadPerUnit float64 `db:"overhead_per_unit" json:"overhead_per_unit,omitempty"`
	IsLocked       bool    `db:"is_locked" json:"is_locked"`
	CreatedAt      string  `db:"created_at" json:"created_at"`
	UpdatedAt      string  `db:"updated_at" json:"updated_at"`
}

type CreateIngredientRequest struct {
	MenuID    string  `json:"menu_id"`
	Name      string  `json:"name"`
	Unit      string  `json:"unit"`
	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type CreatePackagingRequest struct {
	MenuID         string  `json:"menu_id"`
	Name           string  `json:"name"`
	UnitPrice      float64 `json:"unit_price"`
	QuantityPerPcs float64 `json:"quantity_per_pcs"`
}

type CostConfigRequest struct {
	LaborPerUnit    *float64 `json:"labor_per_unit,omitempty"`
	OverheadPerUnit *float64 `json:"overhead_per_unit,omitempty"`
	TargetMargin    *float64 `json:"target_margin,omitempty"` // percentage
}
