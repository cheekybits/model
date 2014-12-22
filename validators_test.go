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
