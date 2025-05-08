package response

import timeext "github.com/kcchandra/golang-hook/pkg/time"

type BaseResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	TimeStamp string      `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
}

func NewBaseResponse(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:      code,
		Message:   message,
		TimeStamp: timeext.NowString(),
		Data:      data,
	}
}
