package validation

import "errors"
import "fmt"
        
var String = str{}

type str struct{}

func (str) IsNotBlank(s string) error {
	if len(s) == 0 {
		return errors.New("can't be blank")
	}

  return nil
}

func (str) MinLen(s string, count int) error {
	if len(s) < count {
		return fmt.Errorf("can't be shorten then %v chars", count)
	}

  return nil
}

func (str) MaxLen(s string, count int) error {
	if len(s) > count {
		return fmt.Errorf("can't be longer then %v chars", count)
	}

  return nil
}
