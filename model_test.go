package model_test

import (
	"encoding/json"
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

func TestExample(t *testing.T) {
	is := is.New(t)

	var Person = model.M{
		"name":   {model.IsString, model.IsRequired},
		"number": {model.IsNumber},
		"ok":     {model.IsBool, model.IsRequired},
		"tags":   {model.IsInterfaceSlice},
	}

	var data map[string]interface{}
	err := json.Unmarshal([]byte(`{
		"name": 123,
		"number": false,
		"tags": ["one", "two", "three"]
	}`), &data)
	is.NoErr(err)

	_, errs := Person.Do(data)
	is.Equal(len(errs), 3)
	is.Equal(errs["name"][0].Error(), "must be string")
	is.Equal(errs["number"][0].Error(), "must be a number")
	is.Equal(errs["ok"][0].Error(), "is required")

	// marshal it
	j, err := json.Marshal(errs)
	is.NoErr(err)
	is.Equal(string(j), `{"name":["must be string"],"number":["must be a number"],"ok":["is required"]}`)

	j, err = json.MarshalIndent(errs, "", "  ")
	is.NoErr(err)
	is.Equal(string(j), `{
  "name": [
    "must be string"
  ],
  "number": [
    "must be a number"
  ],
  "ok": [
    "is required"
  ]
}`)

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

func TestModelRequired(t *testing.T) {
	is := is.New(t)

	fokcalls = 0
	m := model.M{
		"name":   {fok},
		"number": {fok, model.IsRequired},
	}

	data := map[string]interface{}{
		"number": 123,
	}
	newdata, errs := m.Do(data)
	is.Equal(len(errs), 0)
	is.OK(newdata)
	is.Equal(newdata, data)
	is.Equal(fokcalls, 2)

	m = model.M{
		"name":   {fok},
		"number": {fok, model.IsRequired},
	}

	required := m.Required()
	is.NotEqual(required, nil)
	is.NotEqual(len(required), 0)

	newdata, errs = required.Do(data)
	is.Equal(len(errs), 1)
	is.OK(errs["name"])
	is.Equal(errs["name"][0].Error(), "is required")

}

func TestModelRemove(t *testing.T) {
	is := is.New(t)

	fokcalls = 0
	m := model.M{
		"name":   {fok},
		"number": {fok, model.IsRequired},
	}

	is.Nil(m.Remove("number")["number"])

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

func TestNestedData(t *testing.T) {
	is := is.New(t)

	m := model.M{
		"address.postcode":       {model.IsRequired},
		"address.postcode.inner": {model.IsString},
		"address.postcode.outer": {model.IsString},
	}

	d := map[string]interface{}{
		"address": map[string]interface{}{
			"postcode": map[string]interface{}{
				"inner": "NG1",
				"outer": "8BC",
			},
		},
	}
	_, errs := m.Do(d)
	is.Equal(len(errs), 0)

	d = map[string]interface{}{
		"address": map[string]interface{}{
			"postcode": map[string]interface{}{
				"inner": 123,
				"outer": 456,
			},
		},
	}
	_, errs = m.Do(d)
	is.Equal(len(errs), 2)
	is.Equal(errs["address.postcode.inner"][0].Error(), "must be string")
	is.Equal(errs["address.postcode.outer"][0].Error(), "must be string")

	d = map[string]interface{}{
		"address": map[string]interface{}{},
	}
	_, errs = m.Do(d)
	is.Equal(len(errs), 1)
	is.Equal(errs["address.postcode"][0].Error(), "is required")

}
