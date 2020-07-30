package models

// Brand model
type Brand struct {
	BrandID   int    `orm:"column(brand_id);pk"`
	BrandName string `orm:"column(brand_name)"`
}

// TableName for beego orm
func (b *Brand) TableName() string {
	const ormTableName = "brands"
	return ormTableName
}

// NewBrand factory method
func NewBrand(brandID int, brandName string) *Brand {
	return &Brand{
		BrandID:   brandID,
		BrandName: brandName,
	}
}
