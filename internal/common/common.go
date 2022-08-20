package common

import "books.api/internal/model"

func Response(data interface{}) interface{} {
	resp := model.Response{
		Result:  "success",
		Message: "Success",
		Data:    data,
	}

	return resp
}

func ResponseCache(data interface{}, isCached bool) interface{} {
	resp := model.ResponseCache{
		Response: model.Response{
			Result:  "success",
			Message: "Success",
			Data:    data,
		},
		IsCached: isCached,
	}

	return resp
}

func ResponseError(result string, message string, data interface{}) interface{} {
	resp := model.ResponseError{
		Response: model.Response{
			Result:  result,
			Message: message,
			Data:    data,
		},
	}

	return resp
}

func ResponseNotFound() interface{} {
	resp := model.ResponseError{
		Response: model.Response{
			Result:  "notfound",
			Message: "Data not found",
			Data:    nil,
		},
	}

	return resp
}
