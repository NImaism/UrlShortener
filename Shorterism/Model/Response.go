package Model

type HttpResponse struct {
	ResMessage string
	Data       interface{}
}

func SuccessResponse(Data interface{}) HttpResponse {
	return HttpResponse{
		ResMessage: "Success",
		Data:       Data,
	}
}

func ErrorResponse(Data interface{}, errorMessage string) HttpResponse {
	return HttpResponse{
		ResMessage: errorMessage,
		Data:       Data,
	}
}
