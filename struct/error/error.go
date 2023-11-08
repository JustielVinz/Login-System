package errors

type ResponseModel struct {
	RetCode string      `json:"retCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorModel struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
	Error     error  `json:"error"`
}
