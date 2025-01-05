package xerr

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func ErrorHandler(err error) (int, any) {
	// 因为只在参数校验处使用了httpx的error,所以这里只会处理参数校验错误,所以并没有打印日志
	// 如果需要打印日志,可以在这里打印
	return http.StatusOK, nil
}

func DebugErrorHandler(err error) (int, any) {

	return http.StatusOK, Response{
		Code: 100,
		Msg:  err.Error(),
	}

}

// OkHandler 请求成功回调
func OkHandler(ctx context.Context, result any) any {
	return Response{
		Code:   200,
		Msg:    "ok",
		Result: result,
	}
}

func CodeFromError(err error) Code {
	err = errors.Cause(err)
	if code, ok := err.(Code); ok {
		return code
	}

	switch err {
	case context.Canceled:
		return Canceled
	case context.DeadlineExceeded:
		return Deadline
	}

	return ServerErr
}

// JwtError jwt校验回调
func JwtError(w http.ResponseWriter, err error) {
	var jwtErrorResponse Code

	jwtError := err.Error()
	switch jwtError {
	case JwtNotExists:
		jwtErrorResponse = NotAuth
	case JWTTokenExpired:
		jwtErrorResponse = AuthExpired
	default:
		jwtErrorResponse = NotAuth
	}

	response, err := json.Marshal(map[string]interface{}{
		"code": jwtErrorResponse.Code(),
		"msg":  jwtErrorResponse.Message(),
	})

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

// IsDuplicateEntryError 判断是否为重复插入错误
func IsDuplicateEntryError(err error) bool {
	if strings.Contains(err.Error(), "Duplicate entry") {
		return true
	}
	return false
}

// ConfPanicError 配置文件错误(请勿在其他代码中使用,仅用于配置文件错误)
func ConfPanicError(msg string) {
	panic(fmt.Sprintf("\033[31m%s\u001B[0m\n", msg))
}
