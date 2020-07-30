package models

import "rc_go/models"

// GetUserParams for getUser req
// type GetUserParams struct {
// 	AccountID int `json:"accountID" validate:"required"`
// }

// GetUserParams for getUser req
type GetUserParams struct {
	BrandID     int    `form:"brand_id" binding:"required"`
	AccountName string `form:"account_name" binding:"required"`
}

// AddUserParams add User req
type AddUserParams struct {
	models.Account
	AccountID int `json:"account_id,omitempty"`
}
