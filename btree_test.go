package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	as := assert.New(t)
	tree := New(7).Add(5).Add(8).Add(3).Add(1).Add(10)
	expStr := `{
           "Left": {
              "Left": {
                 "Left": {
                    "Left": null,
                    "Value": 1,
                    "Right": null
                 },
                 "Value": 3,
                 "Right": null
              },
              "Value": 5,
              "Right": null
           },
           "Value": 7,
           "Right": {
              "Left": null,
              "Value": 8,
              "Right": {
                 "Left": null,
                 "Value": 10,
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
