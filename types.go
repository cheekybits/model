// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package model

import "errors"

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "bool=BUILTINS"

// IsBool checks to make sure the value is a bool.
// Remains silent if the data is not present.
func IsBool(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(bool); !ok {
			return errors.New("must be bool")
		}
	}
	return nil
}

// IsBoolSlice checks to make sure the value is a bool.
// Remains silent if the data is not present.
func IsBoolSlice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]bool); !ok {
			return errors.New("must be array of bool")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "byte=BUILTINS"

// IsByte checks to make sure the value is a byte.
// Remains silent if the data is not present.
func IsByte(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(byte); !ok {
			return errors.New("must be byte")
		}
	}
	return nil
}

// IsByteSlice checks to make sure the value is a byte.
// Remains silent if the data is not present.
func IsByteSlice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]byte); !ok {
			return errors.New("must be array of byte")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "complex128=BUILTINS"

// IsComplex128 checks to make sure the value is a complex128.
// Remains silent if the data is not present.
func IsComplex128(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(complex128); !ok {
			return errors.New("must be complex128")
		}
	}
	return nil
}

// IsComplex128Slice checks to make sure the value is a complex128.
// Remains silent if the data is not present.
func IsComplex128Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]complex128); !ok {
			return errors.New("must be array of complex128")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "complex64=BUILTINS"

// IsComplex64 checks to make sure the value is a complex64.
// Remains silent if the data is not present.
func IsComplex64(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(complex64); !ok {
			return errors.New("must be complex64")
		}
	}
	return nil
}

// IsComplex64Slice checks to make sure the value is a complex64.
// Remains silent if the data is not present.
func IsComplex64Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]complex64); !ok {
			return errors.New("must be array of complex64")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "error=BUILTINS"

// IsError checks to make sure the value is a error.
// Remains silent if the data is not present.
func IsError(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(error); !ok {
			return errors.New("must be error")
		}
	}
	return nil
}

// IsErrorSlice checks to make sure the value is a error.
// Remains silent if the data is not present.
func IsErrorSlice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]error); !ok {
			return errors.New("must be array of error")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "float32=BUILTINS"

// IsFloat32 checks to make sure the value is a float32.
// Remains silent if the data is not present.
func IsFloat32(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(float32); !ok {
			return errors.New("must be float32")
		}
	}
	return nil
}

// IsFloat32Slice checks to make sure the value is a float32.
// Remains silent if the data is not present.
func IsFloat32Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]float32); !ok {
			return errors.New("must be array of float32")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "float64=BUILTINS"

// IsFloat64 checks to make sure the value is a float64.
// Remains silent if the data is not present.
func IsFloat64(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(float64); !ok {
			return errors.New("must be float64")
		}
	}
	return nil
}

// IsFloat64Slice checks to make sure the value is a float64.
// Remains silent if the data is not present.
func IsFloat64Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]float64); !ok {
			return errors.New("must be array of float64")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "int=BUILTINS"

// IsInt checks to make sure the value is a int.
// Remains silent if the data is not present.
func IsInt(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(int); !ok {
			return errors.New("must be int")
		}
	}
	return nil
}

// IsIntSlice checks to make sure the value is a int.
// Remains silent if the data is not present.
func IsIntSlice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]int); !ok {
			return errors.New("must be array of int")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "int16=BUILTINS"

// IsInt16 checks to make sure the value is a int16.
// Remains silent if the data is not present.
func IsInt16(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(int16); !ok {
			return errors.New("must be int16")
		}
	}
	return nil
}

// IsInt16Slice checks to make sure the value is a int16.
// Remains silent if the data is not present.
func IsInt16Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]int16); !ok {
			return errors.New("must be array of int16")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "int32=BUILTINS"

// IsInt32 checks to make sure the value is a int32.
// Remains silent if the data is not present.
func IsInt32(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(int32); !ok {
			return errors.New("must be int32")
		}
	}
	return nil
}

// IsInt32Slice checks to make sure the value is a int32.
// Remains silent if the data is not present.
func IsInt32Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]int32); !ok {
			return errors.New("must be array of int32")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "int64=BUILTINS"

// IsInt64 checks to make sure the value is a int64.
// Remains silent if the data is not present.
func IsInt64(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(int64); !ok {
			return errors.New("must be int64")
		}
	}
	return nil
}

// IsInt64Slice checks to make sure the value is a int64.
// Remains silent if the data is not present.
func IsInt64Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]int64); !ok {
			return errors.New("must be array of int64")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "int8=BUILTINS"

// IsInt8 checks to make sure the value is a int8.
// Remains silent if the data is not present.
func IsInt8(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(int8); !ok {
			return errors.New("must be int8")
		}
	}
	return nil
}

// IsInt8Slice checks to make sure the value is a int8.
// Remains silent if the data is not present.
func IsInt8Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]int8); !ok {
			return errors.New("must be array of int8")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "rune=BUILTINS"

// IsRune checks to make sure the value is a rune.
// Remains silent if the data is not present.
func IsRune(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(rune); !ok {
			return errors.New("must be rune")
		}
	}
	return nil
}

// IsRuneSlice checks to make sure the value is a rune.
// Remains silent if the data is not present.
func IsRuneSlice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]rune); !ok {
			return errors.New("must be array of rune")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "string=BUILTINS"

// IsString checks to make sure the value is a string.
// Remains silent if the data is not present.
func IsString(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(string); !ok {
			return errors.New("must be string")
		}
	}
	return nil
}

// IsStringSlice checks to make sure the value is a string.
// Remains silent if the data is not present.
func IsStringSlice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]string); !ok {
			return errors.New("must be array of string")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "uint=BUILTINS"

// IsUint checks to make sure the value is a uint.
// Remains silent if the data is not present.
func IsUint(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(uint); !ok {
			return errors.New("must be uint")
		}
	}
	return nil
}

// IsUintSlice checks to make sure the value is a uint.
// Remains silent if the data is not present.
func IsUintSlice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]uint); !ok {
			return errors.New("must be array of uint")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "uint16=BUILTINS"

// IsUint16 checks to make sure the value is a uint16.
// Remains silent if the data is not present.
func IsUint16(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(uint16); !ok {
			return errors.New("must be uint16")
		}
	}
	return nil
}

// IsUint16Slice checks to make sure the value is a uint16.
// Remains silent if the data is not present.
func IsUint16Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]uint16); !ok {
			return errors.New("must be array of uint16")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "uint32=BUILTINS"

// IsUint32 checks to make sure the value is a uint32.
// Remains silent if the data is not present.
func IsUint32(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(uint32); !ok {
			return errors.New("must be uint32")
		}
	}
	return nil
}

// IsUint32Slice checks to make sure the value is a uint32.
// Remains silent if the data is not present.
func IsUint32Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]uint32); !ok {
			return errors.New("must be array of uint32")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "uint64=BUILTINS"

// IsUint64 checks to make sure the value is a uint64.
// Remains silent if the data is not present.
func IsUint64(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(uint64); !ok {
			return errors.New("must be uint64")
		}
	}
	return nil
}

// IsUint64Slice checks to make sure the value is a uint64.
// Remains silent if the data is not present.
func IsUint64Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]uint64); !ok {
			return errors.New("must be array of uint64")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "uint8=BUILTINS"

// IsUint8 checks to make sure the value is a uint8.
// Remains silent if the data is not present.
func IsUint8(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(uint8); !ok {
			return errors.New("must be uint8")
		}
	}
	return nil
}

// IsUint8Slice checks to make sure the value is a uint8.
// Remains silent if the data is not present.
func IsUint8Slice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]uint8); !ok {
			return errors.New("must be array of uint8")
		}
	}
	return nil
}

// //go:generate genny -pkg="model" -in=$GOFILE -out=types.go gen "uintptr=BUILTINS"

// IsUintptr checks to make sure the value is a uintptr.
// Remains silent if the data is not present.
func IsUintptr(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.(uintptr); !ok {
			return errors.New("must be uintptr")
		}
	}
	return nil
}

// IsUintptrSlice checks to make sure the value is a uintptr.
// Remains silent if the data is not present.
func IsUintptrSlice(data map[string]interface{}, keypath string) error {
	if v, ok := GetOK(data, keypath); ok {
		if _, ok := v.([]uintptr); !ok {
			return errors.New("must be array of uintptr")
		}
	}
	return nil
}
