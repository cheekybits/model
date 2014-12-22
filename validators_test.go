package model_test

import (
	"testing"

	"github.com/cheekybits/is"
	"github.com/cheekybits/model"
)

func TestStrict(t *testing.T) {
	is := is.New(t)

	m := model.M{
		"name":   {},
		"number": {},
	}.Before(model.Strict)

	d := map[string]interface{}{
		"name":   "Mat",
		"number": 123,
		"extra":  true,
	}

	_, errs := m.Do(d)
	is.Equal(len(errs), 1)
	is.Equal(errs["model"][0].Error(), "unexpected 'extra'")

}

func TestIsNumber(t *testing.T) {
	is := is.New(t)

	m := model.M{
		"number": {model.IsNumber},
	}

	d := map[string]interface{}{
		"number": 123,
	}
	_, errs := m.Do(d)
	is.Equal(len(errs), 0)

	d = map[string]interface{}{
		"number": "123",
	}
	_, errs = m.Do(d)
	is.Equal(len(errs), 1)
	is.Equal(errs["number"][0].Error(), "must be a number")

}

func TestIsRequired(t *testing.T) {

	is := is.New(t)

	m := model.M{
		"number": {model.IsRequired},
	}

	d := map[string]interface{}{
		"number": 123,
	}
	_, errs := m.Do(d)
	is.Equal(len(errs), 0)

	d = map[string]interface{}{
		"number": nil,
	}
	_, errs = m.Do(d)
	is.Equal(len(errs), 1)
	is.Equal(errs["number"][0].Error(), "is required")

	d = map[string]interface{}{}
	_, errs = m.Do(d)
	is.Equal(len(errs), 1)
	is.Equal(errs["number"][0].Error(), "is required")

}
