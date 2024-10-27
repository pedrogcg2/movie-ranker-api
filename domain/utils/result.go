package utils

type Result[T any] struct {
	Value     T
	Error     string
	IsSuccess bool
}

func Failure[T any](error string) Result[T] {
	return Result[T]{Error: error, IsSuccess: false}
}

func Success[T any](value T) Result[T] {
	return Result[T]{Error: "", IsSuccess: true, Value: value}
}
