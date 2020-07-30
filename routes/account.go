package routes

import (
	"fmt"
	"log"
	"rc_go/models"
	"rc_go/repositories"
	apiModel "rc_go/routes/models"
	httputil "rc_go/utils/http"

	"github.com/gin-gonic/gin"
)

// for read request body only once ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

// GetUser return user info by id
func GetUser(c *gin.Context) {
	// accountID, err := strconv.Atoi(c.Query("accountID"))
	// if err != nil {
	// 	log.Println("failed to convert string to int in get user")
	// }

	// express := cast.ToBool(c.Query("express"))

	// reqParams := &apiModel.GetUserParams{AccountID: accountID}
	// if err := httputil.RequestValidator(reqParams); err != nil {
	// 	httputil.ValidateErrorWithDataResponse(c, &httputil.DataMap{"desc": err.Error()})
	// 	return
	// }
	reqParams := &apiModel.GetUserParams{}
	err := c.ShouldBindQuery(reqParams)
	fmt.Println(reqParams)
	if err != nil {
		fmt.Println(err)
		// log.Println("failed to validate get user params")
		httputil.ValidateErrorResponse(c)
		return
	}

	// c.JSON(http.StatusOK, gin.H{})
	// return
	account, err := repositories.GetAccountRep().Find(&models.Account{BrandID: reqParams.BrandID, AccountName: reqParams.AccountName})
	if err != nil {
		log.Println("failed to get user info by user id", err)
		httputil.ErrorResponse(c, &httputil.ErrorObj{ErrorMsg: "failed to get user by id", ErrorCode: 1})
		return
	}

	httputil.SuccessResponse(c, gin.H{"user": account})
}

// AddUser for add user
func AddUser(c *gin.Context) {
	reqParams := &apiModel.AddUserParams{}
	err := c.ShouldBindJSON(reqParams)
	if err != nil {
		log.Println("failed to validate add user params")
		httputil.ValidateErrorResponse(c)
		return
	}

	account, err := repositories.GetAccountRep().Insert(&reqParams.Account)
	if err != nil {
		log.Println("failed to add account")
		httputil.ErrorResponse(c, gin.H{"err": "failed to add account"})
		return
	}

	httputil.SuccessResponse(c, gin.H{"user": account})
	return
}
