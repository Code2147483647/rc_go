package models

// RuleAction model
type RuleAction struct {
	RuleActionID int `orm:"column(rule_action_id);pk"`
	RuleID       int `orm:"column(rule_id)"`
	ActionID     int `orm:"column(account_id)"`
}

// TableName for beego orm
func (r *RuleAction) TableName() string {
	const ormTableName = "rule_actions"
	return ormTableName
}

// NewRuleAction factory method
func NewRuleAction(ruleID, int, actionID int) *RuleAction {
	return &RuleAction{
		RuleID:   ruleID,
		ActionID: actionID,
	}
}
