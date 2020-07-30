package dao

import (
	"rc_go/models"
)

type AccountDAO interface {
	Insert(account *models.Account) (models.Account, error)
	Find(account *models.Account) ([]*models.Account, error)
	Delete(account *models.Account) error
	FindMany(accountIDs []*int) ([]*models.Account, error)
}