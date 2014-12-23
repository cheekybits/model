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
	"name":   {model.IsString, model.IsRequired},
	"number": {model.IsNumber},
	"ok":     {model.IsBool, model.IsRequired},
	"tags":   {model.IsInterSlice},
}
```

We can process some data provided by the user (perhaps from JSON):

```
var data map[string]interface{}
json.Decode([]byte(`{
	"name": "Mat",
	"number": false,
	"tags": [""],
}`), &data)
```
