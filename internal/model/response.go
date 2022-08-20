package model

type Response struct {
	Result  string      `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseCache struct {
	Response
	IsCached bool `json:"isCached"`
}

type ResponseError struct {
	Response
}
