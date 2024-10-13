package valueobjects

import (
	"strings"
)

type Note struct {
	value float32
}

func Create(value float32) Result[Note] {
	validate := validate(value)

	if strings.Count(validate, "") > 0 {
		return Failure[Note](validate)
	}
	note := Note{value: value}

	return Success[Note](note)
}

func validate(value float32) string {
	var result string = ""

	if value < 0 {
		return "Value should be bigger than 0"
	}
	if value > 10 {
		return "Value should be smaller than 10"
	}

	return result
}
