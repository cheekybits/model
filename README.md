model
=====

Simple data modeling for Go.

### Concepts

  * Work with `map[string]interface{}` types holding raw data
  * Bind many funcs to each field, each func gets called when `Do` is asked to process the underlying data.
  * If the funcs return errors, they are collected and returned in a map of issues.
  * Use dot notation to supported nested objects.

### Example

Given this model:

```
var Person = model.M{
	"name":         {model.IsString, model.IsRequired},
	"number":       {model.IsNumber},
	"ok":           {model.IsBool, model.IsRequired},
	"tags":         {model.IsInterfaceSlice},
	"address.city": {model.IsRequired},
}
```

We can process some data provided by the user (perhaps from JSON):

```
var data map[string]interface{}
err := json.Unmarshal([]byte(`{
	"name": 123,
	"number": false,
	"tags": ["one", "two", "three"]
}`), &data)
```

And check that it is valid, by checking the length of the errors:

```
var errs model.Errs
if data, errs = Person.Do(data); len(errs) > 0 {
	// report errors to user
}
```

Adding `model.Strict` as a `Before` function, ensures no unexpected data appears in the map:

```
var Person = model.M{
	"name":         {model.IsString, model.IsRequired},
	"number":       {model.IsNumber},
	"ok":           {model.IsBool, model.IsRequired},
	"tags":         {model.IsInterfaceSlice},
	"address.city": {model.IsRequired},
}.Before(model.Strict)
```

  * You can write your own `Before` and `After` functions
  * Add your own validators easily enough too, just write a function that has this signature: `func(data map[string]interface{}, keypath string) error`
