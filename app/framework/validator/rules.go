package validator

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"github.com/thedevsaddam/govalidator"
)

func init() {

}

func Run()  {
	maxWord()
}

func maxWord()  {
	// custom rules to take fixed length word.
	// e.g: max_word:5 will throw error if the field contains more than 5 words
	govalidator.AddCustomRule("max_word", func(field string, rule string, message string, value interface{}) error {
		valSlice := strings.Fields(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_word:"))

		if len(valSlice) > l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("the %s field must not be greater than %d words", field, l)
		}

		return nil
	})
}
