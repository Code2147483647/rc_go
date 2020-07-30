package models

// Action model
type Action struct {
	ActionID    int    `orm:"column(action_id);pk"`
	BrandID     int    `orm:"column(brand_id)"`
	ActionName  string `orm:"column(action_name)"`
	RiskLevelID int    `orm:"column(risk_level_id)"`
	ActionCode  string `orm:"column(action_code)"`
}

// TableName for beego orm
func (a *Action) TableName() string {
	const ormTableName = "actions"
	return ormTableName
}

// NewAction factory method
func NewAction(brandID int, actionName string, riskLevelID int, actionCode string) *Action {
	return &Action{
		BrandID:     brandID,
		ActionName:  actionName,
		RiskLevelID: riskLevelID,
		ActionCode:  actionCode,
	}
}
