package dbmodels

type ProductCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// TableName ...
func (t *ProductCategory) TableName() string {
	return "public.product_category"
}
