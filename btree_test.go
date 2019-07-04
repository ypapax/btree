package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	as := assert.New(t)
	tree := New(7).Add(5).Add(8).Add(3).Add(1).Add(10)
	expStr := `{
           "Value": 7,
           "Left": {
              "Value": 5,
              "Left": {
                 "Value": 3,
                 "Left": {
                    "Value": 1,
                    "Left": null,
                    "Right": null
                 },
                 "Right": null
              },
              "Right": null
           },
           "Right": {
              "Value": 8,
              "Left": null,
              "Right": {
                 "Value": 10,
                 "Left": null,
                 "Right": null
              }
           }
        }`
	expTree, err := Parse(expStr)
	if !as.NoError(err) {
		return
	}

	t.Log(tree.String())
	as.Equal(expTree, tree)
}
