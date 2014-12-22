package model

import (
	"strings"
)

const (
	keyModel       = "model"
	nonFieldPrefix = "__"
	keyBefore      = nonFieldPrefix + "before"
	keyAfter       = nonFieldPrefix + "after"
)

type f func(data map[string]interface{}, key string) error

// M is a map providing []F functions for each field.
type M map[string][]f

// Do processes the data with this model and returns new data,
// and a slice of errors.
func (m M) Do(data map[string]interface{}) (map[string]interface{}, map[string][]error) {
	newdata := make(map[string]interface{})
	for k, v := range data {
		newdata[k] = v
	}
	errs := make(map[string][]error)

	// before
	if b, ok := m[keyBefore]; ok {
		if err := b[0](newdata, ""); err != nil {
			errs[keyModel] = append(errs[keyModel], err)
		}
	}

	for k := range m {
		if strings.HasPrefix(k, nonFieldPrefix) {
			continue
		}
		for _, fn := range m[k] {
			if err := fn(newdata, k); err != nil {
				errs[k] = append(errs[k], err)
			}
		}
	}

	// after
	if b, ok := m[keyAfter]; ok {
		if err := b[0](newdata, ""); err != nil {
			errs[keyModel] = append(errs[keyModel], err)
		}
	}

	return newdata, errs
}

// Before adds a pre-process callback to the model.
// Will be called before the fields are processed in Do.
// The error returned from Before functions will be set in the data
// with the "model" key.
// Mutates as well as returns M.
func (m M) Before(fn func(M, map[string]interface{}) error) M {
	m[keyBefore] = []f{func(data map[string]interface{}, _ string) error {
		return fn(m, data)
	}}
	return m
}

// After adds a post-process callback to the model.
// Will be called after the fields are processed in Do.
// The error returned from After functions will be set in the data
// with the "model" key.
// Mutates as well as returns M.
func (m M) After(fn func(M, map[string]interface{}) error) M {
	m[keyAfter] = []f{func(data map[string]interface{}, _ string) error {
		return fn(m, data)
	}}
	return m
}