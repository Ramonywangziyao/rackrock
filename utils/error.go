package utils

type RockError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (err RockError) Error() string {
	return err.Msg
}

func NewRockErr(code int, msg string) RockError {
	return RockError{
		Code: code,
		Msg:  msg,
	}
}

func NewBusinessErr(msg string) RockError {
	return RockError{
		Code: ServiceError.Code,
		Msg:  msg,
	}
}

var (
	AuthError    = NewRockErr(100, "not auth")
	ServiceError = NewRockErr(300, "service error")
)
