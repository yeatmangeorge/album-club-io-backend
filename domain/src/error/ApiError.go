package error

type MessageErr interface {
	Message() string
	Status() int
	Error() string
}

type StandardMessageErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *StandardMessageErr) Error() string {
	return e.ErrError
}

func (e *StandardMessageErr) Message() string {
	return e.ErrMessage
}

func (e *StandardMessageErr) Status() int {
	return e.ErrStatus
}
