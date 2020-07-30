package mysql

import (
	"github.com/jinzhu/gorm"
)

type Manager struct {
	db *db.gorm
}