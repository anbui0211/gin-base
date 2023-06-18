package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func R400(c *gin.Context, data interface{}, err error) {
	errList := strings.Split(err.Error(), ";")

	errMsg := errList[len(errList)-1] // "username: username invalid"

	splitErr := strings.Split(errMsg, ":")

	formatErrMsg := splitErr[0]
	if len(splitErr) > 1 {
		formatErrMsg = strings.TrimSpace(splitErr[1]) // "username invalid"
	}

	if err == nil {
		errMsg = CommonBadRequest
	}

	sendResponse(c, http.StatusBadRequest, data, false, formatErrMsg)
}

func R200(c *gin.Context, data interface{}, msg string) {
	if msg == "" {
		msg = CommonSuccess
	}
	sendResponse(c, http.StatusOK, data, true, msg)
}

const (
	CommonSuccess    = "thành công"
	CommonBadRequest = "dữ liệu không hợp lệ"
)
