package sqlx 

import (
	"log"
	"rc_go/models"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	createAccount  = `insert into accounts(brand_id, account_name, risk_level_id, express) values(:brand_id, :account_name, :risk_level_id, :express)`
	createAccount2 = `insert into accounts(brand_id, account_name, risk_level_id, express) values(?,?,?,?)`
)

// AccountRepository handle account persistance
type AccountRepository struct {
	db *sqlx.DB
}

// NewAccountRepository factory method
func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Insert account
func (a *AccountRepository) Insert(account *models.Account) (*models.Account, error) {
	args := []interface{}{account.BrandID, account.AccountName, account.RiskLevelID, account.Express}
	res, err := a.db.Exec(createAccount2, args...)
	// res, err := a.db.NamedExec(createAccount, account)
	if err != nil {
		log.Println("err: ", err)
		return nil, errors.Errorf("failed to insert account")
	}

	insertID, _ := res.LastInsertId()
	account.AccountID = int(insertID)
	return account, nil
}

// Find account
// func (a *AccountRepository) Find(account *models.Account) (*models.Account, error) {
// 	err := a.db.Read(account)
// 	if err == orm.ErrNoRows {
// 		return nil, nil
// 	} else if err != nil {
// 		log.Println("err: ", err)
// 		return nil, errors.Errorf("failed to fetch account")
// 	}

// 	return account, nil
// }

// // Delete account
// func (a *AccountRepository) Delete(account *models.Account) error {
// 	_, err := a.db.Delete(account)
// 	if err != nil {
// 		log.Println("err: ", err)
// 		return errors.Errorf("failed to fetch account")
// 	}

// 	return nil
// }

// // FindMany find accounts by account_id
// func (a *AccountRepository) FindMany(accountIds []int) ([]*models.Account, error) {
// 	var accounts []*AccountExtension
// 	_, err := a.getQuerySelector().Filter("account_id__in", accountIds).All(&accounts)
// 	if err == orm.ErrNoRows {
// 		return nil, nil
// 	} else if err != nil {
// 		log.Println("err: ", err)
// 		return nil, errors.Errorf("failed to fetch account")
// 	}
// 	var mAccounts []*models.Account
// 	for i := range accounts {
// 		mAccounts = append(mAccounts, accounts[i].ToModel())
// 	}
// 	return mAccounts, nil
// }
