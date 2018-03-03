package sql

import (
	"shendu.com/errors"
)

func NewError(message ...interface{}) *errors.Error {
	return errors.NewCode(errors.ErrInternal, 4, message...)
}
