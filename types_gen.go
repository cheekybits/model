package model

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "Type=BUILTINS"

// import (
// 	"errors"

// 	"github.com/cheekybits/genny/generic"
// )

// // Type is the generic type placeholder.
// type Type generic.Type

// // IsType checks to make sure the value is a Type.
// // Remains silent if the data is not present.
// func IsType(data map[string]interface{}, key string) error {
// 	if v, ok := data[key]; ok {
// 		if _, ok := v.(Type); !ok {
// 			return errors.New(key + " must be Type")
// 		}
// 	}
// 	return nil
// }

// // IsTypeSlice checks to make sure the value is a Type.
// // Remains silent if the data is not present.
// func IsTypeSlice(data map[string]interface{}, key string) error {
// 	if v, ok := data[key]; ok {
// 		if _, ok := v.([]Type); !ok {
// 			return errors.New(key + " must be array of Type")
// 		}
// 	}
// 	return nil
// }
