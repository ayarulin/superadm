package validation

import (
	"fmt"
	"time"
)

var Time = timevalidators{}

type timevalidators struct{}

func (timevalidators) IsNotAfter(t time.Time, after time.Time) error {
	if after.Compare(t) < 0 {
		return fmt.Errorf("can't be after %v", t)
	}

	return nil
}

func (timevalidators) IsNotZero(t time.Time) error {
	if t.IsZero() {
		return fmt.Errorf("can't be zero value %v", t)
	}

	return nil
}
