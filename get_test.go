package model_test

import (
	"testing"

	"github.com/cheekybits/is"
	"github.com/cheekybits/model"
)

func TestGet(t *testing.T) {
	is := is.New(t)

	for _, test := range []struct {
		M   map[string]interface{}
		K   string
		Exp interface{}
		OK  bool
	}{
		{
			map[string]interface{}{"name": "Tylor"},
			"name",
			"Tylor",
			true,
		},
		{
			map[string]interface{}{"address": map[string]interface{}{"city": "London"}},
			"address.city",
			"London",
			true,
		},
		{
			map[string]interface{}{
				"address": map[string]interface{}{
					"postcode": map[string]interface{}{
						"inner": "NG19",
					},
				},
			},
			"address.postcode.inner",
			"NG19",
			true,
		},
		{
			nil,
			"address.postcode.inner",
			nil,
			false,
		},
		{
			map[string]interface{}{"address": map[string]interface{}{"city": "London"}},
			"address.country",
			nil,
			false,
		},
	} {

		actual, actualOK := model.GetOK(test.M, test.K)

		is.Equal(actual, test.Exp)
		is.Equal(actualOK, test.OK)

	}

}
