package repositories

import "github.com/astaxie/beego/orm"

// InitRep init repository
func InitRep(dbInstance orm.Ormer) {
	NewAccountRepository(dbInstance)
}
