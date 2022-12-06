package common

type Response struct {
	Code    int         `json:"code"`
	ErrMsg  string      `json:"errMsg"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func BuildSuccessResponse(data interface{}) Response {
	return Response{
		Code:    0,
		ErrMsg:  "",
		Success: true,
		Data:    data,
	}
}

func BuildSuccessResponseNoData() Response {
	return Response{
		Code:    0,
		ErrMsg:  "",
		Success: true,
		Data:    nil,
	}
}

func BuildFailResponse(errMsg string) Response {
	return Response{
		Code:    0,
		ErrMsg:  errMsg,
		Success: false,
		Data:    nil,
	}
}
