package valueobjects

import (
	"database/sql/driver"
	"fmt"
	"movie-api/domain/utils"
	"strings"
)

type Note struct {
	NoteValue float32
}

func Create(value float32) utils.Result[Note] {
	validate := validate(value)

	if strings.Count(validate, "") > 0 {
		return utils.Failure[Note](validate)
	}
	note := Note{NoteValue: value}

	return utils.Success[Note](note)
}

func validate(value float32) string {
	var result string

	if value < 0 {
		return "Value should be bigger than 0"
	}
	if value > 10 {
		return "Value should be smaller than 10"
	}

	return result
}

func (n *Note) Scan(value interface{}) error {
	note, ok := value.(float32)

	if !ok {
		return fmt.Errorf("failed to read value")
	}

	result := Create(note)
	n = &result.Value
	return nil
}

func (n Note) Value() (driver.Value, error) {
	return n.NoteValue, nil
}
