package wxmch_api

import "fmt"

/*
	当请求处理失败时，除了HTTP状态码表示错误之外，API将在消息体返回错误相应说明具体的错误原因。
	code：详细错误码
	message：错误描述，使用易理解的文字表示错误的原因。
	field: 指示错误参数的位置。当错误参数位于请求body的JSON时，填写指向参数的JSON Pointer 。当错误参数位于请求的url或者querystring时，填写参数的变量名。
	value:错误的值
	issue:具体错误原因
 */
type ErrBody struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Detail ErrDetail `json:"detail"`
}

type ErrDetail struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Issue string `json:"issue"`
	Location string `json:"location"`
}

func (e *ErrBody) Error() string {
	return fmt.Sprintf("微信支付错误码:%s,错误描述:%s", e.Code, e.Message)
}

func (e *ErrBody) IsIdempotent() bool{
	return e.Code == "RESOURCE_ALREADY_EXISTS"
}

type errIdempotent interface {
	IsIdempotent() bool
}