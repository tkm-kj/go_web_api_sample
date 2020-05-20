package serializer

// ネストさせないとタグが反映されなかった
type ErrorContent struct {
	Message string `json:"message"`
}
type ErrorSerializer struct {
	Error *ErrorContent `json:"error"`
}

func NewErrorSerializer(err error) *ErrorSerializer {
	return &ErrorSerializer{
		Error: &ErrorContent{
			Message: err.Error(),
		},
	}
}
