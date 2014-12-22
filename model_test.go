package model_test

import (
	"errors"
	"testing"

	"github.com/cheekybits/is"
	"github.com/cheekybits/model"
)

var fokcalls = 0

func fok(data map[string]interface{}, key string) error {
	fokcalls++
	return nil
}

func TestModel(t *testing.T) {
	is := is.New(t)

	m := model.M{
		"name":   {fok, fok},
		"number": {fok},
	}

	data := map[string]interface{}{
		"name":   "Mat",
		"number": 123,
	}
	newdata, errs := m.Do(data)
	is.Equal(len(errs), 0)
	is.OK(newdata)
	is.Equal(newdata, data)
	is.Equal(fokcalls, 3)

	// errors
	err := errors.New("something went wrong")
	var ferr = func(data map[string]interface{}, key string) error {
		return err
	}
	m = model.M{
		"name":   {fok, ferr},
		"number": {fok},
	}
	newdata, errs = m.Do(data)
	is.Equal(1, len(errs))
	is.Equal(errs["name"][0], err)

}

func TestChangingData(t *testing.T) {
	is := is.New(t)

	var fchange = func(data map[string]interface{}, key string) error {
		data[key+"_copy"] = data[key]
		return nil
	}

	m := model.M{
		"name":   {fchange},
		"number": {fchange},
	}

	data := map[string]interface{}{
		"name":   "Mat",
		"number": 123,
	}

	newdata, errs := m.Do(data)
	is.Equal(len(errs), 0)
	is.Equal(newdata["name"], "Mat")
	is.Equal(newdata["name_copy"], "Mat")
	is.Equal(newdata["number"], 123)
	is.Equal(newdata["number_copy"], 123)
	if _, present := data["name_copy"]; present {
		is.Fail("shouoldn't mutate original map")
	}

}

func TestModelBefore(t *testing.T) {
	is := is.New(t)

	m := model.M{
		"name":   {fok, fok},
		"number": {fok},
	}.Before(func(m model.M, data map[string]interface{}) error {
		data["before"] = true
		return nil
	})

	data := map[string]interface{}{
		"name":   "Mat",
		"number": 123,
		"food":   "sausage",
	}

	newdata, errs := m.Do(data)
	is.Equal(len(errs), 0)
	is.Equal(newdata["before"], true)
	is.Equal(newdata["name"], "Mat")
	is.Equal(newdata["number"], 123)

	m = model.M{
		"name":   {fok, fok},
		"number": {fok},
	}.Before(func(m model.M, data map[string]interface{}) error {
		return errors.New("whoops")
	})
	newdata, errs = m.Do(data)
	is.Equal(len(errs), 1)
	is.Equal(errs["model"][0].Error(), "whoops")

}

func TestModelAfter(t *testing.T) {
	is := is.New(t)

	m := model.M{
		"name":   {fok, fok},
		"number": {fok},
	}.After(func(m model.M, data map[string]interface{}) error {
		data["After"] = true
		return nil
	})

	data := map[string]interface{}{
		"name":   "Mat",
		"number": 123,
		"food":   "sausage",
	}

	newdata, errs := m.Do(data)
	is.Equal(len(errs), 0)
	is.Equal(newdata["After"], true)
	is.Equal(newdata["name"], "Mat")
	is.Equal(newdata["number"], 123)

	m = model.M{
		"name":   {fok, fok},
		"number": {fok},
	}.After(func(m model.M, data map[string]interface{}) error {
		return errors.New("whoops")
	})
	newdata, errs = m.Do(data)
	is.Equal(len(errs), 1)
	is.Equal(errs["model"][0].Error(), "whoops")

}