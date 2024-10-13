package utils

type Result[T any] struct {
	value     T
	error     string
	isSuccess bool
}

func Failure[T any](error string) Result[T] {
	return Result[T]{error: error, isSuccess: false}
}

func Success[T any](value T) Result[T] {
	return Result[T]{error: "", isSuccess: true, value: value}
}
