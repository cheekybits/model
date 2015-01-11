package model

//go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "Type=BUILTINS,interface{}"

import (
	"errors"

	"github.com/cheekybits/genny/generic"
	"github.com/cheekybits/m"
)

// Type is the generic type placeholder.
type Type generic.Type

// IsType checks to make sure the value is a Type.
// Remains silent if the data is not present.
func IsType(data map[string]interface{}, keypath string) error {
	if v, ok := m.GetOK(data, keypath); ok {
		if _, ok := v.(Type); !ok {
			return errors.New("must be Type")
		}
	}
	return nil
}

// IsTypeSlice checks to make sure the value is a Type.
// Remains silent if the data is not present.
func IsTypeSlice(data map[string]interface{}, keypath string) error {
	if v, ok := m.GetOK(data, keypath); ok {
		if _, ok := v.([]Type); !ok {
			return errors.New("must be array of Type")
		}
	}
	return nil
}
