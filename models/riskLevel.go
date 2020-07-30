package models

// RiskLevel model
type RiskLevel struct {
	RiskLevelID int    `orm:"column(risk_level_id);pk"`
	BrandID     int    `orm:"column(brand_id)"`
	RiskName    string `orm:"column(risk_name)"`
	RiskLevel   int    `orm:"column(risk_level)"`
}

// TableName for beego orm
func (r *RiskLevel) TableName() string {
	const ormTableName = "riskLevels"
	return ormTableName
}

// NewRiskLevel factory method
func NewRiskLevel(brandID int, riskName string, riskLevel int) *RiskLevel {
	return &RiskLevel{
		BrandID:   brandID,
		RiskName:  riskName,
		RiskLevel: riskLevel,
	}
}
