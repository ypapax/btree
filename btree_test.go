package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	as := assert.New(t)
	tree := New(7).Add(5).Add(8).Add(3).Add(1).Add(10)
	expStr := `{
           "right/": {
              "right/": {
                 "[": 10
              },
              "[": 8
           },
           "[": 7,
           "left-_": {
              "[": 5,
              "left-_": {
                 "[": 3,
                 "left-_": {
                    "[": 1
                 }
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
