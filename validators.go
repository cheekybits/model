package model

import (
	"errors"

	"github.com/cheekybits/m"
)

var errRequired = errors.New("is required")

// IsRequired checks to make sure key is present and the value
// is non-nil.
func IsRequired(data map[string]interface{}, keypath string) error {
	var v interface{}
	var ok bool
	if v, ok = m.GetOK(data, keypath); !ok {
		return errRequired
	}
	if v == nil {
		return errRequired
	}
	return nil
}

type errUnexpectedField string

func (f errUnexpectedField) Error() string {
	return "unexpected '" + string(f) + "'"
}

// Strict is a Before function that ensures there are no extra data
// fields in the model.
func Strict(m M, data map[string]interface{}) error {
	for k := range data {
		if _, present := m[k]; !present {
			return errUnexpectedField(k)
		}
	}
	return nil
}

var errNotNumber = errors.New("must be a number")

// IsNumber checks to make sure the value is a number.
// Remains silent if the data is not present.
func IsNumber(data map[string]interface{}, keypath string) error {
	if v, ok := m.GetOK(data, keypath); ok {
		switch v.(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64, complex64, complex128:
			return nil
		}
		return errNotNumber
	}
	return nil
}
