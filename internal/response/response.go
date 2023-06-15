package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func R400(c *gin.Context, data interface{}, err error) {
	errList := strings.Split(err.Error(), ";")

	errMsg := errList[len(errList)-1] // "username: username invalid"

	formatErrMsg := strings.TrimSpace(strings.Split(errMsg, ":")[1]) // "username invalid"

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
