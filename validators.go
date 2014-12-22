package model

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
