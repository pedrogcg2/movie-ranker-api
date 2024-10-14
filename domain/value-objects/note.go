package valueobjects

import (
	"strings"

	"github.com/pedrogcg2/movie-ranker-api/domain/utils"
)

type Note struct {
	value float32
}

func Create(value float32) utils.Result[Note] {
	validate := validate(value)

	if strings.Count(validate, "") > 0 {
		return utils.Failure[Note](validate)
	}
	note := Note{value: value}

	return utils.Success[Note](note)
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
