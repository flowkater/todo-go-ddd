package errors

import "fmt"

// AppError는 애플리케이션의 에러를 나타내는 인터페이스입니다.
type AppError interface {
	error
	Status() int
}

// HTTPError는 HTTP 에러를 나타내는 구조체입니다.
type HTTPError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Internal   error  `json:"-"`
}

func (e *HTTPError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Internal)
	}
	return e.Message
}

func (e *HTTPError) Status() int {
	return e.StatusCode
}

// NewHTTPError creates a new HTTPError
func NewHTTPError(statusCode int, message string, internal error) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
		Internal:   internal,
	}
}

// IsNotFound checks if error is a not found error
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*HTTPError); ok {
		return e.StatusCode == 404
	}
	return err.Error() == "ent: todo not found"
}
