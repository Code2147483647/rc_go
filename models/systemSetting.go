package models

// SystemSetting model
type SystemSetting struct {
	PropertyID int `orm:"column(property_id);pk"`
	BrandID  int    `orm:"column(brand_id)"`
	PropertyName string `orm:"column(property_name)"`
	PropertyValue string `orm:"column(property_value)"`
}

// TableName for beego orm
func (s *SystemSetting) TableName() string {
	const ormTableName = "system_setting"
	return ormTableName
}

// NewSystemSetting factory method
func NewSystemSetting(brandID int, propertyName string, propertyValue string) *SystemSetting {
	return &SystemSetting{
		BrandID: brandID,
		PropertyName: propertyName,
		PropertyValue: propertyValue,
	}
}