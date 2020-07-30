package models

// ActionTag model
type ActionTag struct {
	AccountTagID int `orm:"column(account_tag_id);pk"`
	AccountID    int `orm:"column(account_id)"`
	TagID        int `orm:"column(tag_id)"`
}

// TableName for beego orm
func (a *ActionTag) TableName() string {
	const ormTableName = "account_tags"
	return ormTableName
}

// NewActionTag factory method
func NewActionTag(accountTagID int, accountID int, tagID int) *ActionTag {
	return &ActionTag{
		AccountTagID: accountTagID,
		AccountID:    accountID,
		TagID:        tagID,
	}
}
