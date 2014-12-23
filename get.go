package model

import "strings"

const dot = "."

// Get gets the value from the data by the given dot-notation
// keypath. Returns nil if any of the data is missing.
func Get(data map[string]interface{}, keypath string) interface{} {
	value, _ := GetOK(data, keypath)
	return value
}

// GetOK gets the value from the data by the given dot-notation
// keypath, or returns the second argument false if any of the data
// is missing.
func GetOK(data map[string]interface{}, keypath string) (interface{}, bool) {
	return getOK(data, strings.Split(keypath, dot))
}

func getOK(data map[string]interface{}, keys []string) (interface{}, bool) {
	if len(keys) > 1 {
		var sub interface{}
		var ok bool
		if sub, ok = data[keys[0]]; !ok {
			return nil, false
		}
		var submap map[string]interface{}
		if submap, ok = sub.(map[string]interface{}); !ok {
			return nil, false
		}
		return getOK(submap, keys[1:])
	}
	value, ok := data[keys[0]]
	return value, ok
}
