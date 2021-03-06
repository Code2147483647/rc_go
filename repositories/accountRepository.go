package repositories

import (
	"log"
	"rc_go/models"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

// Account rename
type Account = models.Account

var once sync.Once
var instance *AccountRepository

// AccountRepository handle account persistance
type AccountRepository struct {
	db orm.Ormer
}

// getQuerySelector
func (a AccountRepository) getQuerySelector() orm.QuerySeter {
	return a.db.QueryTable("accounts")
}

// NewAccountRepository factory method
func NewAccountRepository(db orm.Ormer) *AccountRepository {
	instance = &AccountRepository{db: db}
	return instance
}

// GetAccountRep return accountRepository instance
func GetAccountRep() *AccountRepository {
	return instance
}

// Insert account
func (a *AccountRepository) Insert(account *Account) (*Account, error) {
	res, err := a.db.Insert(account)
	if err != nil {
		log.Println("err: ", err)
		return nil, errors.Errorf("failed to insert account")
	}

	account.AccountID = int(res)
	return account, nil
}

// Find account
func (a *AccountRepository) Find(account *Account) (*Account, error) {
	err := a.db.Read(account, "account_name", "brand_id", "account_id")
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println("err: ", err)
		return nil, errors.Errorf("failed to fetch account")
	}

	return account, nil
}

// Delete account
func (a *AccountRepository) Delete(account *Account) error {
	_, err := a.db.Delete(account)
	if err != nil {
		log.Println("err: ", err)
		return errors.Errorf("failed to fetch account")
	}

	return nil
}

// FindMany find accounts by account_id
func (a *AccountRepository) FindMany(accountIds []int) ([]*Account, error) {
	var accounts []*Account
	_, err := a.getQuerySelector().Filter("account_id__in", accountIds).All(&accounts)
	if err == orm.ErrNoRows {
		return accounts, nil
	} else if err != nil {
		log.Println("err: ", err)
		return nil, errors.Errorf("failed to fetch accounts")
	}

	return accounts, nil
}
