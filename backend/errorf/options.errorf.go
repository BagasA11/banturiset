package errorf

import (
	e "errors"
)

var (
	ErrRedundant = e.New("data redundant")
)
