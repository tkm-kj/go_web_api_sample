package repository

// 名前は Error で終わらせる https://github.com/golang/go/wiki/Errors#naming
type RecordNotFoundError struct {
}

func (e *RecordNotFoundError) Error() string {
	return "record not found"
}
