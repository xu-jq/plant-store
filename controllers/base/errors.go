package base

var (
	ErrOK  = &GeneralResponse{Code: 0, Message: "ErrOK"}
	Err302 = &GeneralResponse{Code: 302, Message: "Err302"}
	Err404 = &GeneralResponse{Code: 404, Message: "Err404"}
	Err500 = &GeneralResponse{Code: 500, Message: "Err500"}

	ErrInputData = &GeneralResponse{Code: 1001, Message: "ErrInputData"}
	ErrSystem    = &GeneralResponse{Code: 1002, Message: "ErrSystem"}
	ErrDatabase  = &GeneralResponse{Code: 1003, Message: "ErrDatabase"}
	ErrNeedLogin = &GeneralResponse{Code: 1004, Message: "ErrNeedLogin"}
	ErrCSRFToken = &GeneralResponse{Code: 1005, Message: "ErrCSRFToken"}
	ErrUser      = &GeneralResponse{Code: 1006, Message: "用户名已存在"}
	ErrLogin     = &GeneralResponse{Code: 1007, Message: "用户密码错误"}
	ErrGet       = &GeneralResponse{Code: 1008, Message: "用户未登录"}
	ErrParse     = &GeneralResponse{Code: 1009, Message: "获取参数错误"}
)
