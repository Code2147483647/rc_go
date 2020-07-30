package models

// Account model
type Account struct {
	AccountID   int    `orm:"column(account_id);pk" json:"account_id"`
	BrandID     int    `orm:"column(brand_id)" json:"brand_id" binding:"required"`
	AccountName string `orm:"column(account_name)" json:"account_name" binding:"required"`
	RiskLevelID int    `orm:"column(risk_level_id)" json:"risk_level_id"`
	Express     bool   `orm:"column(express)" json:"express"`
}

// TableName for beego orm
func (u *Account) TableName() string {
	const ormTableName = "accounts"
	return ormTableName
}

// NewAccount factory method
func NewAccount(brandID int, accountName string, riskLevelID int, express bool) *Account {
	return &Account{
		BrandID:     brandID,
		AccountName: accountName,
		RiskLevelID: riskLevelID,
		Express:     express,
	}
}
