package models

// RuleTag model
type RuleTag struct {
	RuleTagID int `orm:"column(rule_tag_id);pk"`
	RuleID    int `orm:"column(rule_id)"`
	TagID     int `orm:"column(tag_id)"`
}

// TableName for beego orm
func (r *RuleTag) TableName() string {
	const ormTableName = "rule_tags"
	return ormTableName
}

// NewRuleTag factory method
func NewRuleTag(ruleID int, tagID int) *RuleTag {
	return &RuleTag{
		RuleID: ruleID,
		TagID:  tagID,
	}
}
