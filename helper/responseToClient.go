package helper

type ResponseClientModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseClient(code int, msg string, data interface{}) ResponseClientModel {
	return ResponseClientModel{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

type CustomResponse map[string]interface{}
