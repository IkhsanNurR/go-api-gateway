package response

type Responses struct {
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	Status   int         `json:"status"`
	MetaData MetaData    `json:"meta_data"`
}

type ResponsesOne struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

type FailedResponses struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Error   any         `json:"error"`
}

type MetaData struct {
	Limit   int    `json:"limit"`
	Pages   int    `json:"pages"`
	Total   int    `json:"total"`
	SortBy  string `json:"sort_by"`
	SortKey string `json:"sort_key"`
}

func SuccessResponse(data interface{}, message string, status int) Responses {
	response := Responses{
		Data:    data,
		Message: message,
		Status:  status,
	}

	return response
}

func SuccessOneResponse(data interface{}, message string, status int) ResponsesOne {
	response := ResponsesOne{
		Data:    data,
		Message: message,
		Status:  status,
	}

	return response
}

func FailedResponse(data interface{}, message string, status int, err any) FailedResponses {
	response := FailedResponses{
		Data:    data,
		Message: message,
		Status:  status,
		Error:   err,
	}

	return response
}
