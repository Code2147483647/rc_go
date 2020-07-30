package models

// RuleCondition model
type RuleCondition struct {
	RuleConditionID int    `orm:"column(rule_condition_id);pk"`
	RuleID          int    `orm:"column(rule_id)"`
	Filed           string `orm:"column(filed)"`
	Operator        string `orm:"column(operator)"`
	Value           string `orm:"column(value)"`
}

// TableName for beego orm
func (r *RuleCondition) TableName() string {
	const ormTableName = "rule_conditions"
	return ormTableName
}

// NewRuleCondition factory method
func NewRuleCondition(ruleID int, filed string, operator string, value string) *RuleCondition {
	return &RuleCondition{
		RuleID:   ruleID,
		Filed:    filed,
		Operator: operator,
		Value:    value,
	}
}
