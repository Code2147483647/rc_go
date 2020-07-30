package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

// DataMap for wraping data
type DataMap = map[string]interface{}

// ErrorObj represent error
type ErrorObj struct {
	ErrorMsg  string                  `json:"error_msg"`
	ErrorCode int                     `json:"error_code"`
	Data      *map[string]interface{} `json:"data"`
	Success   bool                    `json:"success"`
}

// NewValidateError for errorObj that represent validate error to client
func NewValidateError(data *DataMap) *ErrorObj {
	return &ErrorObj{ErrorMsg: "params is invalid", ErrorCode: 1, Success: false, Data: data}
}

// RequestValidator for validate api
func RequestValidator(data interface{}) error {
	return validate.Struct(data)
}

// SuccessResponse response to client for success
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// ValidateErrorResponse error response for validate error
func ValidateErrorResponse(c *gin.Context) {
	c.JSON(http.StatusOK, NewValidateError(nil))
}

// ValidateErrorWithDataResponse error response for validate error
func ValidateErrorWithDataResponse(c *gin.Context, data *DataMap) {
	c.JSON(http.StatusOK, NewValidateError(data))
}

// ErrorResponse response to client for error
func ErrorResponse(c *gin.Context, errObj interface{}) {
	c.JSON(http.StatusOK, errObj)
}
