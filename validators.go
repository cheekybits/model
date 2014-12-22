package model

import "errors"

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

var errNotNumber = errors.New("expected number")

func IsNumber(data map[string]interface{}, key string) error {
	if v, ok := data[key]; ok {
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
