package xerr

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

/**
常用通用固定错误
*/

type CodeError struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

// GetErrCode 返回给前端的错误码

func (e *CodeError) GetErrCode() uint32 {
	return e.Code
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.Message
}

func (e *CodeError) ErrorResponse() CodeError {
	return CodeError{
		Code:    e.Code,
		Message: e.Message,
	}
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Message)
}

func CustomErr(errCode uint32, ctx context.Context, err error) error {
	var codeErr *CodeError
	if IsCodeErr(errCode) {
		codeErr = &CodeError{Code: errCode, Message: MapErrMsg(errCode)}
	} else {
		codeErr = &CodeError{Code: errCode, Message: err.Error()}
	}
	logx.WithContext(ctx).Error(errors.Wrapf(err, "resp[%+v]", codeErr))
	return codeErr
}

func ErrHandler(err error) any {
	switch err.(type) {
	// 如果错误类型为CodeError，就返回错误类型的结构体
	case *CodeError:
		return err
	default:
		// 系统错误，500 错误提示
		return CodeError{Code: 500, Message: err.Error()}
	}
}
