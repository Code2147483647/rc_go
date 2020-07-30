package models

// Rule model
type Rule struct {
	RuleID   int    `orm:"column(rule_id);pk"`
	BrandID  int    `orm:"column(brand_id)"`
	RuleName string `orm:"column(rule_name)"`
	ActionID int    `orm:"column(action_id:)"`
}

// TableName for beego orm
func (r *Rule) TableName() string {
	const ormTableName = "rules"
	return ormTableName
}

// NewRule factory method
func NewRule(ruleID int, brandID int, ruleName string, actionID int) *Rule {
	return &Rule{
		RuleID:   ruleID,
		BrandID:  brandID,
		RuleName: ruleName,
		ActionID: actionID,
	}
}
